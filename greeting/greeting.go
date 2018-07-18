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
