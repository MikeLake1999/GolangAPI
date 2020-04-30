package routes

import (
	"encoding/json"
	"gallery/models"
	"gallery/routes"
	"gallery/services"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
func TestAuthentication(t *testing.T) {
	t.Run("Should authenticate with valid account", func(t *testing.T) {
		json := strings.NewReader(`{"email":"nopex","password":"123"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/v1/authentication", json)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Should throw error with invalid email", func(t *testing.T) {
		json := strings.NewReader(`{"email":"nope1","password":"123"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/v1/authentication", json)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})

	t.Run("Should throw error with invalid password", func(t *testing.T) {
		json := strings.NewReader(`{"email":"nopex","password":"1"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/v1/authentication", json)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})
}

func TestRegistration(t *testing.T) {
	t.Run("Should register with email and password", func(t *testing.T) {
		t.Skip("Skipped after passed")
		json := strings.NewReader(`{"email":"nope2","password":"123"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/v1/registration", json)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Should throw error if email already exist", func(t *testing.T) {
		json := strings.NewReader(`{"email":"nopex","password":"123"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/v1/registration", json)
		ts.ServeHTTP(writer, request)
		if writer.Code != 400 {
			t.Errorf("Response code should be 400, was: %d", writer.Code)
		}
	})
}

func TestGetAccount(t *testing.T) {
	// get token first
	postData := strings.NewReader(`{"email":"nopex","password":"123"}`)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/v1/authentication", postData)
	ts.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code should be Ok, was: %d", writer.Code)
	}

	token := writer.Body.String()

	t.Run("Shoud get account for authorized user", func(t *testing.T) {
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/v1/account", nil)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}

		var account models.Account
		err := json.Unmarshal(writer.Body.Bytes(), &account)
		if err != nil {
			t.Errorf("Response should be an account object")
		}

		if account.Email != "nopex" {
			t.Errorf("Returned account should be nopex, was: %s", account.Email)
		}
	})

	t.Run("Shoud throw error if not authorized", func(t *testing.T) {
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/v1/account", nil)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})
}
func TestUpdateAccount(t *testing.T) {
	// get token first
	postData := strings.NewReader(`{"email":"nopex","password":"123"}`)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/v1/authentication", postData)
	ts.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code should be Ok, was: %d", writer.Code)
	}

	token := writer.Body.String()

	t.Run("Shoud update account for authorized user", func(t *testing.T) {
		t.Skip("Skipped after passed")
		postData := strings.NewReader(`{"email":"nopex", "name":"dat", "address":"TH","phone":"0000"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/account", postData)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Shoud update account for authorized user", func(t *testing.T) {
		t.Skip("Skipped after passed")
		postData := strings.NewReader(`{"password":"1234"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/account/password", postData)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Shoud throw error if not authorized", func(t *testing.T) {
		postData := strings.NewReader(`{"email":"nopex", "name":"dat", "address":"TH","phone":"0000"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/account", postData)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})
	t.Run("Shoud throw error if not authorized", func(t *testing.T) {
		postData := strings.NewReader(`{"password":"1234"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/account/password", postData)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})
}

func TestDeleteAccount(t *testing.T) {
	// get token first
	postData := strings.NewReader(`{"email":"nopex","password":"123"}`)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/v1/authentication", postData)
	ts.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code should be Ok, was: %d", writer.Code)
	}

	token := writer.Body.String()

	t.Run("Shoud delete account for authorized user", func(t *testing.T) {
		t.Skip("Skipped after passed")

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/v1/account", nil)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Shoud throw error if not authorized", func(t *testing.T) {
		t.Skip("Skipped after passed")

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/v1/account", nil)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

}
