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

## Updating CSV file
1. Open a browser and navigate to https://github.com/Informasjonsforvaltning/ai-project-service
1. Click on the file `ai_projects_norwegian_state - Oversatt_v1.csv`
1. Click on the pencil icon (Edit this file). Keep this tab open.
1. Open a new window or tab and open the original spreadsheet on Google Docs
1. Make changes to the spreadsheet
1. Click on File -> Download -> Comma Seperated Values (.csv)
1. Open this file with a text editor
1. Select all content and replace the csv content on github (the previous csv editing tab)
1. Click on commit changes and write a commit message and click on *Propose changes*
1. Assign a reviewer and click on *Create pull request*
