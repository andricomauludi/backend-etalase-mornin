package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	//get the cookie off request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [main]"})
	}
	//decode / validate it

	// Parse takes the token string and a function for looking up the key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [Exp]"})
		}

		//find the user with token sub
		var user models.User
		models.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [Id not found]"})
		}

		//attach to req
		c.Set("user", user)

		//continue

		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token not valid]"})
	}

}

func Authorization(c *gin.Context) {
	//get the cookie off request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [main]"})
	}
	//decode / validate it

	// Parse takes the token string and a function for looking up the key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [Exp]"})
		}

		//find the user with token username
		var user models.User
		models.DB.First(&user, claims["username"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [Id not found]"})
		}

		//attach to req
		c.Set("user", user)

		//continue

		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token not valid]"})
	}

}
