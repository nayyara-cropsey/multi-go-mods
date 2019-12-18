# multi-go-mods
Testing multiple go-mods

# example
```
cd /example
go build ./...
go build main.go
./main

INFO[0000] Running                                       context=share-command operation=validation
INFO[0000] Making request: [GET] https://api.github.com/user  component=http-client
ERRO[0000] Response body = {"message":"Requires authentication","documentation_url":"https://developer.github.com/v3/users/#get-the-authenticated-user"}  component=http-client
ERRO[0000] Got error response: {"message":"Requires authentication","documentation_url":"https://developer.github.com/v3/users/#get-the-authenticated-user"}: status code: 401  context=share-command operation=validation
Error: status code: 401
Usage:
  example [flags]

Flags:
  -h, --help      help for example
      --version   version for example

FATA[0000] Failed with error: status code: 401
```