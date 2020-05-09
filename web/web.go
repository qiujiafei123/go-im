package web

import (
	"fmt"
	"net/http"
)

type Web struct {
}

func New() *Web {
	return &Web{}
}

func (s *Web) Run() {
	addr := fmt.Sprintf(":%d", 7070)
	http.ListenAndServe(addr, http.FileServer(http.Dir("./web/")))
}
