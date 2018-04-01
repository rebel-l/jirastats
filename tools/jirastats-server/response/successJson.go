package response

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type SuccessJson struct {
	payload interface{}
	res http.ResponseWriter
}

func NewSuccessJson(payload interface{}, res http.ResponseWriter) *SuccessJson {
	s := new(SuccessJson)
	s.payload = payload
	s.res = res
	return s
}

func (s *SuccessJson) SendOK() {
	s.send(http.StatusOK)
}

func (s *SuccessJson) send(statusCode int) {
	s.res.Header().Set(ContentHeader, ContentTypeJson)
	s.res.WriteHeader(statusCode)
	err := json.NewEncoder(s.res).Encode(s.payload)
	if err != nil {
		log.Errorf("Wasn't able to write body: %s", err)
	}
}
