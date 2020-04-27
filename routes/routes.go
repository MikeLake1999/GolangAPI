package routes

import (
	"gallery/middleware"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Create() (g *gin.Engine) {

	g = gin.Default()

	v1 := g.Group("/v1")
	{
		v1.POST("/registration", Registration)
		v1.POST("/authentication", Authentication)

		upload := v1.Group("/upload")
		{
			upload.POST("/single", func(ctx *gin.Context) {
				file, err := ctx.FormFile("file")
				if err != nil {
					ctx.AbortWithError(400, err)
					return
				}

				filename := filepath.Base(file.Filename)
				if err := ctx.SaveUploadedFile(file, filename); err != nil {
					ctx.AbortWithError(400, err)
					return
				}
				ctx.Status(200)
			})
		}
		account := v1.Group("/account")
		{
			// Use authentication middleware
			account.Use(middleware.RequireAuthentication())
			// add account handlers
			account.GET("", GetAccount)
			account.PUT("", UpdateAccount)
			account.PUT("/password", UpdatePassword)
			account.DELETE("", DeleteAccount)
		}

		gallery := v1.Group("/gallery")
		{
			// Use authentication middleware
			gallery.Use(middleware.RequireAuthentication())
			// add gallery handlers
			gallery.POST("", CreateGallery)
			gallery.GET("", GetAllGalleries)
			gallery.GET("/:id", GetGallery)
			gallery.PUT("/:id", UpdateGallery)
			gallery.DELETE("/:id", DeleteGallery)
		}
		photo := v1.Group("/photo")
		{
			// Use authentication middleware
			photo.Use(middleware.RequireAuthentication())
			// add photo handlers
			photo.POST("", CreatePhoto)
			photo.GET("/:id", GetPhoto)
			photo.PUT("/:id", UpdatePhoto)
			photo.DELETE("/:id", DeletePhoto)
			photo.POST("/:id/reaction", CreateReaction)
			photo.DELETE("/:id/reaction", DeleteReaction)
		}
		public := v1.Group("/public")
		{
			// add public handlers

			public.GET("/gallery", GetPublicGalleries)
			public.GET("/gallery/:id", GetPublicGallery)
			public.GET("/photo/:id", GetPublicPhoto)
			public.GET("/account/:id", GetPublicAccount)

		}

	}

	return
}
