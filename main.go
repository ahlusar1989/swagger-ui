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

// cliArgs defines the configuration that the CLI
// expects. By using a struct we can very easily
// aggregate them into an object and check what are
// the expected types.
// If we need to mock this later it's just a matter
// of reusing the struct.
type cliArgs struct {
	Port int `arg:"-p" help:"port to listen to"`
}

var (
	// args is a reference to an instantiation of
	// the configuration that the CLI expects but
	// with some values set.
	// By setting some values in advance we provide
	// default values that the user might provide
	// or not.
	args = &cliArgs{
		Port: 8080,
	}
)

// main performs the main routine of the application:
//	1.	parses the args;
//	2.	analyzes the declaration of the API
//	3.	sets the implementation of the handlers
//	4.	listens on the port we want
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
