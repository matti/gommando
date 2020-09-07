.PHONY:
build:
	docker build -t gommando .

.PHONY:
examples: build
	docker run gommando go run examples/ping/main.go
	docker run gommando go run examples/stdout/main.go
	docker run gommando go run examples/stderr/main.go
