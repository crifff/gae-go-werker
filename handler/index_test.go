package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"google.golang.org/appengine/aetest"
)

func TestIndex(t *testing.T) {
	var ins aetest.Instance
	var err error
	ins, err = aetest.NewInstance(nil)
	defer ins.Close()
	if err != nil {
		t.Fatal(err)
	}

	req, err := ins.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	Index1(res, req)

	if res.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("statusCode: %d", res.Code))
	}
	body, _ := ioutil.ReadAll(res.Body)

	if string(body) != "this page is module1" {
		t.Fatal(fmt.Sprintf("body: %s", string(body)))
	}

}
