package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rakamins-pbi/final-task-pbi-rakamin-fullstack-HadadFadilah/Internals/models"
)

func AddPhoto(ctx *gin.Context) {
	var PhotoData struct {
		Title    string `json:"title" binding:"required"`
		Caption  string `json:"caption" binding:"required"`
		PhotoUrl string `json:"url" binding:"required"`
	}

	err := ctx.ShouldBind(&PhotoData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userid := value.(uint)

	photo, err := models.CreatePhoto(PhotoData.Title, PhotoData.Caption, PhotoData.PhotoUrl, uint(userid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage":         "Photo successfully created",
		"createdByUserId": photo.UserID,
	})

}

func DeletePhoto(ctx *gin.Context) {
	Param := ctx.Param("photoid")

	photoId, err := strconv.ParseUint(Param, 10, 64)
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

	var photo models.Photos

	if err := models.DB.First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cant Found"})
		return
	}
	if photo.UserID != userid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized this image is not yours lmao"})
		return
	}

	if err := models.DB.Unscoped().Delete(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Errorr"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "The photo has been deleted successfully"})
}

func UpdatePhoto(ctx *gin.Context) {
	var body struct {
		Title    string `json:"title" binding:"required"`
		Caption  string `json:"caption" binding:"required"`
		PhotoUrl string `json:"url" binding:"required"`
	}

	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	Param := ctx.Param("photoid")
	photoId, err := strconv.ParseUint(Param, 10, 64)
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

	var photo models.Photos

	if err := models.DB.First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	if photo.UserID != userid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized this image is not yours lmao"})
		return
	}

	photo.Caption = body.Caption
	photo.Title = body.Title
	photo.PhotoUrl = body.PhotoUrl
	if err := models.DB.Save(&photo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Errorr"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage":  "The photo has been Updated successfully",
		"title":    photo.Title,
		"caption":  photo.Caption,
		"photoUrl": photo.PhotoUrl,
	})
}
