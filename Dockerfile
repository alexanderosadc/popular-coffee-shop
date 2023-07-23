FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o popular-cofee-shop

ENTRYPOINT [ "./popular-cofee-shop" ]
