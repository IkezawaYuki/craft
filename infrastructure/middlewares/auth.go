package middlewares

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func verifyFirebaseIDToken(ctx gin.Context, auth *auth.Client) (*auth.Token, error) {

}
