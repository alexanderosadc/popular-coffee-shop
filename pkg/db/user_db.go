package db

import (
	"errors"
	"fmt"
	"sync"

	"database/sql"

	"github.com/alexanderosadc/popular-coffee-shop/pkg/domain"
	_ "github.com/lib/pq"
)

type Repository interface {
	ConnectToDB(host, port, user, password, dbname string) error
	CreateOrUpdateUser(userID, membershipType string) error
	GetUserCofee(userID, requestedCofeeType string) (int, error)
	GetUser(userID string) (*domain.User, error)
	CloseConnection()
}

type SqlRepo struct {
	mu           sync.RWMutex
	db           *sql.DB
	cofeeClients map[string]domain.User
}

func (r *SqlRepo) ConnectToDB(host, port, user, password, dbname string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	r.db = db

	fmt.Println("successfully connected to postgress")
	return nil
}

func (r *SqlRepo) CreateOrUpdateUser(userID string, membershipType string) error {
	user, _ := r.GetUser(userID)
	if user == nil {
		r.createUser(userID, membershipType)
		return nil
	}
	r.mu.Lock()
	user.Membership = membershipType
	r.mu.Lock()

	return nil
}

func (r *SqlRepo) GetUserCofee(userID string, requestedCofeeType string) (int, error) {
	return 0, nil
}

func (r *SqlRepo) GetUser(userID string) (*domain.User, error) {
	r.mu.RLock()
	user, ok := r.cofeeClients[userID]
	r.mu.RUnlock()
	if ok {
		return &user, nil
	}

	return nil, errors.New("there is no such user in DB")
}

func (r *SqlRepo) createUser(userID string, membershipType string) {
	r.mu.Lock()
	user := domain.User{Membership: membershipType}
	r.cofeeClients[userID] = user
	r.mu.Unlock()
}

func (r *SqlRepo) CloseConnection() {
	if err := r.db.Close(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("connection closed")
}
