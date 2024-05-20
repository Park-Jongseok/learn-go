```sh
go mod init learn-go
mkdir hello
cd hello
go mod init hello
touch hello.go
cd ../
go mod edit -replace hello=./hello
go mod tidy
go run .
```
