package app1_test

import (
	"../../go-kickstart-from-java/app1"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox"
	"log"
	"net/http"
	"testing"
	"time"
)

func StartServer(port string, t *testing.T) {
	service := app1.NewService()
	router := toolbox.NewServiceRouter(
		toolbox.ServiceRouting{
			HTTPMethod: "GET",
			URI:        "/v1/reverse/{ids}",
			Handler:    service.Reverse,
			Parameters: []string{"ids"},
		},
		toolbox.ServiceRouting{
			HTTPMethod: "POST",
			URI:        "/v1/reverse/",
			Handler:    service.Reverse,
			Parameters: []string{"ids"},
		},
	)

	http.HandleFunc("/v1/", func(writer http.ResponseWriter, reader *http.Request) {
		err := router.Route(writer, reader)
		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Printf("Started test server on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func TestServiceRouter(t *testing.T) {
	go func() {
		StartServer("8082", t)
	}()

	time.Sleep(2 * time.Second)
	var result = make([]int, 0)
	{

		err := toolbox.RouteToService("get", "http://127.0.0.1:8082/v1/reverse/1,7,3", nil, &result)
		if err != nil {
			t.Errorf("Failed to send get request  %v", err)
		}
		assert.EqualValues(t, []int{3, 7, 1}, result)

	}

}
