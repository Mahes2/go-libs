test:
	go test -coverprofile=coverage.out ./...

cover: test
	go tool cover -html=coverage.out

benchmark:
	go test -v ./... -bench=. -run=xxx -benchmem