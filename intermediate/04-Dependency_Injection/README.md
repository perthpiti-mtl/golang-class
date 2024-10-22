***Dependency Injection***

Dependency Injection (DI) is a design pattern that allows a class or function to receive its dependencies from external sources rather than creating them internally. This promotes loose coupling, easier testing, and better organization of code.

**Manual Dependency Injection (Without a Library)**

Let’s consider a simple example: an application that fetches data from an API and stores it in a database.

*Step 1: Define Interfaces for Dependencies*

First, define interfaces for the services your application depends on.

```go
// connector/http_client.go
package connector

type HTTPClient interface {
    Get(url string) ([]byte, error)
}
```
```go
// repository/database.go
package repository

type Database interface {
    Save(data []byte) error
}
```

*Step 2: Implement the Interfaces*

Provide concrete implementations for these interfaces.
```go
// connector/http_client_impl.go
package connector

import (
    "io"
    "net/http"
)

type RealHTTPClient struct{}

func (c *RealHTTPClient) Get(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

func NewRealHTTPClient() HTTPClient {
	return &RealHTTPClient{}
}
```
```go
// repository/database_impl.go
package repository

import "fmt"

type RealDatabase struct{}

func (db *RealDatabase) Save(data []byte) error {
    // Imagine this saves data to a real database
    fmt.Println("Data saved to the database")
    return nil
}

func NewRealDatabase() Database {
	return &RealDatabase{}
}
```

*Step 3: Create the Main Application Struct*

Inject dependencies via the constructor.
```go
// app/app.go
package app

type App struct {
    httpClient HTTPClient
    database   Database
}

func NewApp(httpClient HTTPClient, database Database) *App {
    return &App{
        httpClient: httpClient,
        database:   database,
    }
}

func (a *App) Run(url string) error {
    data, err := a.httpClient.Get(url)
    if err != nil {
        return err
    }
    return a.database.Save(data)
}
```

*Step 4: Wire Up Dependencies in main*

Manually create and inject dependencies.

```go
// main.go
package main

import (
	"example.com/04-di/m/connector"
	"example.com/04-di/m/repository"
	"example.com/04-di/app"
)

func main() {
	httpClient := connector.NewRealHTTPClient()
	database := repository.NewRealDatabase()
	appRunner := app.NewApp(httpClient, database)

	if err := appRunner.Run("https://api.thecatapi.com/v1/images/search"); err != nil {
		panic(err)
	}
}

```

Advantages of Manual DI
- Explicit Control: You have full control over how dependencies are constructed and injected.
- External Dependencies: Doesn’t rely on external libraries.

Disadvantages of Manual DI
- Boilerplate Code: Manually wiring dependencies can become cumbersome in large applications.
- or-Prone: Easy to make mistakes when manually managing dependencies.

-------
**Dependency Injection Using wire**

[wire](https://github.com/google/wire) is a code generation tool provided by Google that automates connecting components using dependency injection in Go.

*Step 1: Install wire*
```bash
go install github.com/google/wire/cmd/wire@latest
```
*Step 2: Define Providers*

Create provider functions that return instances of your dependencies.
```go
// providers.go
//go:build wireinject
// +build wireinject

package main

import (
	"example.com/04-di/m/app"
	"example.com/04-di/m/connector"
	"example.com/04-di/m/repository"
	"github.com/google/wire"
)

func InitializeApp() *app.App {
	wire.Build(
		connector.NewRealHTTPClient,
		repository.NewRealDatabase,
		app.NewApp,
	)
	return nil
}
```

**Step 3: Generate the Injector
```bash
wire
```

This will generate a wire_gen.go file with the wiring logic.


**Generated wire_gen.go**
```go
// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"example.com/04-di/m/app"
	"example.com/04-di/m/connector"
	"example.com/04-di/m/repository"
)

// Injectors from provider.go:

func InitializeApp() *app.App {
	httpClient := connector.NewRealHTTPClient()
	database := repository.NewRealDatabase()
	appApp := app.NewApp(httpClient, database)
	return appApp
}

```

*Step 4: Use the Generated Injector in main*
```go
// main.go
package main

func main() {
	appRunner := InitializeApp()
	if err := appRunner.Run("https://api.thecatapi.com/v1/images/search"); err != nil {
		panic(err)
	}
}
```

**Advantages of Using wire**
- Reduced Boilerplate: Automatically wires dependencies, reducing manual code.
- Compile-Time Safety: Errors in wiring are caught at compile time.
- Scalability: Easier to manage dependencies as the application grows.

**Disadvantages of Using wire**

- Learning Curve: Requires understanding how wire works.
- Generated Code: Adds an extra build step and generated code to your project.
- Limited Flexibility: Less explicit control over dependency construction.

