# electric-challenge
Summarized from problem statement:

* Write a service that accepts a file
(in some way) calculates the amount of times each hostname has appeared and return
the data in hostname => count format

* Create a deployable artifact
* Create a CI/CD pipeline



## Requirements
golang 1.14.x
docker


## Running tests

```bash
go mod download
go test -v .
```

## Running The service
```bash
go run .
```

Or if you want to build the binary

```bash
go build -o electric .
./electric
```

# API Documentation

Allow unauthenticated user to upload log file to be processed

**URL** : `/api/v1/upload`
**Method** : `POST`
**Auth Required** : `NO`
**Permissions required** : `None`
**Data constraints**
formField
```
file=path/to/logfile
```

**Header Constraints**
The server is processing the request as a multipart fileupload

```
Content-type: multipart/form-data
```

Example with curl from cli

```bash
curl -X POST -H "Content-Type: multipart/form-data" http://localhost:8080/upload -F "file=@./fixtures/electric.log"
```
