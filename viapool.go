package viapool

import (
	"errors"
	"net/url"
)

var (
	BaseURL = &url.URL{Host: "www.viabtc.com", Scheme: "https", Path: "/"}

	// ErrInvalidArgs error invalid arguments
	ErrInvalidArgs = errors.New("error invalid arguments")

	// ErrRequestTimeout error request reach timeout
	ErrRequestTimeout = errors.New("error request reach timeout")
)
