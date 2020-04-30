package routes

import (
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

func TestCreateGallery(t *testing.T) {
	// get token first
	postData := strings.NewReader(`{"email":"nopex","password":"123"}`)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/v1/authentication", postData)
	ts.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code should be Ok, was: %d", writer.Code)
	}

	token := writer.Body.String()
	t.Run("Should register with email and password", func(t *testing.T) {
		t.Skip("Skipped after passed")
		json := strings.NewReader(`{"name":"OOP","brief":"123"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/v1/gallery", json)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Should throw error if email already exist", func(t *testing.T) {
		json := strings.NewReader(`{"name":"OOP","brief":"123"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/v1/gallery", json)
		ts.ServeHTTP(writer, request)
		if writer.Code != 400 {
			t.Errorf("Response code should be 400, was: %d", writer.Code)
		}
	})
}

func TestGetGallery(t *testing.T) {
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
		request, _ := http.NewRequest("GET", "/v1/gallery", nil)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}

	})

	t.Run("Shoud throw error if not authorized", func(t *testing.T) {
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/v1/gallery", nil)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})
}
func TestUpdateGallery(t *testing.T) {
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
		postData := strings.NewReader(`"name":"OOP1","brief":"1234"`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/gallery", postData)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Shoud update account for authorized user", func(t *testing.T) {
		t.Skip("Skipped after passed")

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/gallery/1/publication")
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Shoud throw error if not authorized", func(t *testing.T) {
		postData := strings.NewReader(`{"name":"OOP","brief":"123"}`)
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/gallery", postData)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})
	t.Run("Shoud throw error if not authorized", func(t *testing.T) {

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/v1/gallery/1/publication", nil)
		ts.ServeHTTP(writer, request)
		if writer.Code != 401 {
			t.Errorf("Response code should be 401, was: %d", writer.Code)
		}
	})
}

func TestDeleteGallery(t *testing.T) {
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
		request, _ := http.NewRequest("DELETE", "/v1/gallery/2", nil)
		request.Header.Set("Authorization", "Bearer "+token)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

	t.Run("Shoud throw error if not authorized", func(t *testing.T) {
		t.Skip("Skipped after passed")

		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/v1/gallery/2", nil)
		ts.ServeHTTP(writer, request)
		if writer.Code != 200 {
			t.Errorf("Response code should be Ok, was: %d", writer.Code)
		}
	})

}
