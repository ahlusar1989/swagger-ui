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

