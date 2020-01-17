## Install
```
$ go build
```
## Usage
```
$ go run main.go
```
## API
### SignUp
```
[post] :8000/api/v1/users

In:
{
    username: string
    password: string
}

Out:
{
    Code: int
    Message: string
    Data: struct
}
```
### SignIn
```
[post] :8000/api/v1/login

In:
{
    username: string
    password: string
}

Out:
{
    Code: int
    Message: string
    Data: map[string]string{
        token:         
    }
}
```
### CheckToken
```
[post] :8000/api/v1/token

In:
-H "token: string"

Out:
{
    Code: int
    Message: string
    Data: nil
}
```
