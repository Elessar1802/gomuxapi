package encoder

import (
	"encoding/json"
  "net/http"
)

type Response struct {
  Code int            `json:"-"`
	Success bool        `json:"success,notnull"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type Encoder struct {
  w http.ResponseWriter
	encoder *json.Encoder
}

func NewEncoder(w http.ResponseWriter) Encoder {
	return Encoder{w, json.NewEncoder(w)}
}

func (enc Encoder) Encode(r Response) {
  enc.w.WriteHeader(r.Code) // write the corresponding http code
  r.Success = r.Error == nil
	enc.encoder.Encode(r)
}
