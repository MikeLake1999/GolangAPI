package routes

import (
	"errors"
	"fmt"
	"gallery/services"

	"github.com/gin-gonic/gin"
)

func GetPublicGallery(ctx *gin.Context) {
	id := ctx.Param("id")

	gallery, err := services.GetPhotosPublicGallery(id)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, gallery)
}
func GetPublicPhoto(ctx *gin.Context) {
	id := ctx.Param("id")

	photo, err := services.GetGalleryPublicPhoto(id)
	fmt.Println(photo)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, photo)
}
func GetPublicAccount(ctx *gin.Context) {
	id := ctx.Param("id")

	account, err := services.GetPublicAccount(id)
	fmt.Println(account)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, account)
}
func GetPublicGalleries(ctx *gin.Context) {
	gallery, err := services.GetPublicGalleries()
	fmt.Println(gallery)
	if err != nil {
		ctx.AbortWithError(404, errors.New("Not Found"))
		return
	}

	ctx.JSON(200, gallery)
}
