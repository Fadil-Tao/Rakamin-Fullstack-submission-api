package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rakamins-pbi/final-task-pbi-rakamin-fullstack-HadadFadilah/Internals/models"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(ctx *gin.Context) {
	// bind data masuk dengan struct
	var data struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	err := ctx.ShouldBind(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dasar Objek user
	var user models.User
	// Query ke database mencari email
	models.DB.First(&user, "email = ?", data.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid email",
		})
		return
	}
	// Pencocokan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	// Membuat jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Mengirim kembali token menggunakan http only cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*7, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"massage": "successfull"})
}

func HandleRegister(ctx *gin.Context) {
	// bind data masuk dengan struct
	var data struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := ctx.ShouldBind(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check apakah user name atau email sudah terpakai
	if !models.UserAvailable(data.Email, data.Username) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email or username alreadey taken"})
		return
	}
	// Hashing password
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Membuat user dan menyimpan ke database
	user, err := models.CreateUser(data.Email, data.Username, string(hashedpassword))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error failed to create user"})
	}
	// Pesan sukses dengan json
	ctx.JSON(http.StatusOK, gin.H{"massage": "user successfully created", "user": user.Email})
}

func GetUserDetails(ctx *gin.Context) {
	Param := ctx.Param("userid")

	userId, err := strconv.ParseUint(Param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := models.DB.First(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := models.DB.Model(&user).Preload("Photos").Find(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"photos":   user.Photos,
	})
}

func UpdateUser(ctx *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !models.UserAvailable(body.Email, body.Username) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email or username alreadey taken"})
		return
	}

	Param := ctx.Param("userid")
	userId, err := strconv.ParseUint(Param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	value, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userid := value.(uint)

	var user models.User
	if err := models.DB.First(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.ID != userid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You cant edit other's profile"})
		return
	}
	
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Email = body.Email
	user.Username = body.Username
	user.Password = string(hashedpassword)

	if err := models.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Errorr"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage":  "The User has been Updated successfully",
		"Email":    user.Email,
		"Username":  user.Username,
	})
}


func DeleteUser(ctx *gin.Context){
	Param := ctx.Param("userid")
	userId, err := strconv.ParseUint(Param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Photo ID"})
		return
	}
	value, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userid := value.(uint)
	
	var user models.User 

	if err := models.DB.First(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cant Found"})
		return
	}
	if user.ID != userid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized you cant delete other's profile"})
		return
	}

	if err := models.DB.Unscoped().Delete(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Errorr"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "The Account has been deleted successfully"})
}


func ClearCookie(ctx *gin.Context){
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK,gin.H{"massage":"Cookie Has Been Deleted"})
}