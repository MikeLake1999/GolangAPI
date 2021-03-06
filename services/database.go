package services

import (
	"errors"
	"fmt"
	"gallery/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func ConnectDB(connection string) (err error) {

	// Create connection
	DB, err = gorm.Open("mysql", connection)
	if err != nil {
		return
	}

	err = DB.AutoMigrate(
		&models.Account{},
		&models.Galleries{},
		&models.Photos{},
		&models.Reactions{},
	).Error
	if err != nil {
		return
	}

	DB.Model(&models.Galleries{}).AddForeignKey(
		"account_id", "accounts(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Photos{}).AddForeignKey(
		"account_id", "accounts(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Photos{}).AddForeignKey(
		"gallery_id", "galleries(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Reactions{}).AddForeignKey(
		"account_id", "accounts(id)", "CASCADE", "CASCADE",
	)
	DB.Model(&models.Reactions{}).AddForeignKey(
		"photo_id", "photos(id)", "CASCADE", "CASCADE",
	)

	return
}

// Account Query
func GetAccountByID(id uint) (account *models.Account, err error) {
	Logger.Debugf("Get account information by id=[%d]", id)
	account = &models.Account{}
	err = DB.First(account, id).Error
	return
}

func GetAccountByEmail(email string) (account *models.Account, err error) {
	Logger.Debugf("Get account email=[%s]", email)
	account = &models.Account{}
	err = DB.Where("email = ?", email).First(account).Error
	return
}
func GetPublicAccount(id string) (account *models.Account, err error) {
	Logger.Debugf("Get public account id=[%d]", id)
	account = &models.Account{}
	err = DB.Select("name, avatar").Where("id = ?", id).Find(&account).Error
	return

}

func SaveAccount(account *models.Account) (err error) {
	err = DB.Save(account).Error
	return
}
func Authenticate(email string, password string) (tokenStr string, err error) {
	Logger.Debugf("Authentication email=[%s], password=[%s]", email, password)
	account, err := GetAccountByEmail(email)
	if err != nil {
		return

	}
	fmt.Println(account)
	bytePassword := []byte(password)

	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	account.Password = string(passwordHash)
	if account.Password != string(passwordHash) {
		return "", errors.New("Invalid Email or Password!")
	}
	tokenStr, err = CreateToken(account.Id)
	return
}
func Register(email string, password string, name string, address string, phone string) (account models.Account, err error) {
	Logger.Debugf("Registration email=[%s], password=[%s], name=[%s], address=[%s], phone=[%s]", email, password, name, address, phone)
	bytePassword := []byte(password)

	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	account.Password = string(passwordHash)
	account = models.Account{
		Email:    email,
		Password: string(passwordHash),
		Name:     name,
		Address:  address,
		Phone:    phone,
	}
	err = DB.Create(&account).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}
