run:
	go run cmd/tour/main.go

setup-tooling:
	go install github.com/cosmtrek/air@latest

run-with-reload:
	air
