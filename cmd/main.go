package main

import (
	"blog/adapters/inbound/controller"
	"blog/adapters/outbound/postgresql"
	"blog/application/services/author"
	authorsession "blog/application/services/author_session"
	services "blog/application/services/post"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	authorRepo := postgresql.NewAuthorRepositoryImpl(db)
	authorSessionRepo := postgresql.NewAuthorSessionRepositoryImpl(db)
	postRepo := postgresql.NewPostRepositoryImpl(db)

	sessionService := authorsession.NewAuthorSessionService(authorSessionRepo)
	authorService := author.NewAuthorService(authorRepo, sessionService)
	postService := services.NewPostService(postRepo)

	router := gin.Default()
	apiV1 := router.Group("/api/v1")

	controller.NewPostController(apiV1, postService)
	controller.NewAuthorController(apiV1, authorService)
	router.Run(":8080")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in f", r)
		}
	}()
}

func initDatabase() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&postgresql.Post{}, &postgresql.Author{}, &postgresql.AuthorSession{})

	return db
}
