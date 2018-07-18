package greeting_test

import (
	. "hello/greeting"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloFunc(t *testing.T) {
	expected := `{"txt":"hello world"}`
	r, _ := http.NewRequest(http.MethodGet, "/hello?name=world", nil)
	w := httptest.NewRecorder()

	Hello(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))
	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}
}

func TestHelloPostFunc(t *testing.T) {
	expected := `{"txt":"hello world"}`
	msg := `{"name":"world"}`
	r, _ := http.NewRequest(http.MethodPost, "/hello-post", strings.NewReader(msg))
	w := httptest.NewRecorder()

	HelloPost(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := strings.TrimSpace(string(body))
	if actual != expected {
		t.Errorf("expect '%s' but got '%s'", expected, actual)
	}
}
