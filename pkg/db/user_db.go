package db

import (
	"fmt"
	"sync"

	"github.com/alexanderosadc/popular-coffee-shop/pkg/domain"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	ConnectToDB(host, port, user, password, dbname string) error
	Create(user *domain.User) error
	ReadByID(ID string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(ID string) error
}

type SqlRepo struct {
	mu           sync.RWMutex
	db           gorm.DB
	cofeeClients map[string]domain.User
}

func (r *SqlRepo) ConnectToDB(host, port, user, password, dbname string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("successfully connected to db")
	r.db = *db

	if err := db.AutoMigrate(&domain.User{}, &domain.CofeeType{}); err != nil {
		return err
	}

	fmt.Println("successfully migration of tables")
	return nil
}

func (r *SqlRepo) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// ReadByID retrieves a User from the database based on the given ID
func (r *SqlRepo) ReadByID(ID string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", ID).Preload("PurchaseHistory").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates an existing User in the database
func (r *SqlRepo) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

// Delete deletes a User from the database based on the given ID
func (r *SqlRepo) Delete(Id string) error {
	return r.db.Where("id = ?", Id).Delete(&domain.User{}).Error
}
