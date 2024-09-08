package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Get the Authorization header from the request
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [main]"})
		return
	}

	// Log the Authorization header to check if it's received
	fmt.Println("Received Authorization header:", authHeader)

	// Split the "Bearer" prefix
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token format]"})
		return
	}

	tokenString := parts[1]

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token invalid]"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [Exp]"})
			return
		}

		// Find the user with the token subject
		var user models.User
		models.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [Id not found]"})
			return
		}

		// Attach the user to the context
		c.Set("user", user)

		// Continue with the request
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token not valid]"})
	}
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(c *gin.Context) {
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

			stringrole := claims["role"].(string)
			rolesVal := stringToIntSlice(stringrole)
			fmt.Println(stringrole)
			// return

			roles := rolesVal
			validation := make(map[int]int)
			for _, val := range roles {
				validation[val] = 0
			}

			// for _, val := range validRoles {
			// 	if _, ok := validation[val]; !ok {
			// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [Role]"})
			// 	}
			// }
			for i := 0; i < len(validRoles); i++ {
				for y := 0; y < len(roles); y++ {
					if roles[y] == validRoles[i] {
						c.Next()

					}

				}
			}

			//continue

			c.Next()

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token not valid]"})
		}

	}
}
func stringToIntSlice(str string) []int {
	strSlice := strings.Split(str, " ")
	intSlice := make([]int, 0, len(strSlice))

	for _, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			// handle error if necessary
			// continue
		}
		intSlice = append(intSlice, num)
	}

	return intSlice
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define allowed origins
		allowedOrigins := map[string]bool{
			// "http://localhost:3000": true,
			"http://www.ceumonny.com:3000": true,
			"http://ceumonny.com:3000":     true,
			// Add other allowed origins here if needed
		}

		origin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed and set the Access-Control-Allow-Origin header accordingly
		if allowedOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			// Optionally handle disallowed origins
			// For example, you could set a default allowed origin or simply omit the header
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://www.ceumonny.com:3000")
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
