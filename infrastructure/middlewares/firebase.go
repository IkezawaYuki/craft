package middlewares

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"os"
)

func Firebase() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := option.WithCredentialsFile(os.Getenv("KEY_JSON_PATH"))
		config := &firebase.Config{
			ProjectID: os.Getenv("PROJECT_ID"),
		}

	}
}
