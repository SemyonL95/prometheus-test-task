app-serve:
	docker-compose up --build -d

.PHONY: app-serve

run-tests:
	go test ./internals/cache/inmemory_test.go ./internals/cache/inmemory.go
.PHONY: run-tests

app-down:
	docker-compose down

.PHONY: app-down