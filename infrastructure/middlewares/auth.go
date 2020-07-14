package middlewares

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

func verifyFirebaseIDToken(ctx *gin.Context, auth *auth.Client) (*auth.Token, error) {
	headerAuth := ctx.GetHeader("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)
	return jwtToken, err
}

func FirebaseGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authClientGin, ok := c.Get("firebase")
		if !ok {
			c.JSON(500, "firebase is empty")
			return
		}
		authClient := authClientGin.(*auth.Client)
		jwtToken, err := verifyFirebaseIDToken(c, authClient)
		if err != nil {
			c.JSON(401, "not authentication")
			return
		}
		c.Set("auth", jwtToken)
		c.Next()
	}
}

func FirebaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authClientGin, ok := c.Get("firebase")
		if !ok {
			c.JSON(500, "firebase is empty")
			return
		}
		authClient := authClientGin.(*auth.Client)
		jwtToken, _ := verifyFirebaseIDToken(c, authClient)
		c.Set("auth", jwtToken)
		c.Next()
	}
}
