package tpjson

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Receive from client, via pointer
func ReceiveJSON(r *http.Request, jsonPtr interface{}) error {
	if buf, err := io.ReadAll(r.Body); err != nil {
		return err
	} else if err := json.Unmarshal(buf, jsonPtr); err != nil {
		return err
	}
	return nil
}

// Send to client, via object
func SendJSON(w http.ResponseWriter, jsonObj interface{}) error {
	if buf, err := json.Marshal(jsonObj); err != nil {
		return err
	} else if _, err := fmt.Fprint(w, string(buf)); err != nil {
		return err
	}
	return nil
}
