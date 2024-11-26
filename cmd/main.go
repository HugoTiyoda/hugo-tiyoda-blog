package main

import (
	"blog/adapters/inbound/controller"
	"blog/adapters/outbound/postgresql"
	"blog/application/services/author"
	authorsession "blog/application/services/author_session"
	services "blog/application/services/post"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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

	startSessionCleaner(sessionService)
	AuthMiddleware(sessionService)

	controller.NewAuthorController(apiV1, authorService)
	
	protected := router.Group("/api/protected")
	protected.Use(AuthMiddleware(sessionService))
	{
		controller.NewPostController(protected, postService)
	}
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

func startSessionCleaner(sessionService *authorsession.AuthorSessionService) {
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := sessionService.CleanExpiredSessions(); err != nil {
					log.Printf("failed to clean expired sessions: %v", err)
				}
			}
		}
	}()
}

func AuthMiddleware(sessionService *authorsession.AuthorSessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		session, err := sessionService.ValidateSession(tokenString)
		if err != nil {
			switch err.Error() {
			case "session expired":
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			case "session not found":
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			}
			return
		}

		c.Set("session", session)
		c.Set("authorId", session.AuthorId)

		c.Next()
	}
}
