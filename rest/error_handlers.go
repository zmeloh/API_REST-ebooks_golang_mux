package rest

import (
	"encoding/json"
	"net/http"
)

func ServerResponse(w http.ResponseWriter, httpCode int, res any) {
	w.WriteHeader(httpCode)
	if e, ok := res.(error); ok {
		json.NewEncoder(w).Encode(map[string]string{"message": e.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"message": res})
}
