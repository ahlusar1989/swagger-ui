// main declares the CLI that spins up the server
// It takes some arguments, validates if they're valid
// and match the expected type and then initialize the
// server.
package main

import (
	_ "github.com/ahlusar1989/swagger-ui/restapi"
	_ "github.com/ahlusar1989/swagger-ui/restapi/operations"
	_ "github.com/go-openapi/loads"
	_ "github.com/go-openapi/runtime/middleware"
	_ "github.com/go-openapi/swag"
	"net/http"
)

func main() {

	finishedChannel := make(chan bool)

	// serve and swagger.json + assets on another go routine
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	<-finishedChannel
}
