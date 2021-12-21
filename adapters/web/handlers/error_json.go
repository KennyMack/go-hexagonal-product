package handlers

import "encoding/json"

func jsonError(msg string) []byte {
	errorResult := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	r, err := json.Marshal(errorResult)
	if err != nil {
		return []byte(err.Error())
	}

	return r
}
