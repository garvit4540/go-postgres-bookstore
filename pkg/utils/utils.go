package utils

import (
	"encoding/json"
	"github.com/garvit4540/go-postgres-bookstore/pkg/trace"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		trace.PushTrace(trace.ErrorReadingBody, nil)
		panic(err)
	}
	err = json.Unmarshal([]byte(body), x)
	if err != nil {
		trace.PushTrace(trace.ErrorJsonMarshalUnMarshall, map[string]interface{}{"error": err})
		panic(err)
	}
	return
}
