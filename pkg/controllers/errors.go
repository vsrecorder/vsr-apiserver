package controllers

import "errors"

var (
	ErrInvalidParameter      = errors.New("invalid parameter")
	ErrOfficialEventNotFound = errors.New("official event not found")
	RecordNotFound           = errors.New("record not found")
)
