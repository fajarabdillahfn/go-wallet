package wrapper

import (
	"encoding/json"
	"log"
	"net/http"
)

type Header struct {
	Message int `json:"message"`
}
type Response struct {
	Header Header      `json:"header"`
	Data   interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	if wrap != "" {
		wrapper := make(map[string]interface{})

		wrapper[wrap] = data

		js, err := json.Marshal(wrapper)
		if err != nil {
			return err
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(status)
		_, err = w.Write(js)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		js, err := json.Marshal(data)
		if err != nil {
			return err
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(status)
		_, err = w.Write(js)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, message string, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	if message == "" {
		message = err.Error()
	}

	theError := jsonError{
		Message: message,
	}

	err = WriteJSON(w, statusCode, theError, "error")
	if err != nil {
		log.Println(err)
	}
}
