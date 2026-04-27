package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	Router *mux.Router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Router: mux.NewRouter(),
	}
}

func (s *HTTPServer) Run(port string) error {

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), s.Router)
	if err != nil {
		return err
	}

	return nil

}
