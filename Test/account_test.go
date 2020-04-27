package test

import (
	"gallery/routes"
	"gallery/services"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var ts *gin.Engine

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	_ = services.ConnectDB("root:Hieu2951999@tcp(127.0.0.1:3306)/Galleries?parseTime=true")
	ts = routes.Create()
}

func tearDown() {
	defer services.DB.Close()
}

func TestGetAccountPublic(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/v1/public/account/4", nil)
	ts.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code should be Ok, was: %d", writer.Code)
	}
	if writer.HeaderMap.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Errorf(
			"Content-Type should be application/json, was %s",
			writer.HeaderMap.Get("Content-Type"),
		)
	}
}

// func TestCreateEntry(t *testing.T) {

// 	account := models.Account{
// 		Email:    "nopexx",
// 		Password: "123",
// 		Name:     "hieu",
// 		Avatar:   "none",
// 		Address:  "none",
// 		Phone:    "0123456789",
// 	}

// 	req, err := http.NewRequest("POST", "/registration", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler, err := http.HandlerFunc(services.Register(account.Email, account.Password, account.Name, account.Avatar, account.Address, account.Phone))
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// }
