package app3

import (
	"log"
	"github.com/viant/toolbox"
	"fmt"
	"net/http"
)

type Server struct {
	port string
	service Service
}

func (s Server) Start() {
	router := toolbox.NewServiceRouter(
		toolbox.ServiceRouting{
			HTTPMethod: "POST",
			URI:        "/v1/query/",
			Handler:    s.service.Query,
			Parameters: []string{"request"},
		},
	)

	http.HandleFunc("/v1/", func(writer http.ResponseWriter, reader *http.Request) {
		err := router.Route(writer, reader)
		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Printf("Started test server on port %v\n", s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, nil));
}


func NewServer(port string, service Service) *Server {
	return &Server{
		service:service,
		port:port,
	}
}