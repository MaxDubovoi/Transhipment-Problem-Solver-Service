package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
 
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
 )
 
 func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
 }
 
 func TestHelloWorld(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"ok": "Hello world",
	}
 
	// Grab our router
	router := SetupRouter()
 
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "api/v1/test")
 
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
 
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
 
	// Grab the value & whether or not it exists
	value, exists := response["ok"]
 
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["ok"], value)
 }
 func SetupRouter() *gin.Engine {
	router := gin.Default()
  
	v1 := router.Group("api/v1") 
	{
	  v1.GET("/test", GetTest)
	  v1.POST("/compute", PostComputeDistribution)
	}
  
	return router
  }