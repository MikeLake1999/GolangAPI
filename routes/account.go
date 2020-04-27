package routes

import (
	"errors"
	"fmt"
	"gallery/services"

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

	account, err := services.Register(cred.Email, cred.Password, cred.Name, cred.Avatar, cred.Address, cred.Phone)
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
	services.Logger.Infof("Get account information by id=[%d]", accountId)
	account, err := services.GetAccountByID(accountId.(uint))

	newAccount := &Credential{}
	if err = ctx.BindJSON(&newAccount); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if newAccount.Name != account.Name {
		account.Name = newAccount.Name
	}

	if newAccount.Avatar != account.Avatar {
		account.Avatar = newAccount.Avatar
	}
	if newAccount.Email != account.Email {
		account.Email = newAccount.Email
	}
	if newAccount.Password != account.Password {
		account.Password = newAccount.Password
	}
	if newAccount.Address != account.Address {
		account.Address = newAccount.Address
	}
	if newAccount.Phone != account.Phone {
		account.Phone = newAccount.Phone
	}

	err = services.SaveAccount(account)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)

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
