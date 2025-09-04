daemon:
	@go build -o ./bin/espnffd ./cmd/espnffd 

espndb:
	@go build -o ./bin/espndb ./cmd/espndb