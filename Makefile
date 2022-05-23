.PHONY: run

FILENAME := ./cmd/interview/customers.csv

run:
	go run cmd/interview/main.go

build:
	go build cmd/interview/main.go

time:
	time go run cmd/interview/main.go

FILENAME := cmd/interview/customers-1M.csv
time-1millon:
	time go run cmd/interview/main.go