func UpdateAccount(email string, name string, address string, phone string, id uint) (account models.Account, err error) {
	Logger.Debugf("Update account information by id=[%d], email=[%s], name=[%s], address=[%s], phone=[%s]", id, email, name, address, phone)
	account = models.Account{
		Email:   email,
		Name:    name,
		Address: address,
		Phone:   phone,
	}
	err = DB.Model(&account).Where("id = ?", id).Updates(&account).Error
	if err != nil {
		fmt.Println(err)
	}
	return

}
func UpdateAvatar(avatar string, id uint) (account models.Account, err error) {
	Logger.Debugf("Update avatar by account id=[%d], avatar=[%s]", id, avatar)
	account = models.Account{
		Avatar: avatar,
	}
	err = DB.Model(&account).Where("id = ?", id).Updates(&account).Error
	if err != nil {
		fmt.Println(err)
	}
	return

}
func UpdatePassword(password string, id uint) (account models.Account, err error) {
	Logger.Debugf("Update password account  id=[%d], password=[%s]", id, password)
	bytePassword := []byte(password)

	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	account.Password = string(passwordHash)

	err = DB.Model(&account).Where("id = ?", id).Update("password", string(passwordHash)).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteAccount(id uint) (err error) {
	Logger.Debugf("Delete account by ID=[%d]", id)

	account := &models.Account{}
	err = DB.First(account, id).Error
	if err != nil {
		return
	}
	err = DB.Delete(account).Error
	return
}

// Gallery Query
func CreateGallery(name string, brief string, id uint) (gallery models.Galleries, err error) {
	Logger.Debugf("Create gallery by account id =[%d], gallery name=[%s], brief=[%s]", id, name, brief)
	gallery = models.Galleries{
		Name:      name,
		Brief:     brief,
		AccountId: id,
		Active:    "Inactive",
	}
	err = DB.Create(&gallery).Error
	if err != nil {
		fmt.Println(err)
	}
	return

}
func SaveGallery(gallery *models.Galleries) (err error) {
	err = DB.Save(gallery).Error
	return
}
func GetGallery(id string) (gallery *models.Galleries, err error) {
	Logger.Debugf("Get gallery id=[%d]", id)
	gallery = &models.Galleries{}
	err = DB.Where("id = ?", id).First(gallery).Error
	return
}
func GetAllGalleries(accountId uint) (galleries []models.Galleries, err error) {
	Logger.Debugf("Get all galleries by account id=[%d]", accountId)
	galleries = []models.Galleries{}
	err = DB.Where("account_id = ?", accountId).Find(&galleries).Error
	fmt.Println(len(galleries))
	for i := 0; i < len(galleries); i++ {
		galleries[i].Photos = []models.Photos{}
		err = DB.Model(galleries[i]).Limit(3).Related(&(galleries[i].Photos), "Photos").Error
		if err != nil {
			return
		}

	}
	fmt.Println(galleries)
	return
}
func Publication(id string) (galleries models.Galleries, err error) {
	Logger.Debugf("Publication gallery id=[%d]", id)
	galleries = models.Galleries{
		Active: "active",
	}
	err = DB.Model(&galleries).Where("galleries.id = ?", id).Updates(&galleries).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}
func UpdateGallery(id string, name string, brief string) (galleries models.Galleries, err error) {
	Logger.Debugf("Update gallery id=[%d], name=[%s], brief=[%s]", id, name, brief)
	galleries = models.Galleries{
		Name:  name,
		Brief: brief,
	}
	err = DB.Model(&galleries).Where("galleries.id = ?", id).Updates(&galleries).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetPublicGalleries() (galleries []models.Galleries, err error) {
	galleries = []models.Galleries{}
	err = DB.Where("active = active").Find(&galleries).Error
	if err != nil {
		return
	}
	fmt.Println(len(galleries))
	for i := 0; i < len(galleries); i++ {
		galleries[i].Photos = []models.Photos{}
		err = DB.Model(galleries[i]).Limit(3).Related(&(galleries[i].Photos), "Photos").Error

		if err != nil {
			return
		}

	}
	fmt.Println(galleries)
	return
}
func GetPhotosGallery(id string) (photos []models.Photos, err error) {
	Logger.Debugf("Get photo by gallery id=[%d]", id)
	photos = []models.Photos{}

	err = DB.Select("photos.id, photos.gallery_id, photos.name, galleries.active, photos.account_id, photos.description, photos.path, photos.size, photos.sum_reaction").Joins("join galleries ON galleries.id = ?  AND galleries.id = photos.gallery_id", id).Find(&photos).Error
	fmt.Println(len(photos))
	for i := 0; i < len(photos); i++ {
		err = DB.Model(photos[i]).Related(&(photos[i].Reactions), "Reactions").Error
		photos[i].SumReaction = len(photos[i].Reactions)
		fmt.Println(err)
		if err != nil {
			return
		}
	}
	fmt.Println(photos)
	return
}

func GetPhotosPublicGallery(id string) (photos []models.Photos, err error) {
	Logger.Debugf("Get photo by public gallery id=[%d]", id)

	photos = []models.Photos{}

	err = DB.Select("photos.id, photos.gallery_id, photos.name, galleries.active, photos.account_id, photos.description, photos.path, photos.size, photos.sum_reaction").Joins("join galleries ON galleries.id = ? AND galleries.active = 'active' AND galleries.id = photos.gallery_id", id).Find(&photos).Error
	fmt.Println(len(photos))
	for i := 0; i < len(photos); i++ {
		err = DB.Model(photos[i]).Related(&(photos[i].Reactions), "Reactions").Error
		photos[i].SumReaction = len(photos[i].Reactions)
		println(len(photos[i].Reactions))
		if err != nil {
			return
		}

	}
	fmt.Println(photos)
	return
}
func DeleteGallery(id string) (err error) {
	Logger.Debugf("Delete gallery id=[%d]", id)
	Logger.Debugf("Delete gallery by ID=[%d], id")

	gallery := &models.Galleries{}
	err = DB.First(gallery, id).Error
	if err != nil {
		return
	}
	err = DB.Delete(gallery).Error
	return
}

// Photo Query

func CreatePhoto(accountId uint, gallery_id int, name string, description string, path string, size int64) (photo models.Photos, err error) {
	Logger.Debugf("Create photo by account id=[%d], gallery id=[%d], photo name=[%s], descriprion=[%s], path=[%s], size=[%d]", accountId, gallery_id, name, description, path, size)
	photo = models.Photos{
		AccountId:   accountId,
		GalleryId:   gallery_id,
		Name:        name,
		Description: description,
		Path:        path,
		Size:        size,
	}
	err = DB.Create(&photo).Error
	if err != nil {
		fmt.Println(err)
	}
	return

}
func UpdatePhoto(id string, name string, description string) (photo models.Photos, err error) {
	Logger.Debugf("Update photo id=[%d], photo name=[%s], descriprion=[%s]", id, name, description)
	photo = models.Photos{
		Name:        name,
		Description: description,
	}
	err = DB.Model(&photo).Where("photos.id = ?", id).Updates(&photo).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetGalleryId(id int) (photo *[]models.Photos, err error) {

	photo = new([]models.Photos)
	err = DB.Where("gallery_id = ?", id).Find(photo).Error
	return

}
func GetGalleryPublicPhoto(id string) (photos []models.Photos, err error) {
	Logger.Debugf("Get public photo id=[%d]", id)

	photos = []models.Photos{}

	err = DB.Select("photos.id, photos.gallery_id, photos.name, galleries.active, photos.account_id, photos.description, photos.path, photos.size, photos.sum_reaction").Joins("join galleries ON galleries.active = 'active' AND galleries.id = photos.gallery_id").Where("photos.id = ?", id).Find(&photos).Error
	for i := 0; i < len(photos); i++ {
		photos[i].Reactions = []models.Reactions{}

		err = DB.Model(photos[i]).Related(&(photos[i].Reactions), "Reactions").Error
		photos[i].SumReaction = len(photos[i].Reactions)
		if err != nil {
			return
		}

	}
	fmt.Println(photos)
	return
}
func SavePhoto(photo *models.Photos) (err error) {
	err = DB.Save(photo).Error
	return
}
func GetPhotoAndReaction(id string) (photos []models.Photos, err error) {
	Logger.Debugf("Get reaction of photo id=[%d]", id)
	photos = []models.Photos{}

	err = DB.Where("id = ?", id).Find(&photos).Error
	fmt.Println(len(photos))
	for i := 0; i < len(photos); i++ {

		err = DB.Model(photos[i]).Related(&(photos[i].Reactions), "Reactions").Error
		fmt.Println(err)
		if err != nil {
			return
		}

	}
	fmt.Println(photos)
	return
}
func GetPhoto(id string) (photo *models.Photos, err error) {
	photo = &models.Photos{}
	err = DB.Where("id = ?", id).First(photo).Error
	return
}
func GetPublicPhoto(id int) (photo *[]models.Photos, err error) {
	photo = new([]models.Photos)
	err = DB.Where("gallery_id = ?", id).Find(photo).Error
	return
}
func DeletePhoto(id string) (err error) {
	Logger.Debugf("Delete photo by ID=[%d], id")

	photo := &models.Photos{}
	err = DB.First(photo, id).Error
	if err != nil {
		return
	}
	err = DB.Delete(photo).Error
	return
}

func CreateReaction(account_id uint, photo_id int) (reactive models.Reactions, err error) {
	Logger.Debugf("Create reaction by account id=[%d], photo id=[%d]", account_id, photo_id)
	reactive = models.Reactions{
		AccountId: account_id,
		PhotoId:   photo_id,
	}

	err = DB.Create(&reactive).Error
	if err != nil {
		fmt.Println(err)
	}
	return

}
func GetReactionByAccountId(id int) (reaction *models.Reactions, err error) {
	reaction = &models.Reactions{}
	err = DB.Where("photo_id = ?", id).First(reaction).Error
	return
}
func DeleteReaction(id string) (err error) {
	Logger.Debugf("Delete reaction by photo id=[%d], id")

	reaction := &models.Reactions{}

	err = DB.Where("photo_id = ?", id).Delete(reaction).Error
	return
}
