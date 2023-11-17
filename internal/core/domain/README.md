## Domain
Insert all the models into their own files. For example, `user.go` for user model. The model should be a struct with the following format:
```go
type User struct {
    ID        string
    FirstName string
    LastName  string
    Email     string
    Password  string
}
```