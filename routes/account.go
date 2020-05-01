package routes

import (
	"errors"
	"fmt"
	"gallery/services"
	"sync"

	"github.com/gin-gonic/gin"
)

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
type PasswordHash struct {
	Password string `json:"password"`
}

func Authentication(ctx *gin.Context) {
	cred := &Credential{}
	if err := ctx.BindJSON(cred); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(401, errors.New("Invalid email or Password"))
		return
	}
	services.Logger.Infof("Authentication with email=[%s], Password=[%s]", cred.Email, cred.Password)
	token, err := services.Authenticate(cred.Email, cred.Password)
	if err != nil {
		ctx.AbortWithError(401, errors.New("Invalid email or password"))
		return
	}
	ctx.String(200, token)
}
func Registration(ctx *gin.Context) {
	cred := &Credential{}
	if err := ctx.BindJSON(cred); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}
	services.Logger.Infof("Registration account on Email=[%s], Password=[%s], Name=[%s], Address=[%s], Phone=[%s]", cred.Email, cred.Password, cred.Name, cred.Address, cred.Phone)
	account, err := services.Register(cred.Email, cred.Password, cred.Name, cred.Address, cred.Phone)
	fmt.Println(account)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}
	ctx.Status(200)
}
func GetAccount(ctx *gin.Context) {
	accountId, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	services.Logger.Infof("Get account information by id=[%d]", accountId)
	account, err := services.GetAccountByID(accountId.(uint))

	if err != nil {
		ctx.AbortWithError(404, errors.New("Account Not Found"))
		return
	}
	ctx.JSON(200, account)
}
func UpdateAccount(ctx *gin.Context) {

	// get id param
	accountId, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	cred := &Credential{}
	if err := ctx.BindJSON(cred); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}

	account, err := services.UpdateAccount(cred.Email, cred.Name, cred.Address, cred.Phone, accountId.(uint))
	services.Logger.Infof("Update account Id=[%d], Email=[%s], Name=[%s], Address=[%s], Phone=[%s] ", accountId, cred.Email, cred.Name, cred.Address, cred.Phone)
	fmt.Println(account)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)

}
func UpdateAvatar(ctx *gin.Context) {
	cred := &Credential{}
	accountId, exist := ctx.Get("account_id")
	if !exist {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	path, size, err := saveUploadedFile(ctx)

	fmt.Println(size)
	cred.Avatar = path

	image := services.ImageSize{
		Path: path,
	}
	w, h, err := services.GetDimension(image.Path)

	fmt.Printf("Width: %d, Height: %d\n", w, h)

	resolutions := []int{64}
	wg := sync.WaitGroup{}
	wg.Add(len(resolutions))
	for _, size := range resolutions {
		go func(wg *sync.WaitGroup, size int) {
			defer wg.Done()
			_, err := image.Resize(size)
			fmt.Println(err)
		}(&wg, size)

	}
	wg.Wait()

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Can not save uploaded file"))
		return
	}

	account, err := services.UpdateAvatar(cred.Avatar, accountId.(uint))
	services.Logger.Infof("Update avatar by account Id=[%d], Avatar=[%s]", accountId, cred.Avatar)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Error"))
		return
	}

	ctx.JSON(200, account)
}
func UpdatePassword(ctx *gin.Context) {

	// get id param
	accountId, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	password := ctx.Param("password")

	account, err := services.UpdatePassword(password, accountId.(uint))
	services.Logger.Infof("Update password by account Id=[%d], Password=[%s]", accountId, password)
	fmt.Println(account)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)

}

func DeleteAccount(ctx *gin.Context) {
	accountId, exist := ctx.Get("account_id")
	if !exist {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	services.Logger.Info("Delete account by id=[%d]", accountId)
	err := services.DeleteAccount(accountId.(uint))
	if err != nil {
		ctx.AbortWithError(404, errors.New("Account Not Found!"))
		return
	}
	ctx.Status(200)
}
