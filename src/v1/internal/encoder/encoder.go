package encoder

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int          `json:"-"`
	Success bool         `json:"success,notnull"`
	Payload interface{}  `json:"payload,omitempty"`
	Error   interface{}  `json:"error,omitempty"`
	Cookie  *http.Cookie `json:"-"`
}

type Encoder struct {
	w http.ResponseWriter
  encoder *json.Encoder
}

func NewEncoder(w http.ResponseWriter) Encoder {
	return Encoder{w, json.NewEncoder(w)}
}

func (enc Encoder) Encode(r Response) {
  // The SetCookie needs to be called before the WriteHeader call
  // otherwise it doesn't work
	if r.Cookie != nil {
		http.SetCookie(enc.w, r.Cookie)
	}
	enc.w.WriteHeader(r.Code) // write the corresponding http code
	r.Success = r.Error == nil
	json.NewEncoder(enc.w).Encode(r)
}
