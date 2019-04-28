package Controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//Add Something
func TestNewPlayer(t *testing.T) {
	var jsonStr = []byte(`{"name":"Jean","nickname":"Neja"}`)
	if _, err := http.NewRequest("POST", "localhost:8080/player", bytes.NewBuffer(jsonStr)); err != nil {
		t.Errorf("Can't send request")
	}
	recorder := httptest.NewRecorder()
	if recorder.Code != 200 {
		t.Errorf("response code should be 200, got %v/n", recorder.Code)
	}
}

func TestNewGame(t *testing.T) {
	var jsonStr = []byte(`{
		"id":1,
		"type":"NORMAL",
		"participants": [2, 3],
		"win":false,
		"date":"Tue Apr 23 18:08:15 2019"
	}`)
	if _, err := http.NewRequest("POST", "localhost:8080/game", bytes.NewBuffer(jsonStr)); err != nil {
		t.Errorf("Can't send request")
	}
	recorder := httptest.NewRecorder()
	if recorder.Code != 200 {
		t.Errorf("response code should be 200, got %v/n", recorder.Code)
	}
}

//See one thing
func TestSeePlayer(t *testing.T) {
	var jsonStr = []byte(`{"nickname": "Don Joe"}`)
	var expected =  `{"Name":"Jon Doe","Nickname":"Don Joe","ID":"5cbe6c5328cad7bcb4e76a1a"}`
	req , err := http.NewRequest("GET", "localhost:8080/player", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Errorf("Can't send request")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SeePlayer)
	handler.ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Errorf("response code should be 200, got %v/n", rr.Code)
	}
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",rr.Body.String(), expected)
	}
}

func TestSeeGame(t *testing.T) {
	var jsonStr = []byte(`{"idgame":10}`)
	var expected =  `{"IdGame":10,"Type":"NORMAL","Participants":[2,3],"Win":false,"Date":"Tue Apr 23 18:08:15 2019"}`
	req , err := http.NewRequest("GET", "localhost:8080/game", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Errorf("Can't send request")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SeeGame)
	handler.ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Errorf("response code should be 200, got %v/n", rr.Code)
	}
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",rr.Body.String(), expected)
	}
}