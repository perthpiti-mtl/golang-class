***JSON marshal & unmarshal***

**1. Importing the Package**

First, you need to import the encoding/json package:

```go
import "encoding/json"
```
----------------------------------------------
**2. Marshaling Data to JSON**

*Basic Example*

Suppose you have a simple struct:

```go
type User struct {
    Name string
    Age  int
}
```

*To convert an instance of User to JSON:*

```go
user := User{Name: "Alice", Age: 30}
jsonData, err := json.Marshal(user)
if err != nil {
log.Fatal(err)
}
fmt.Println(string(jsonData))
```

*Output:*

```json
{"Name": "Alice","Age": 30}
```

**Using JSON Tags**

You can control the JSON output using struct tags:

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

Now, the output keys will be lowercase:

```json
{"name": "Alice","age": 30}
```

**Omitting Empty Fields**

Use the omitempty tag to skip fields with zero values:
```go
type User struct {
    Name      string `json:"name"`
    Age       int    `json:"age,omitempty"`
    Occupation string `json:"occupation,omitempty"`
}
```
If Occupation is empty, it won’t appear in the JSON output.

**Encoding Indented JSON**

For pretty-printed JSON:
```go
jsonData, err := json.MarshalIndent(user, "", "    ")
```

----------------------------------------------
**3. Unmarshaling JSON to Data Structures**

*Basic Example*

Given JSON data:

```json
{"name":"Bob","age":25}
```

Unmarshal it into a User struct:
```go
jsonData := []byte(`{"name":"Bob","age":25}`)
var user User
err := json.Unmarshal(jsonData, &user)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", user)
```
*Output:*
```
{Name:Bob Age:25}
```

**Handling Unknown JSON Fields**

By default, if JSON contains fields not present in your struct, Unmarshal will ignore them. To catch unknown fields:
```go
decoder := json.NewDecoder(bytes.NewReader(jsonData))
decoder.DisallowUnknownFields()
err := decoder.Decode(&user)
```
**Unmarshaling into Interface{}**

If you don’t know the structure in advance:
```go
var result interface{}
json.Unmarshal(jsonData, &result)
fmt.Printf("%+v\n", result)

// Type assert data to map[string]interface{}
m, ok := data.(map[string]interface{})
if !ok {
panic("Expected a JSON object")
}

// Accessing simple fields
name, _ := m["name"].(string)
age, _ := m["age"].(int64)

fmt.Println("Name:", name)
fmt.Println("Age:", age)
```

**Working with Nested Structures**
```go
type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

type User struct {
    Name    string  `json:"name"`
    Age     int     `json:"age"`
    Address Address `json:"address"`
}
```
JSON data:
```json
{
    "name": "Charlie",
    "age": 28,
    "address": {
        "street": "123 Main St",
        "city": "Anytown"
    }
}
```

Unmarshal as before, and nested structures will be populated accordingly.

**Working with Slices and Maps**

Slices

JSON arrays map to Go slices:
```go
jsonData := []byte(`["apple", "banana", "cherry"]`)
var fruits []string
json.Unmarshal(jsonData, &fruits)
```
Maps

JSON objects can be unmarshaled into map[string]interface{}:
```go
jsonData := []byte(`{"name":"Dana","age":22}`)
var userMap map[string]interface{}
json.Unmarshal(jsonData, &userMap)
```

**Handling Time Fields**

Go’s time.Time can be marshaled/unmarshaled with the appropriate format:
```go
type Event struct {
    Name      string    `json:"name"`
    Timestamp time.Time `json:"timestamp"`
}

event := Event{
    Name:      "Meeting",
    Timestamp: time.Now(),
}
jsonData, _ := json.Marshal(event)
```
Note: By default, time.Time is encoded in RFC 3339 format.
```json
{
  "name": "Meeting",
  "timestamp": "2023-10-01T12:34:56Z"
}
```

----------------------------------------------
**Error Handling**

Always check for errors when marshaling/unmarshaling:
```go
jsonData, err := json.Marshal(user)
if err != nil {
    // handle error
}

err = json.Unmarshal(jsonData, &user)
if err != nil {
    // handle error
}
```
**Common Pitfalls**

- Exported Fields: Only exported struct fields (those starting with capital letters) are marshaled/unmarshaled.
- Zero Values: Be cautious with zero values (like 0 for int, "" for string). Use omitempty if needed.
- Data Types: Ensure JSON data types match your struct field types.