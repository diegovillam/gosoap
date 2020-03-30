package gosoap

import (
	"fmt"
)

// Soap Request
type Request struct {
	Method string
	Params Params
	Action string
}

func NewRequest(m string, p Params, a string) *Request {
	return &Request{
		Method: m,
		Params: p,
		Action: a,
	}
}

type RequestStruct interface {
	SoapBuildRequest() *Request
}

func NewRequestByStruct(s RequestStruct) (*Request, error) {
	if s == nil {
		return nil, fmt.Errorf("'s' cannot be 'nil'")
	}

	return s.SoapBuildRequest(), nil
}
