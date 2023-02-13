# ai-project-service

## Introduction

This project is a simple CSV parser written in Go. It consists of several Go source code files that implement a system for reading data from a CSV file and returning it as a two-dimensional slice of strings.

The code consists of the following folders:

main.go - The entry point for the program. This file sets up an HTTP server and routes incoming HTTP requests to the appropriate handlers.

Handler Layer - This folder defines the HTTP handlers for the application. It implements the logic for processing HTTP requests and returning the appropriate response.

Service layer - This folder defines the interface and implementation for the CSV service. The service is responsible for reading data from a CSV file and returning it as a two-dimensional slice of strings. The implementation of the service uses the encoding/csv package from the Go standard library to read the data from the file.

## Running the project

```go
go run main.go
```

This will start an HTTP server on localhost:8080. You can access the parsed data from the CSV file by making an HTTP GET request to the following endpoint:

```bash
http://localhost:8080/data
```

The response will be a JSON object that contains the parsed data from the CSV file.
