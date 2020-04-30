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

func TestGetPublicAccount(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/v1/public/account/2", nil)
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
func TestGetPublicGallery(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/v1/public/gallery", nil)
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

func TestGetPublicGalleryById(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/v1/public/gallery/2", nil)
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
func TestGetPublicPhoto(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/v1/public/photo/2", nil)
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