package authcontroller

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
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email/ pass off req body
	var body models.User //mengambil body dari post api dan menyocokkan dengan yang ada di model

	messagebody := make(map[string]interface{}) //membuat message untuk response

	// Insert the inner JSON object into the outer JSON object
	messagebody["message"] = "Body not found or error, please try again"

	if c.ShouldBindJSON(&body) != nil { //apabila body yang diberikan tidak mengembalikan apa apa
		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messagebody, //memberikan pesan eror
		})

		return
	}

	//Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10) //melakukan hashing pada password untuk disimpan

	messagepass := make(map[string]interface{})

	// Insert the inner JSON object into the outer JSON object
	messagepass["message"] = "Failed to hash password, please try again"

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messagepass, //memberikan pesan eror
		})

		return
	}

	//Create the user
	//1 superadmin
	//2 admin
	//3 cashier
	//4 dashboard management
	//5 customer
	user := models.User{Fullname: body.Fullname, Username: body.Username, Password: string(hash), Role: "5"} //membuat user pada mysql dengan api post body yang dikirimkan
	result := models.DB.Create(&user)

	messageinsert := make(map[string]interface{})

	// Insert the inner JSON object into the outer JSON object
	messageinsert["message"] = "Failed to create user, consider to change your username"

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messageinsert, //memberikan pesan eror
		})

		return
	}

	//respond

	// var isian = map[string]string{"username": body.Username, "Nama Lengkap": body.Fullname}
	// var responses = map[string]string{"message": "Account Created!"}
	// responses["data"] = isian

	outer := make(map[int]interface{})

	// Create the inner JSON object
	inner := make(map[string]interface{})
	inner["username"] = body.Username
	inner["fullname"] = body.Fullname

	// Insert the inner JSON object into the outer JSON object
	outer[0] = inner

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": inner})
}

func Login(c *gin.Context) {
	//get the email and pass off req body

	var body struct {
		Fullname string `gorm:"varchar(100)" json:"fullname"`
		Username string `gorm:"unique" json:"username"`
		Password string `gorm:"varchar(50)" json:"password"`
	}

	if c.ShouldBindJSON(&body) != nil { //apabila body yang diberikan tidak mengembalikan apa apa
		messagebody := make(map[string]interface{})

		// Insert the inner JSON object into the outer JSON object
		messagebody["message"] = "Body not found or error, please try again"
		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messagebody, //memberikan pesan eror
		})

		return
	}

	//look up requested user
	var user models.User
	models.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		messagebody := make(map[string]interface{})

		// Insert the inner JSON object into the outer JSON object
		messagebody["message"] = "invalid username"

		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messagebody, //memberikan pesan eror
		})

		return
	}

	//compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) //melakukan compare antara hashing password dan password yang diberikan oleh post api

	if err != nil {
		messagebody := make(map[string]interface{})

		// Insert the inner JSON object into the outer JSON object
		messagebody["message"] = "invalid password"

		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messagebody, //memberikan pesan eror
		})

		return
	}

	//generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"role":     user.Role,
		"fullname": user.Fullname,
		"exp":      time.Now().Add(time.Hour * 24 * 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) //mengambil env secret

	if err != nil {
		messagebody := make(map[string]interface{})

		// Insert the inner JSON object into the outer JSON object
		messagebody["message"] = "Failed to create token"

		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messagebody, //memberikan pesan eror
		})

		return
	}

	//sent it back

	c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true) //menyimpan cookie authorization jwt
	messagelogin := make(map[string]interface{})

	// Insert the inner JSON object into the outer JSON object
	messagelogin["message"] = "Login Successful!"
	messagelogin["token"] = tokenString
	c.JSON(http.StatusOK, gin.H{

		"status": 1,
		"data":   messagelogin,
	})
}

func Logout(c *gin.Context) {
	// Clear the Authorization cookie
	// c.SetCookie("Authorization", "", -1, "/", "", false, true)

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data":   "Logged out successfully",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data":   user,
	})
}

func Showall(c *gin.Context) {

	var user []models.User //array dan ambil model product

	models.DB.Find(&user)
	// returndata := make(map[string]interface{})
	// returndata["username"] = user.username
	// returndata["username"] = user.username

	// Insert the inner JSON object into the oute
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": user}) //untuk return json nya
}

func CheckAuthHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [no header]"})
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
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [expired]"})

			return
		}

		var user models.User
		models.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [no user]"})

			return
		}

		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "data": "You Logged in"})

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token no valid]"})

	}
}
func CheckRoleHandler(c *gin.Context) {
	// Get Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [no header]"})
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

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	// Handle token parsing errors
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token invalid]"})
		return
	}

	// Token is valid, extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check for token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token expired]"})
			return
		}

		// Get the user from the database using the user ID from token claims
		var user models.User
		models.DB.First(&user, claims["sub"]) // Assuming "sub" is user ID in JWT claims

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [no user found]"})
			return
		}

		// Convert role to string and parse it to an integer if necessary
		userRoleStr := claims["role"].(string) // Assuming role is stored as string in JWT
		userRole, err := strconv.Atoi(userRoleStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "Invalid role format"})
			return
		}

		// Extract the username from claims
		username := claims["fullname"].(string) // Assuming username is stored in JWT claims

		// Check if the user has the required role (role 1, 2, or 3)
		if userRole == 1 || userRole == 2 || userRole == 3 {
			// Authorized, return success response
			c.JSON(http.StatusOK, gin.H{"status": 1, "data": "User is authorized", "role": userRole, "fullname": username})
		} else {
			// Unauthorized, role does not match
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": -1, "data": "You are unauthorized [insufficient role]"})
		}
	} else {
		// Token is not valid or claims could not be parsed
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": -1, "data": "You are unauthorized [token not valid]"})
	}
}
