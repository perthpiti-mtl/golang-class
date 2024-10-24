# Configuration & ENV

[envconfig](https://github.com/kelseyhightower/envconfig) is a Go package that populates struct fields from environment variables.
It helps:
- Map environment variables to Go struct fields.
- Specify default values.
- Mark variables as required.
- Support complex types like slices, maps, and custom types.
--------------------------------------------
## Installation

First, install the package using go get:
```bash
go get github.com/kelseyhightower/envconfig
```
Alternatively, if you’re using Go modules, you can import the package in your code, and go mod tidy will handle the installation.

## Defining a Configuration Struct

Create a struct that represents your application’s configuration parameters:
```go
type Config struct {
    Port        int      `envconfig:"PORT" default:"8080"`
    Host        string   `envconfig:"HOST" default:"localhost"`
    Debug       bool     `envconfig:"DEBUG" default:"false"`
    AllowedHosts []string `envconfig:"ALLOWED_HOSTS" default:"localhost"`
    DatabaseURL string   `envconfig:"DATABASE_URL" required:"true"`
}
```
- Fields: Represent configuration parameters.
- Struct Tags: Control how envconfig processes each field.

## Loading Environment Variables

Use envconfig.Process to load environment variables into your struct:
```go
var cfg Config
err := envconfig.Process("", &cfg)
if err != nil {
    log.Fatal(err.Error())
}
```
- The first argument is an optional prefix (more on this later).
- &cfg is a pointer to your configuration struct.

--------------------------------------------
## Using Struct Tags

Struct tags customize how envconfig maps environment variables to struct fields.

### Custom Environment Variable Names

By default, envconfig uses the uppercase name of the struct field. To customize:
```go
type Config struct {
    MyVar string `envconfig:"MY_CUSTOM_VAR"`
}
```

### Default Values
```go
type Config struct {
    Timeout time.Duration `default:"30s"`
}
```

### Required Variables

Mark variables as required:
```go
type Config struct {
    APIKey string `required:"true"`
}
```
If APIKEY isn’t set, envconfig.Process returns an error.

--------------------------------------------
### Using Prefixes

You can namespace your environment variables using prefixes:
```go
err := envconfig.Process("MYAPP", &cfg)
```
This makes envconfig look for variables like MYAPP_PORT, MYAPP_HOST, etc.

--------------------------------------------

## Example

Here’s a full example:
```go
package main

import (
    "fmt"
    "log"
    "github.com/kelseyhightower/envconfig"
)

type Config struct {
    Port         int      `envconfig:"PORT" default:"8080"`
    Host         string   `envconfig:"HOST" default:"localhost"`
    Debug        bool     `envconfig:"DEBUG" default:"false"`
    AllowedHosts []string `envconfig:"ALLOWED_HOSTS" default:"localhost"`
    DatabaseURL  string   `envconfig:"DATABASE_URL" required:"true"`
}

func main() {
    var cfg Config
    err := envconfig.Process("", &cfg)
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Printf("Config: %+v\n", cfg)
    // Your application logic here
}
```
### Setting Environment Variables:
```bash
export DATABASE_URL="postgres://user:password@localhost/dbname"
export PORT=9090
export ALLOWED_HOSTS=localhost,example.com
```
### Running the Application:
```bash
go run main.go
```
--------------------------------------------

## Using a .env File with envconfig

To read environment variables from a .env file, we’ll use the [github.com/joho/godotenv](https://github.com/joho/godotenv) package, which loads the variables into the environment so that envconfig can access them.

### Step 1: Install godotenv
```bash
go get github.com/joho/godotenv
```
### Step 2: Create a .env File

Create a .env file in the root of your project:

.env
```dotenv
PORT=9090
HOST=localhost
DEBUG=true
ALLOWED_HOSTS=localhost,example.com
DATABASE_URL=postgres://user:password@localhost/dbname
```

### Step 3: Define Your Configuration Struct

```go
package main

import (
    "fmt"
    "log"

    "github.com/joho/godotenv"
    "github.com/kelseyhightower/envconfig"
)

type Config struct {
    Port         int      `envconfig:"PORT" default:"8080"`
    Host         string   `envconfig:"HOST" default:"localhost"`
    Debug        bool     `envconfig:"DEBUG" default:"false"`
    AllowedHosts []string `envconfig:"ALLOWED_HOSTS" default:"localhost"`
    DatabaseURL  string   `envconfig:"DATABASE_URL" required:"true"`
}
```

### Step 4: Load the .env File and Process the Environment Variables
```go
func main() {
    // Load variables from .env into the environment
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    var cfg Config
    err = envconfig.Process("", &cfg)
    if err != nil {
        log.Fatalf("Error processing env variables: %v", err)
    }

    fmt.Printf("Config: %+v\n", cfg)
}
```
### Explanation
- godotenv.Load(): Reads the .env file and loads the variables into the environment.
- envconfig.Process("", &cfg): Populates the Config struct with the environment variables.

### Running the Application
```bash
go run main.go
```
Output
```bash
Config: {Port:9090 Host:localhost Debug:true AllowedHosts:[localhost example.com] DatabaseURL:postgres://user:password@localhost/dbname}
```
--------------------------------------------

## Using Nested Structs with envconfig

You can organize your configuration parameters using nested structs. This is useful for grouping related settings.

### Step 1: Define Nested Structs
```go
type ServerConfig struct {
    Port int    `envconfig:"PORT" default:"8080"`
    Host string `envconfig:"HOST" default:"localhost"`
}

type DatabaseConfig struct {
    URL      string `envconfig:"URL" required:"true"`
    Username string `envconfig:"USERNAME" default:"root"`
    Password string `envconfig:"PASSWORD"`
}

type Config struct {
    Server   ServerConfig   `envconfig:"SERVER"`
    Database DatabaseConfig `envconfig:"DATABASE"`
    Debug    bool           `envconfig:"DEBUG" default:"false"`
}
```
### Step 2: Update the .env File

.env
```dotenv
SERVER_PORT=9090
SERVER_HOST=localhost
DATABASE_URL=postgres://user:password@localhost/dbname
DATABASE_USERNAME=admin
DATABASE_PASSWORD=secret
DEBUG=true
```

### Step 3: Load and Process the Environment Variables
```go
func main() {
    // Load variables from .env into the environment
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    var cfg Config
    err = envconfig.Process("", &cfg)
    if err != nil {
        log.Fatalf("Error processing env variables: %v", err)
    }

    fmt.Printf("Config: %+v\n", cfg)
}
```
Explanation
- Nested Structs: ServerConfig and DatabaseConfig are nested within Config.
- Struct Tags: The envconfig tag in Config (envconfig:"SERVER" and envconfig:"DATABASE") acts as a prefix for the nested fields.
- Environment Variables: envconfig looks for variables like SERVER_PORT, SERVER_HOST, DATABASE_URL, etc.

### Running the Application
```bash
go run main.go
```
### Output
```bash
Config: {Server:{Port:9090 Host:localhost} Database:{URL:postgres://user:password@localhost/dbname Username:admin Password:secret} Debug:true}
```