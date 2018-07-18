package greeting

import (
	"encoding/json"
	"net/http"
)

type Massage struct {
	Txt string `json:"txt"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	values := r.URL.Query()
	name := values.Get("name")
	msg := Massage{Txt: "hello " + name}

	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}

type User struct {
	Name string `json:"name"`
}

func HelloPost(w http.ResponseWriter, r *http.Request) {
	var u User
	body := json.NewDecoder(r.Body)
	body.Decode(&u)
	msg := Massage{Txt: "hello " + u.Name}

	w.Header().Set("centent-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
