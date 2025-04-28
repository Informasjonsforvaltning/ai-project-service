# AI Project Service

This application provides an API for retrieving AI projects that are supported by the norwegian government. It has a
simple CSV parser written in Go. It consists of several Go source code files that implement a system for reading data
from a CSV file and returning it as a two-dimensional slice of strings.

For a broader understanding of the system’s context, refer to
the [architecture documentation](https://github.com/Informasjonsforvaltning/architecture-documentation) wiki. For more
specific context on this application, see the **Portal** subsystem section.

## Getting Started

These instructions will give you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

Ensure you have the following installed:

- Go

Clone the repository.

```sh
git clone https://github.com/Informasjonsforvaltning/ai-project-service.git
cd ai-project-service
```

#### Install required dependencies

```shell
go get
```

#### Start application

```sh
go run main.go
```

### API Documentation (OpenAPI)

The API documentation is available at ```openapi.yaml```.

### Running tests

```shell
go test ./test
```

To generate a test coverage report, use the following command:

```shell
go test -v -race -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic ./test
```

## Updating CSV file

1. Open a browser and navigate to https://github.com/Informasjonsforvaltning/ai-project-service
2. Click on the file `ai_projects_norwegian_state - Oversatt_v1.csv`
3. Click on the pencil icon (Edit this file). Keep this tab open.
4. Open a new window or tab and open the original spreadsheet on Google Docs
5. Make changes to the spreadsheet
6. Click on File -> Download -> Comma Seperated Values (.csv)
7. Open this file with a text editor
8. Select all content and replace the csv content on github (the previous csv editing tab)
9. Click on commit changes and write a commit message and click on *Propose changes*
10. Assign a reviewer and click on *Create pull request*
