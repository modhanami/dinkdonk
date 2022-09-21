package main

import (
	"github.com/gin-gonic/gin"
	"github.com/modhanami/dinkdonk/user-service/logger"
	"github.com/modhanami/dinkdonk/user-service/model/gormmodel"
	"github.com/modhanami/dinkdonk/user-service/repository/gormrepository"
	"github.com/modhanami/dinkdonk/user-service/server"
	"github.com/modhanami/dinkdonk/user-service/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	r := gin.Default()

	db, err := initGormDB()
	if err != nil {
		return err
	}
	err = createFakeUsers(db)
	if err != nil {
		return err
	}

	l := logger.NewZapLogger()
	repo := gormrepository.NewUserRepository(db, nil)
	svc := service.NewUserService(repo)
	handler := server.NewUserHandler(svc, l)

	handler.Mount(r)

	l.Info("starting user service")

	return r.Run("localhost:30027")
}

func initGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&gormmodel.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createFakeUsers(db *gorm.DB) error {
	users := []gormmodel.User{
		{
			Name:     "User 1",
			Email:    "user1@dinkdonk.com",
			Password: "password1",
		},
		{
			Name:     "User 2",
			Email:    "user2@dinkdonk.com",
			Password: "password2",
		},
	}

	return db.Create(&users).Error
}
