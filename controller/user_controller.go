package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/valentinoanugerah/crud_go/database"
	"github.com/valentinoanugerah/crud_go/models"
	"golang.org/x/crypto/bcrypt"
	"os"
)



func GetUsers(c *gin.Context){
	var users  []models.User

	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func RegisterUser(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	 user.RoleID = 2 

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password) , bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash passowrd"})
		return
	}

	user.Password = string(hashedPassword)
	
	result := database.DB.Create(&user)
 if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}

type LoginInput struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))
func Login(c *gin.Context) {
	var input LoginInput
	var user models.User

	// 1. Ambil data JSON dari request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Cari user di database berdasarkan email
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		 "role_id": user.RoleID,
		"exp": time.Now().Add(time.Hour *1).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}



	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}