# SQL Database

The pgx library is a pure Go driver and toolkit for PostgreSQL, offering high performance and advanced features compared to other drivers.

## Installation
install the pgx package:
```bash
go get github.com/jackc/pgx/v5
```

## Import Necessary Packages

Create a new Go file (e.g., main.go) and import the required packages:
```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/jackc/pgx/v5"
)
```
## Configure the Database Connection

You can set up the connection using either a connection string or a configuration struct.

### Connection String
```go
connStr := "postgres://username:password@localhost:5432/dbname"
```

### Configuration Struct
```go
config, err := pgx.ParseConfig("")
if err != nil {
    log.Fatalf("Unable to parse config: %v", err)
}
config.Host = "localhost"
config.Port = 5432
config.User = "username"
config.Password = "password"
config.Database = "dbname"
```

## Establish a Connection

Using the Connection String
```go
conn, err := pgx.Connect(context.Background(), connStr)
if err != nil {
    log.Fatalf("Unable to connect to database: %v", err)
}
defer conn.Close(context.Background())
```

Using the Configuration Struct
```go
conn, err := pgx.ConnectConfig(context.Background(), config)
if err != nil {
    log.Fatalf("Unable to connect to database: %v", err)
}
defer conn.Close(context.Background())
```

## Test the Connection

Execute a simple query to verify the connection:
```go
var greeting string
err = conn.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
if err != nil {
    log.Fatalf("QueryRow failed: %v", err)
}
fmt.Println(greeting)
```

## Connection Pooling with pgxpool

The pgx library includes a connection pool for efficient management of multiple database connections.

Using pgxpool
```go
import (
    "context"
    "fmt"
    "log"

    "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    connStr := "postgres://username:password@localhost:5432/dbname"

    // Create a connection pool
    pool, err := pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("Unable to create connection pool: %v", err)
    }
    defer pool.Close()

    // Test the connection
    var greeting string
    err = pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
    if err != nil {
        log.Fatalf("QueryRow failed: %v", err)
    }
    fmt.Println(greeting)
}
```

### Configuring the Connection Pool

Customize the pool settings using pgxpool.Config:
```go
config, err := pgxpool.ParseConfig(connStr)
if err != nil {
    log.Fatalf("Unable to parse config: %v", err)
}

config.MaxConns = 10
config.MinConns = 2
// Additional configurations...

pool, err := pgxpool.NewWithConfig(context.Background(), config)
```

--------------------------------------------

## Full example

Suppose that we want to store favorite image links in database that have the following structure.
```go
type Favorite struct {
    ID        int       `json:"id"`
    ImageUrl  string    `json:"image_url"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Step 1: Create the favorites Table in PostgreSQL

First, you need to create a table in your PostgreSQL database that corresponds to the Favorite struct.
```sql
CREATE TABLE favorites (
    id SERIAL PRIMARY KEY,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```
You can execute this SQL statement using a tool like psql or any PostgreSQL client.

### Step 2: Set Up Your Go Project

Initialize your Go module and install the required packages:

```bash
go mod init your_module_name
go get github.com/jackc/pgx/v5
```

### Step 3: Import Necessary Packages

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/jackc/pgx/v5"
)
```

### Step 4: Establish a Database Connection

```go
func connectDB() (*pgx.Conn, error) {
    connStr := "postgres://username:password@localhost:5432/dbname"
    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
        return nil, fmt.Errorf("unable to connect to database: %v", err)
    }
    return conn, nil
}
```
Remember to replace username, password, and dbname with your actual credentials.

### Step 5: Implement Functions to Store and Retrieve Favorites

Function to Insert a Favorite
```go
func insertFavorite(conn *pgx.Conn, favorite Favorite) (int, error) {
    var id int
    err := conn.QueryRow(
        context.Background(),
        "INSERT INTO favorites (image_url, created_at) VALUES ($1, $2) RETURNING id",
        favorite.ImageUrl,
        favorite.CreatedAt,
    ).Scan(&id)
    if err != nil {
        return 0, fmt.Errorf("insert failed: %v", err)
    }
    return id, nil
}
```
Function to Retrieve All Favorites
```go
func getFavorites(conn *pgx.Conn) ([]Favorite, error) {
    rows, err := conn.Query(context.Background(), "SELECT id, image_url, created_at FROM favorites")
    if err != nil {
        return nil, fmt.Errorf("query failed: %v", err)
    }
    defer rows.Close()

    var favorites []Favorite
    for rows.Next() {
        var fav Favorite
        err := rows.Scan(&fav.ID, &fav.ImageUrl, &fav.CreatedAt)
        if err != nil {
            return nil, fmt.Errorf("scan failed: %v", err)
        }
        favorites = append(favorites, fav)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("rows error: %v", err)
    }

    return favorites, nil
}
```

### Step 6: Putting It All Together

Here’s the complete example:
```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/jackc/pgx/v5"
)

type Favorite struct {
    ID        int       `json:"id"`
    ImageUrl  string    `json:"image_url"`
    CreatedAt time.Time `json:"created_at"`
}

