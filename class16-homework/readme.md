## class16

- Marshal/Unmarshal: For working with data already in memory ([]byte)
- Encode/Decode: For working with streams (io.Reader/io.Writer like HTTP responses, files)

### Example 1: Unmarshal (already have data)
```go
// When you already have the JSON as a string/byte slice
jsonStr := `{"name":"John","age":30}`
var person Person
err := json.Unmarshal([]byte(jsonStr), &person)
```

### Example 2: Decoder (reading from stream)
```go
// When reading from HTTP response, file, or network
resp, err := http.Get("https://api.example.com/user/1")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var user User
err = json.NewDecoder(resp.Body).Decode(&user)
```