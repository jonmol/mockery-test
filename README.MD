Run `go generate generate.go` using go 1.20.1 and you'll get:
```
mockery-test$ go generate generate.go 
22 Mar 23 10:22 CET INF Starting mockery dry-run=false version=v2.9.4
22 Mar 23 10:22 CET INF Walking dry-run=false version=v2.9.4
2023/03/22 10:22:47 internal error: package "context" without types was imported from "github.com/jonmol/mockery-test"
generate.go:3: running "mockery": exit status 1
``` 

mockery version:
```
mockery-test$ mockery --all
22 Mar 23 10:31 CET INF Starting mockery dry-run=false version=v2.15.0
22 Mar 23 10:31 CET INF Walking dry-run=false version=v2.15.0
2023/03/22 10:31:46 internal error: package "context" without types was imported from "github.com/PlusPeter/hotelcalifornia/server/handlers"
mocks.go:3: running "mockery": exit status 1
```
Switching to 1.17 works
