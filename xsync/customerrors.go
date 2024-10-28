package main

import "errors"

var ErrTooManyRequests = errors.New("too many requests")

func NewErrTooManyRequests() error {
	return ErrTooManyRequests
}

var ErrNotFound = errors.New("not found")

func NewErrNotFound() error {
	return ErrNotFound
}
