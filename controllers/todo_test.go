package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"   
	"github.com/gin-gonic/gin"
 )
 
 func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
 }


 type Response struct {
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	Data      string    `json:"data"`
}

func TestGetAllTodos(t *testing.T){
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	test := GetAllTodos

	if test == nil {
		t.Error("JSP")
	} else{
		t.Log(c.Get("status"))
	}
}