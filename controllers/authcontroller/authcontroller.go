package authcontroller

import (
	"net/http"
	"os"
	"time"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email/ pass off req body
	var body models.User

	messagebody := make(map[string]interface{})

	// Insert the inner JSON object into the outer JSON object
	messagebody["message"] = "Body not found or error, please try again"

	if c.ShouldBindJSON(&body) != nil { //apabila body yang diberikan tidak mengembalikan apa apa
		c.JSON(http.StatusBadRequest, gin.H{"status": -1,
			"data": messagebody, //memberikan pesan eror
		})

		return
	}

	//Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

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
	user := models.User{NamaLengkap: body.NamaLengkap, Username: body.Username, Password: string(hash)}
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

	// var isian = map[string]string{"username": body.Username, "Nama Lengkap": body.NamaLengkap}
	// var responses = map[string]string{"message": "Account Created!"}
	// responses["data"] = isian

	outer := make(map[int]interface{})

	// Create the inner JSON object
	inner := make(map[string]interface{})
	inner["username"] = body.Username
	inner["nama_lengkap"] = body.NamaLengkap

	// Insert the inner JSON object into the outer JSON object
	outer[0] = inner

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": outer})
}

func Login(c *gin.Context) {
	//get the email and pass off req body

	var body struct {
		NamaLengkap string `gorm:"varchar(100)" json:"nama_lengkap"`
		Username    string `gorm:"unique" json:"username"`
		Password    string `gorm:"varchar(50)" json:"password"`
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
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

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
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

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
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data":   "Login Successful! ",
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data":   "Logged in",
	})
}
