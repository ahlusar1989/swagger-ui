# Swagger Server POC

Step 1: Install go-swagger

* `
brew tap go-swagger/go-swagger && 
brew install go-swagger
`

Step 2: Install dependencies using Glide (install using `brew install glide`)

`glide install`

Step 3:

`go build main.go`

Step 4: 

`go run main.go`

# Generating a Server

## Resources used to create Swagger YAML that includes definitions and operations:
* [Swagger Toolbox](https://swagger-toolbox.firebaseapp.com/) - YAML --> JSON

In order to generate a server:

`swagger generate server -f ./swagger/swagger.yml`

This command will spill out the actions it takes as it generates your new REST server. Do not follow the advice at the end of the output. There’s an alternate mechanism to do so. Just use glide to update dependencies: `glide up -v`.

# Generating the UI

## Downloading SwaggerUI files

SwaggerUI can be downloaded from their GitHub Repo. Once downloaded, place the content of dist folder somewhere in your Go project. Then generate the spec:

`swagger generate spec -o ./swagger/swagger.json --scan-models`

After that, move swagger.json file to swaggerui folder, and inside index.html change url to ./swagger.json (url: "./swagger.json").

The `main.go` has two go routines that spin up two goroutines - one for each server. 

The Swagger UI dist is served using the following logic and handler function:

```
fs := http.FileServer(http.Dir("./swaggerui"))
http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
```
