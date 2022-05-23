FILENAME := files/customers.csv

run:
	go run cmd/interview/main.go

build:
	go build cmd/interview/main.go

time:
	time go run cmd/interview/main.go

time-1millon:
	time go run cmd/interview/main.go
