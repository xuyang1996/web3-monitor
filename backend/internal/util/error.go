package util

import (
	"backend/internal/types"
	"encoding/json"
)

type ErrorResponse struct {
	types.CommonResponse
}

func (e *ErrorResponse) Error() string {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(jsonBytes)
}

func NewErrorResponse(code int, message string) error {
	resp := new(ErrorResponse)
	resp.Code = code
	resp.Message = message
	return resp
}

func NewErrorResponseByCode(code int) error {
	message, ok := MessageMap[code];
	if !ok {
		return NewErrorResponse(Unknown, MessageMap[Unknown])
	}
	return NewErrorResponse(code, message)
}
