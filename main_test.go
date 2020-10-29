package main

import (
	"net/http"
	"net/http/httptest"    
	"encoding/json"
	"testing"
	"github.com/gin-gonic/gin"
	"fmt"
	rand "math/rand"

	"github.com/cavdy-play/go_mongo/config"
	"github.com/cavdy-play/go_mongo/routes"
)
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func rndStr(n int) string {
	rnd_str := make([]rune, n)
	for i := range rnd_str {
		rnd_str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(rnd_str)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
   w := httptest.NewRecorder()
   req, err := http.NewRequest(method, path, nil)
   if err != nil {
	   fmt.Println(err)
   } else {
	   r.ServeHTTP(w, req)
	   fmt.Println(w.Body)
   }
   return w
}

func SetupRouter() *gin.Engine {
	// Database
	config.Connect()
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	routes.Routes(router)
	return router
}

func TestGetAllTodos(t *testing.T)  {
	w := performRequest(SetupRouter(), "GET", "/todos")
	// Test if ok
	if w.Code == http.StatusOK {
		t.Log("Data in database")
	} else {
		t.Error("No data in database")
	}
}

type JsonCreate struct {
    Title     string
    Body      string
    Completed string
}
type ArgsCreate struct {
    Json    string
    isError bool
}

func TestCreateTodo(t *testing.T)  {
	args := []ArgsCreate{
		{Json:`{"Title":"TEST1","Body":"TEST1","Completed":"TEST1"}`, isError:false},
		{Json:`{"Title":1,"Body":"TEST1","Completed":"TEST1"}`, isError:true},
		{Json:`{"Title":"TEST1","Body":{"its":"atrap"}},"Completed":"TEST1"}`, isError:true},
	}

	for _, a := range args {
		var jsonObj JsonCreate
     
		err := json.Unmarshal([]byte(a.Json), &jsonObj)
		
		if err != nil {
			if a.isError{
				t.Log("PASS : error expected")
			} else {
				t.Error("FAILED : error unexpected")
			}
		} else {
			if a.isError{
				t.Error("FAILED : error expected")
			} else {
				t.Log("PASS")
			}
		}
	}
}
