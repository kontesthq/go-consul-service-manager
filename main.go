package main

import (
	"fmt"
	"github.com/ayushs-2k4/go-consul-service-manager/consulservicemanager"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	portInt := 9876
	consulService := consulservicemanager.NewConsulService("localhost", 5150)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error fetching hostname: %v", err)
	}
	consulService.Start(hostname, portInt, "Consul-Service-Manager-Test-Service", []string{})

	router := http.NewServeMux()

	router.HandleFunc("GET /hello", HelloGETHandler)

	server := http.Server{
		Addr:    ":" + strconv.Itoa(portInt), // Use the field name Addr for the address
		Handler: router,                      // Use the field name Handler for the router
	}

	fmt.Println("Server listening at applicationPort: " + strconv.Itoa(portInt))

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func HelloGETHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World! GET from Consul-Service-Manager\n")
}
