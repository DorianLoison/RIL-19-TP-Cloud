package main

import (
	"net/http"
	"net/http/httptest"    
	"encoding/json"
	"testing"
	"github.com/gin-gonic/gin"
	"fmt"
	"io"

	"github.com/cavdy-play/go_mongo/config"
	"github.com/cavdy-play/go_mongo/routes"
)

type JsonCreate struct {
    Title     string
    Body      string
    Completed string
}
type ArgsCreate struct {
    Json    string
    isError bool
}
type JsonResponse struct {
    Status  int
    Message string
}

func performRequest(r http.Handler, method, path string, b io.Reader) *httptest.ResponseRecorder {
   w := httptest.NewRecorder()
   req, err := http.NewRequest(method, path, b)
   if err != nil {
	   fmt.Println(err)
   } else {
	   r.ServeHTTP(w, req)
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
	w := performRequest(SetupRouter(), "GET", "/todos", nil)
	// Test if ok
	if w.Code == http.StatusOK {
		t.Log("Data in database")
	} else {
		t.Error("No data in database")
	}
}

func TestCreateTodo(t *testing.T)  {
	args := []ArgsCreate{
		{Json:`{"Title":"TEST1","Body":"TEST1","Completed":"TEST1"}`, isError:false},
		{Json:`{"Title":1,"Body":"TEST1","Completed":"TEST1"}`, isError:true},
		{Json:`{"Title":"TEST1","Body":{"its":"atrap"},"Completed":"TEST1"}`, isError:true},
	}

	for _, a := range args {
		var jsonObj JsonCreate
     
		err := json.Unmarshal([]byte(a.Json), &jsonObj)
		
		if err != nil {
			if a.isError{
				t.Logf("PASS : expected error : %v", err)
			} else {
				t.Errorf("FAILED : unexpected error : %v", err)
			}
		} else {
			if a.isError{
				t.Error("FAILED : expected error")
			} else {
				t.Log("PASS")
			}
		}
	}
	//Just to find route, won't create anything
	w := performRequest(SetupRouter(), "POST", "/todo", nil)
	var jsonObj JsonResponse
	err := json.Unmarshal([]byte(w.Body.String()), &jsonObj)
	if err == nil {
		if jsonObj.Status == http.StatusOK{
			t.Logf("PASS : Route found")
		} else {
			t.Errorf("FAILED : unexpected message : %v", jsonObj.Message)
		}
	} else {
		t.Errorf("FAILED : unexpected error : %v", err)
	}
}

func TestGetSingleTodo(t *testing.T)  {
	//Just to find route, won't get anything
	w := performRequest(SetupRouter(), "GET", "/todo/:todoId", nil)
	var jsonObj JsonResponse
	err := json.Unmarshal([]byte(w.Body.String()), &jsonObj)
	if err == nil {
		// No todo to be found with nil
		if jsonObj.Status == http.StatusNotFound{
			t.Logf("PASS : Route found")
		} else {
			t.Errorf("FAILED : unexpected message : %v", jsonObj.Message)
		}
	} else {
		t.Errorf("FAILED : unexpected error : %v", err)
	}
}


func TestEditTodo(t *testing.T)  {
	//Just to find route, won't update anything
	w := performRequest(SetupRouter(), "PUT", "/todo/:todoId", nil)
	var jsonObj JsonResponse
	err := json.Unmarshal([]byte(w.Body.String()), &jsonObj)
	if err == nil {
		if jsonObj.Status == http.StatusOK{
			t.Logf("PASS : Route found")
		} else {
			t.Errorf("FAILED : unexpected message : %v", jsonObj.Message)
		}
	} else {
		t.Errorf("FAILED : unexpected error : %v", err)
	}
}

func TestDeleteTodo(t *testing.T)  {
	//Just to find route, won't delete anything
	w := performRequest(SetupRouter(), "DELETE", "/todo/:todoId", nil)
	var jsonObj JsonResponse
	err := json.Unmarshal([]byte(w.Body.String()), &jsonObj)
	if err == nil {
		if jsonObj.Status == http.StatusOK{
			t.Logf("PASS : Route found")
		} else {
			t.Errorf("FAILED : unexpected message : %v", jsonObj.Message)
		}
	} else {
		t.Errorf("FAILED : unexpected error : %v", err)
	}
}