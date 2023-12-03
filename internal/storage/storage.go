package storage

import "errors"

var (
	ErrorURLNotFound = errors.New("url not found")
	ErrorURLExists   = errors.New("url exists")
)
