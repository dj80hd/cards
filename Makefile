test:
	go fmt ./...
	go test -race
	golangci-lint run

cover: test
	goverage -covermode=set -coverprofile=cov.out `go list ./...`
	gocov convert cov.out | gocov report

coverhtml: cover
	go tool cover -html=cov.out
