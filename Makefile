.PHONY:
all: build examples

.PHONY:
build:
	docker build -t gommando .

.PHONY:
examples: build
	docker run gommando go run examples/stdout/main.go
	docker run gommando go run examples/stderr/main.go
	docker run gommando go run examples/stdboth/main.go
	docker run gommando go run examples/ping/main.go
