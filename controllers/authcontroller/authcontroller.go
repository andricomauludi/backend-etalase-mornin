package authcontroller

import (
	"net/http"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email/ pass off req body
	var body models.User

	if c.ShouldBindJSON(&body) != nil { //apabila body yang diberikan tidak mengembalikan apa apa
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Body not found", //memberikan pesan eror
		})

		return
	}

	//Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Eror": "Failed to hash password", //memberikan pesan eror
		})

		return
	}

	//Create the user
	user := models.User{NamaLengkap: body.NamaLengkap, Username: body.Username, Password: string(hash)}
	result := models.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Eror": "Failed to create user, change ur username", //memberikan pesan eror
		})

		return
	}

	//respond

	c.JSON(http.StatusOK, gin.H{"Success": body})
}
