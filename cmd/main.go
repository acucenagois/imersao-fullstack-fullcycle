package main

import (
	"github.com/jinzhu/gorm"
	"github.com/acucenagois/imersaoFullstackFullcycle/application/grpc"
	"github.com/acucenagois/imersaoFullstackFullcycle/infrastructure/db"
	"os"
)

var database *gorm.DB

func main() {
	database := db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}