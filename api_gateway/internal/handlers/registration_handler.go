package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rosariocannavo/api_gateway/internal/db"
	"github.com/rosariocannavo/api_gateway/internal/models"
	"github.com/rosariocannavo/api_gateway/internal/nats"
	"github.com/rosariocannavo/api_gateway/internal/repositories"
	"github.com/rosariocannavo/api_gateway/internal/utils"
)

func HandleRegistration(c *gin.Context) {
	userRepo := repositories.NewUserRepository(db.Client)

	var userForm models.UserForm

	//retrieve the partial user information from form
	if err := c.BindJSON(&userForm); err != nil {

		message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "registration_handler/HandleRegistration", http.StatusBadRequest, "error: Invalid request payload")
		nats.NatsConnection.PublishMessage(message)

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	//check if the user is already registered
	isPresent, err := userRepo.CheckIfUserIsPresent(userForm.Username, userForm.MetamaskAddress)

	if err != nil {

		message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "registration_handler/HandleRegistration", http.StatusInternalServerError, "error: Error fetching database")
		nats.NatsConnection.PublishMessage(message)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching database"})
		return
	}

	if isPresent {

		message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "registration_handler/HandleRegistration", http.StatusForbidden, "error: User already present")
		nats.NatsConnection.PublishMessage(message)

		c.JSON(http.StatusForbidden, gin.H{"error": "User already present"})
		return

	} else {

		//if user is not present
		//hash his psw and store him in the db
		//give him a role based on # of transaction
		//give him a nonce

		var user models.User

		// hash the user password
		hashedPwd, err := utils.HashPassword(userForm.Password)

		if err != nil {
			message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "registration_handler/HandleRegistration", http.StatusInternalServerError, "error: Error hashing password")
			nats.NatsConnection.PublishMessage(message)

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
			return
		}

		// generate nonce for metamask sign auth
		nonce, err := utils.GenerateRandomNonce()

		if err != nil {
			message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "registration_handler/HandleRegistration", http.StatusInternalServerError, "error: Bad nonce generation")
			nats.NatsConnection.PublishMessage(message)

			c.JSON(http.StatusInternalServerError, gin.H{"error": " Bad nonce generation"})
			return
		}

		user.Username = userForm.Username
		user.Password = hashedPwd
		user.MetamaskAddress = userForm.MetamaskAddress
		user.Nonce = nonce
		user.Role = utils.CheckUserBalance(user.MetamaskAddress)

		// write the user in the database
		userRepo.CreateUser(&user)

		message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s %s, role: %s", time.Now().UTC().Format(time.RFC3339), "registration_handler/HandleRegistration", http.StatusOK, "message: User registered succesfully. username: ", user.Username, user.Role)
		nats.NatsConnection.PublishMessage(message)

		c.JSON(http.StatusOK, gin.H{"message": "User registered succesfully"})
	}
}
