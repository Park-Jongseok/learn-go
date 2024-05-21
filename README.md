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

```sh
git diff --no-color -U5 > detailed_chagnes.diff
```
