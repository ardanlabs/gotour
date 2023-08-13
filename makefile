run:
	go run cmd/tour/main.go

setup-tooling:
	go install github.com/cosmtrek/air@latest

run-with-reload:
	air

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor