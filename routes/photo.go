package routes

import (
	"errors"
	"fmt"
	"gallery/services"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type PhotoType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type ReactionType struct {
	Reaction string `json:"reaction"`
}

func saveUploadedFileToDirectory(c *gin.Context) (path string, size int64, err error) {

	file, err := c.FormFile("image")
	filename := filepath.Base(file.Filename)
	path = "image/" + filename
	size = file.Size
	if err = c.SaveUploadedFile(file, path); err != nil {
		fmt.Println(err)
		return
	}
	return
}
func CreatePhoto(ctx *gin.Context) {

	accountId, exist := ctx.Get("account_id")
	if !exist {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	id := ctx.Param("id")
	gallery, err := services.GetGallery(id)
	fmt.Println(gallery.Id)
	path, size, err := saveUploadedFileToDirectory(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Can not save uploaded file"))
		return
	}
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	photo := &PhotoType{}
	if err := ctx.BindJSON(photo); err != nil {
		ctx.AbortWithError(400, err)
	}
	photos, err := services.CreatePhoto(accountId.(uint), gallery.Id, photo.Name, photo.Description, path, size)
	fmt.Println(photos)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}
	fmt.Println(gallery)
	ctx.JSON(200, gallery)
}

func GetPhoto(ctx *gin.Context) {
	id := ctx.Param("id")

	photo, err := services.GetPhoto(id)
	fmt.Println(photo.Id)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, photo)
}
func UpdatePhoto(ctx *gin.Context) {
	id := ctx.Param("id")

	photo, err := services.GetPhoto(id)
	fmt.Println(photo)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	newPhoto := &PhotoType{}
	if err = ctx.BindJSON(&newPhoto); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if photo.Name != newPhoto.Name {
		photo.Name = newPhoto.Name
	}
	if photo.Description != newPhoto.Description {
		photo.Description = newPhoto.Description
	}

	err = services.SavePhoto(photo)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return

}
func DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("id")

	photo, err := services.GetPhoto(id)
	fmt.Println(photo)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	err1 := services.DeletePhoto(id)
	if err1 != nil {
		ctx.AbortWithError(404, errors.New("Photo Not Found!"))
		return

	}
	ctx.Status(200)
}
func CreateReaction(ctx *gin.Context) {
	accountId, exist := ctx.Get("account_id")
	if !exist {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	id := ctx.Param("id")
	photo, err := services.GetPhoto(id)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	reaction := &ReactionType{}
	if err := ctx.BindJSON(reaction); err != nil {
		ctx.AbortWithError(400, err)
	}

	reactive, err := services.CreateReaction(accountId.(uint), photo.Id, reaction.Reaction)
	fmt.Println(reactive)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}

	ctx.JSON(200, reaction)
}
func DeleteReaction(ctx *gin.Context) {
	id := ctx.Param("id")

	err1 := services.DeleteReaction(id)
	if err1 != nil {
		ctx.AbortWithError(404, errors.New("Photo Not Found!"))
		return

	}
	ctx.Status(200)
}
