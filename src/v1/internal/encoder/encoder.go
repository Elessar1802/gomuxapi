package encoder

import (
	"encoding/json"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"io"
)

type Response struct {
	Success bool        `json:"success,notnull"`
	Payload interface{} `json:"payload,omitempty"`
	Error   *err.Error  `json:"error,omitempty"`
}

type Encoder struct {
	encoder *json.Encoder
}

func NewEncoder(w io.Writer) Encoder {
	return Encoder{json.NewEncoder(w)}
}

func (enc Encoder) Encode(payload interface{}, e *err.Error) {
	var res Response
	if e != nil {
		res = Response{
			Success: false,
			Payload: nil,
			Error:   e,
		}
		enc.encoder.Encode(res)
		return
	}
	res = Response{
		Success: true,
		Payload: payload,
	}
	enc.encoder.Encode(res)
}
