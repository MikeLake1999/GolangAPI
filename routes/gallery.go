package routes

import (
	"errors"
	"fmt"
	"gallery/services"

	"github.com/gin-gonic/gin"
)

type Galla struct {
	Name   string `json:"name"`
	Brief  string `json:"brief"`
	Active string `json:"active"`
}

func CreateGallery(ctx *gin.Context) {
	accountId, exist := ctx.Get("account_id")
	if !exist {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	gallary := &Galla{}
	if err := ctx.BindJSON(gallary); err != nil {
		ctx.AbortWithError(400, err)
	}
	gallery, err := services.CreateGallery(gallary.Name, gallary.Brief, accountId.(uint))
	services.Logger.Infof("Create gallery by Account ID=[%d], Gallery Name=[%s], Brief=[%s], ", accountId, gallary.Name, gallary.Brief)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}
	fmt.Println(gallery)
	ctx.JSON(200, gallery)
}
func GetAllGalleries(ctx *gin.Context) {
	accountId, exist := ctx.Get("account_id")
	if !exist {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	galleries, err := services.GetAllGalleries(accountId.(uint))
	services.Logger.Infof("Get all galleries By Account ID=[%d]", accountId)
	fmt.Println(galleries)
	if err != nil {
		ctx.AbortWithError(404, errors.New("Not Found"))
		return
	}

	ctx.JSON(200, galleries)
}
func GetGallery(ctx *gin.Context) {
	id := ctx.Param("id")

	gallery, err := services.GetGallery(id)
	services.Logger.Infof("Get gallery by Gallery ID=[%d]", id)
	fmt.Println(gallery)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, gallery)
}
func GetPhotoOfGallery(ctx *gin.Context) {
	id := ctx.Param("id")

	gallery, err := services.GetPhotosGallery(id)
	services.Logger.Infof("Get Photo of Gallery ID=[%d]", id)
	fmt.Println(gallery)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, gallery)
}
func Publication(ctx *gin.Context) {
	id := ctx.Param("id")

	gallery, err := services.Publication(id)
	services.Logger.Infof("Public with Gallery ID=[%d]", id)
	fmt.Println(gallery)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, gallery)
}
func UpdateGallery(ctx *gin.Context) {
	id := ctx.Param("id")

	newGallery := &Galla{}
	if err := ctx.BindJSON(&newGallery); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	galleries, err := services.UpdateGallery(id, newGallery.Name, newGallery.Brief)
	services.Logger.Infof("Update Gallery ID=[%d], Gallery Name=[%s], Brief=[%s]", id, newGallery.Name, newGallery.Brief)
	fmt.Println(galleries)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}
func DeleteGallery(ctx *gin.Context) {
	id := ctx.Param("id")

	gallery, err := services.GetGallery(id)
	services.Logger.Infof("Delete Galleries ID=[%d]", id)
	fmt.Println(gallery)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	err1 := services.DeleteGallery(id)
	if err1 != nil {
		ctx.AbortWithError(404, errors.New("Gallery Not Found!"))
		return

	}
	ctx.Status(200)
}