func main() {
    // Connect to the database
    conn, err := connectDB()
    if err != nil {
        log.Fatalf("Database connection error: %v", err)
    }
    defer conn.Close(context.Background())

    // Create a new favorite
    newFavorite := Favorite{
        ImageUrl:  "https://example.com/image.png",
        CreatedAt: time.Now(),
    }

    // Insert the favorite into the database
    id, err := insertFavorite(conn, newFavorite)
    if err != nil {
        log.Fatalf("Insert error: %v", err)
    }
    fmt.Printf("Inserted favorite with ID: %d\n", id)

    // Retrieve all favorites
    favorites, err := getFavorites(conn)
    if err != nil {
        log.Fatalf("Retrieval error: %v", err)
    }

    // Print retrieved favorites
    for _, fav := range favorites {
        fmt.Printf("ID: %d, ImageURL: %s, CreatedAt: %s\n", fav.ID, fav.ImageUrl, fav.CreatedAt)
    }
}

func connectDB() (*pgx.Conn, error) {
    connStr := "postgres://username:password@localhost:5432/dbname"
    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
        return nil, fmt.Errorf("unable to connect to database: %v", err)
    }
    return conn, nil
}

func insertFavorite(conn *pgx.Conn, favorite Favorite) (int, error) {
    var id int
    err := conn.QueryRow(
        context.Background(),
        "INSERT INTO favorites (image_url, created_at) VALUES ($1, $2) RETURNING id",
        favorite.ImageUrl,
        favorite.CreatedAt,
    ).Scan(&id)
    if err != nil {
        return 0, fmt.Errorf("insert failed: %v", err)
    }
    return id, nil
}

func getFavorites(conn *pgx.Conn) ([]Favorite, error) {
    rows, err := conn.Query(context.Background(), "SELECT id, image_url, created_at FROM favorites")
    if err != nil {
        return nil, fmt.Errorf("query failed: %v", err)
    }
    defer rows.Close()

    var favorites []Favorite
    for rows.Next() {
        var fav Favorite
        err := rows.Scan(&fav.ID, &fav.ImageUrl, &fav.CreatedAt)
        if err != nil {
            return nil, fmt.Errorf("scan failed: %v", err)
        }
        favorites = append(favorites, fav)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("rows error: %v", err)
    }

    return favorites, nil
}
```

### Using a Connection Pool

For production applications, it’s recommended to use a connection pool:
```go
import (
    "github.com/jackc/pgx/v5/pgxpool"
)

// Replace *pgx.Conn with *pgxpool.Pool in your functions
func connectDB() (*pgxpool.Pool, error) {
    connStr := "postgres://username:password@localhost:5432/dbname"
    pool, err := pgxpool.New(context.Background(), connStr)
    if err != nil {
        return nil, fmt.Errorf("unable to create connection pool: %v", err)
    }
    return pool, nil
}
```
Adjust your functions accordingly to accept *pgxpool.Pool instead of *pgx.Conn.


Here’s the complete example using pool:
```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Favorite struct {
	ID        int       `json:"id"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// Connect to the database using a connection pool
	pool, err := connectDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer pool.Close()

	// Create a new favorite
	newFavorite := Favorite{
		ImageUrl:  "https://example.com/image.png",
		CreatedAt: time.Now(),
	}

	// Insert the favorite into the database
	id, err := insertFavorite(pool, newFavorite)
	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}
	fmt.Printf("Inserted favorite with ID: %d\n", id)

	// Retrieve all favorites
	favorites, err := getFavorites(pool)
	if err != nil {
		log.Fatalf("Retrieval error: %v", err)
	}

	// Print retrieved favorites
	for _, fav := range favorites {
		fmt.Printf("ID: %d, ImageURL: %s, CreatedAt: %s\n", fav.ID, fav.ImageUrl, fav.CreatedAt)
	}
}

func connectDB() (*pgxpool.Pool, error) {
	connStr := "postgres://username:password@localhost:5432/dbname"

	// Parse the connection string into a configuration
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	// Optionally configure the pool settings
	config.MaxConns = 10        // maximum number of connections in the pool
	config.MinConns = 2         // minimum number of connections in the pool
	config.MaxConnIdleTime = 5 * time.Minute

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return pool, nil
}

func insertFavorite(pool *pgxpool.Pool, favorite Favorite) (int, error) {
	var id int
	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO favorites (image_url, created_at) VALUES ($1, $2) RETURNING id",
		favorite.ImageUrl,
		favorite.CreatedAt,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert failed: %v", err)
	}
	return id, nil
}

func getFavorites(pool *pgxpool.Pool) ([]Favorite, error) {
	rows, err := pool.Query(context.Background(), "SELECT id, image_url, created_at FROM favorites")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var favorites []Favorite
	for rows.Next() {
		var fav Favorite
		err := rows.Scan(&fav.ID, &fav.ImageUrl, &fav.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}
		favorites = append(favorites, fav)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return favorites, nil
}
```