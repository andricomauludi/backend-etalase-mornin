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

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token not valid]"})

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
			continue
		}
		intSlice = append(intSlice, num)
	}

	return intSlice
}

// func parseStringToIntSlice(str string) []int {
// 	values := make([]int, 0)

// 	// Split string by comma
// 	strValues := strings.Split(str, ",")

// 	// Convert each string value to int
// 	for _, s := range strValues {
// 		num, err := strconv.Atoi(strings.TrimSpace(s))
// 		if err != nil {
// 			return nil
// 		}
// 		values = append(values, num)
// 	}

// 	return values
// }

// func Authorization2(validRoles []int) gin.HandlerFunc {
// 	return func(context *gin.Context) {

// 		if len(context.Keys) == 0 {
// 			ReturnUnauthorized(context)
// 		}

// 		rolesVal := context.Keys["Roles"]
// 		fmt.Println("roles", rolesVal)
// 		if rolesVal == nil {
// 			ReturnUnauthorized(context)
// 		}

// 		roles := rolesVal.([]int)
// 		validation := make(map[int]int)
// 		for _, val := range roles {
// 			validation[val] = 0
// 		}

// 		for _, val := range validRoles {
// 			if _, ok := validation[val]; !ok {
// 				ReturnUnauthorized(context)
// 			}
// 		}
// 	}
// }
