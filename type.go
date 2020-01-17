package main

import (
	"net/http"
)

type conn struct {
	client          *http.Client
	destinationHost string
}
