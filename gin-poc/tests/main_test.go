package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-poc/middleware"
	"gin-poc/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.HeaderMiddleware())
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api/v1")
	routes.RegisterUserRoutes(api.Group("/users"))
	return r
}

func TestGetUsers(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetUsersQueryParam(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/?name=Alice", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateUser(t *testing.T) {
	router := setupRouter()
	newUser := routes.User{Name: "Charlie"}
	jsonValue, _ := json.Marshal(newUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestCreateUserInvalid(t *testing.T) {
	router := setupRouter()
	newUser := routes.User{Name: "A"} // too short
	jsonValue, _ := json.Marshal(newUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestHeaderMiddleware(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, "GinDockerDemo", w.Header().Get("X-Custom-Header"))
}

func TestCORS(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/api/v1/users/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
}
