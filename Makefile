setup:
	@dep ensure

run-connector:
	@go run main.go

run-room:
	@go run main.go -frontend=false -type=room

deps:
	@docker-compose up -d

kill-deps:
	@docker-compose down
