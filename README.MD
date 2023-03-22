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
22 Mar 23 10:25 CET INF Starting mockery dry-run=false version=v2.9.4
22 Mar 23 10:25 CET INF Walking dry-run=false version=v2.9.4
2023/03/22 10:25:05 internal error: package "context" without types was imported from "github.com/jonmol/mockery-test"
```
