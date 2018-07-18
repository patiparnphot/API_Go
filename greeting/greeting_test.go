package greeting_test

import (
	. "hello/greeting"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloFunc(t *testing.T) {
	expected := "hello world"
	r, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	Hello(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := string(body)
	if actual != expected {
		t.Errorf("expect %s but got %s", expected, actual)
	}
}
