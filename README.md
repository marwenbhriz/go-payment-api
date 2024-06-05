module cmd/go-api-sample-todo

go 1.22.3

1. run mysql db
docker run --name payments-db -e MYSQL_ROOT_PASSWORD=root -d -p 3307:3306 mysql:latest
docker exec -it payments-db bash
mysql -u root -proot
CREATE DATABASE payments;

go mod init github.com/marwenbhriz/go-payment-api
go get .
go get github.com/marwenbhriz/go-rest-api/controllers/productcontroller
go get github.com/marwenbhriz/go-rest-api/models

go get gorm.io/driver/mysql