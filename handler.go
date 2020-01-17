package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *conn) mainHandler(w http.ResponseWriter, r *http.Request) {
	turl, err := url.Parse(temenosUrl + r.RequestURI)
	if err != nil {
		panic(err)
	}
	r.RequestURI = ""
	r.URL = turl
	r.Host = c.destionationHost
	resp, err := c.client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
	return
}
