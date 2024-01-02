package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
	c *uint
}

func NewHello(l *log.Logger) *Hello {
	var c uint = 0
	return &Hello{l, &c}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	*h.c++
	var err error = nil
	defer func() {
		apiLog(h.l, h.c, &r.RequestURI, err)
	}()

	fmt.Fprintf(rw, "Woohoo, server is running Beee!")
}
