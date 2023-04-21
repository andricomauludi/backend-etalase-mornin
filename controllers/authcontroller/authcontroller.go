package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, l *http.Request) {

}
func Register(w http.ResponseWriter, l *http.Request) { //mengambil input json yg diterima

	var userInput models.User
	decoder := json.NewDecoder(l.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Failed to decode json")
	}
	defer l.Body.Close()

	//hash password menggunakan bcrypt

	//membuat hashpassword dengan input password dari user dan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	// memasukkan password yg diinput ke dalam models
	userInput.Password = string(hashPassword)

	//insert ke db
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("Failed to save data")
	}

	//respon json
	response, _ := json.Marshal(map[string]string{"message": "success to insert"})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
func Logout(w http.ResponseWriter, l *http.Request) {

}

//test
