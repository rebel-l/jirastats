package endpoints

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ErrorJson struct {
	Error string `json:"error"`
	res http.ResponseWriter
}

func NewErrorJson(msg string, res http.ResponseWriter) *ErrorJson {
	e := new(ErrorJson)
	e.Error = msg
	e.res = res
	return e
}

func (e *ErrorJson) SendInternalServerError() {
	e.send(http.StatusInternalServerError)
}

func (e *ErrorJson) SendNotFound() {
	e.send(http.StatusNotFound)
}

func (e *ErrorJson) send(statusCode int) {
	log.Error(e.Error)
	e.res.Header().Set(ContentHeader, ContentTypeJson)
	e.res.WriteHeader(statusCode)
	err := json.NewEncoder(e.res).Encode(e)
	if err != nil {
		log.Errorf("Wasn't able to write body: %s", err)
	}
}
