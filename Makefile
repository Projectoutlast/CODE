run-app:
	go run cmd/app/main.go
run-migrate-up:
	go run cmd/migrator/main.go --storage-path=./storage/code.db --migrations-path=./migrations
run-migrate-down:
	go run cmd/migrator/main.go --storage-path=./storage/code.db --migrations-path=./migrations --direction=down
fill:
	go run fillDB/main.go
test:
	go test -cover ./... | grep -E '^(ok|FAIL)'