package routes

import (
	"errors"
	"fmt"
	"gallery/services"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GalleryId   int    ` json:"gallery_id"`
	Path        string ` json:"path"`
	Size        int64  ` json:"size"`
	Count       int    ` json:"count"`
}

func saveUploadedFile(c *gin.Context) (path string, size int64, err error) {

	file, err := c.FormFile("image")
	filename := filepath.Base(file.Filename)
	path = filename
	size = file.Size
	if err = c.SaveUploadedFile(file, path); err != nil {
		fmt.Println(err)
		return
	}
	return
}
func CreatePhoto(ctx *gin.Context) {
	photo := &PhotoType{}
	accountId, exist := ctx.Get("account_id")
	if !exist {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	galleryId := ctx.PostForm("gallery_id")

	id, err := strconv.ParseInt(galleryId, 10, 64)

	if err != nil {
		ctx.AbortWithError(400, errors.New("Can not convert gallery id"))
		return
	}

	photo.GalleryId = int(id)
	photo.Name = ctx.PostForm("name")
	photo.Description = ctx.PostForm("description")

	path, size, err := saveUploadedFile(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Can not save uploaded file"))
		return
	}

	photo.Path = path
	photo.Size = size

	photos, err := services.CreatePhoto(accountId.(uint), photo.GalleryId, photo.Name, photo.Description, photo.Path, photo.Size)
	services.Logger.Infof("Create photo by Account ID=[%d], Gallery Id=[%s], Photo Name=[%s], Description=[%s], Path=[%s], Size=[%d]", accountId, photo.GalleryId, photo.Name, photo.Description, photo.Path, photo.Size)
	fmt.Println(photos)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}
	fmt.Println(photos)
	ctx.JSON(200, photos)
}

func GetPhoto(ctx *gin.Context) {
	id := ctx.Param("id")

	photo, err := services.GetPhotoAndReaction(id)
	services.Logger.Infof("Get photo By Photo ID=[%d]", id)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, photo)
}
func UpdatePhoto(ctx *gin.Context) {
	id := ctx.Param("id")

	newPhoto := &PhotoType{}
	if err := ctx.BindJSON(&newPhoto); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	photo, err := services.UpdatePhoto(id, newPhoto.Name, newPhoto.Description)
	services.Logger.Infof("Update Photo ID=[%d], Name=[%s], Description=[%s]", id, newPhoto.Name, newPhoto.Description)
	fmt.Println(photo)
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
	services.Logger.Infof("Delete Photo ID=[%d]", id)
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

	reactive, err := services.CreateReaction(accountId.(uint), photo.Id)
	services.Logger.Infof("Reaction Photo ID=[%d] By Account ID=[%d]", photo.Id, accountId)
	fmt.Println(reactive)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}

	ctx.JSON(200, photo)
}
func DeleteReaction(ctx *gin.Context) {
	id := ctx.Param("id")

	err1 := services.DeleteReaction(id)
	services.Logger.Infof("Unreaction Photo ID=[%d]", id)
	if err1 != nil {
		ctx.AbortWithError(404, errors.New("Photo Not Found!"))
		return

	}
	ctx.Status(200)
}
