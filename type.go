package main

import (
	"net/http"
)

type conn struct {
	client           *http.Client
	destionationHost string
}
