package main

import (
	"flag"
	"github.com/hunkeelin/mtls/v2/klinserver"
	"net/http"
	"net/http/cookiejar"
)

var (
	port             = flag.String("p", "2020", "the port to host it on")
	destionationHost = flag.String("d", "", "the host destionation with http/https")
)

func main() {
	flag.Parse()
	cJar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	hclient := &http.Client{
		Jar: cJar,
	}
	c := conn{
		client:          hclient,
		destinationHost: *destionationHost,
	}
	con := http.NewServeMux()
	con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c.mainHandler(w, r)
	})
	j := &klinserver.ServerConfig{
		BindPort: *port,
		BindAddr: "",
		ServeMux: con,
	}
	panic(klinserver.Server(j))

}
