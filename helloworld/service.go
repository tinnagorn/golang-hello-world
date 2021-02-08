package helloworld

import (
	"encoding/json"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HelloWorld(requestID string, req *Request) ([]byte, error) {
	result, err := json.Marshal(Response{
		Code:    0,
		Message: "Hello World",
	})
	return result, err
}
