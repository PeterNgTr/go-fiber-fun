### Basic REST API in Go using Fiber

Fiber is a new Go-based web framework which has exploded onto the scene and generated a lot of interest from the programming community. The repository for the framework has consistently been on the GitHub Trending page for the Go programming language and as such, let's try building a simple REST API.

So, in this tutorial, we’ll be covering how you can get started building your own REST API systems in Go using this new Fiber framework!

By the end of this tutorial, we will have covered:

* Project Setup
* Building a Simle CRUD REST API for a Book management system
* Breaking out the project into a more extensible format with additional packages.
* Let’s dive in!

### Requirements

* Go - https://go.dev/doc/install

### Installation

* `go get` to install the required package

### Swagger docs

```
swag init --parseDependency --parseInternal
```

### Building our REST API Endpoints

We’ll create the following endpoints:

* GET /api/v1/book - which will return all of the books that you have read during lock down.
* GET /api/v1/book/:id - which takes in a path parameter for the book ID and returns just a solitary book
* POST /api/v1/book - which will allow us to add new books to the list
* DELETE /api/v1/book/:id - which will allow us to delete a book from the list in case we add any books by mistake

### Start the REST API systems

```
go run main.go

Connection Opened to Database
{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>} Adell Marks Mr. Norbert Cummerata I 5}
{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>} Conrad Koch Mr. Terrill Kuhic 5}
{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>} Ms. Barbara Smith Meda Hirthe 5}
{{0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>} Verona Moore DVM Ludwig Koelpin 5}
Database Migrated
Listening at http://localhost:3000
 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.37.1                   │ 
 │               http://127.0.0.1:3000               │ 
 │       (bound on host 0.0.0.0 and port 3000)       │ 
 │                                                   │ 
 │ Handlers ............ 10  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 94185 │ 
 └───────────────────────────────────────────────────┘ 

```
### Authentication

By default, once you start the server, admin user would be created, using the generated token to call the other APIs
```
curl --request POST \
  --url http://localhost:3000/login \
  --header 'Content-Type: application/json' \
  --data '{
  "identity": "admin",
  "password": "12345678"
}'
{"message":"Success login","status":"success","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQxMDE4MzgsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.I9BR0cyt2lhYfnpWGrWaPR4K9-G0xgOK27UFOPTogPU"}% 
```

```
curl --request GET \
  --url http://localhost:3000/api/v1/book \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQxMDE5NDQsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.7l5uLq4XvMGU1ZxRxZWj6f6jXXaM86EX2hq5xGIMyik'
```

### Testing the Endpoints

Now that we have our endpoints defined and talking to the database, the next step is to test these manually to verify if they work as intended:

```
$ curl http://localhost:3000/api/v1/book
[{"ID":3,"CreatedAt":"2020-04-24T09:20:37.622829+01:00","UpdatedAt":"2020-04-24T09:20:37.622829+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5},{"ID":4,"CreatedAt":"2020-04-24T09:29:47.573672+01:00","UpdatedAt":"2020-04-24T09:29:47.573672+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5}]

$ curl http://localhost:3000/api/v1/book/1
{"ID":3,"CreatedAt":"2020-04-24T09:20:37.622829+01:00","UpdatedAt":"2020-04-24T09:20:37.622829+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5}

$ curl -X POST http://localhost:3000/api/v1/book
{"ID":5,"CreatedAt":"2020-04-24T09:49:16.405426+01:00","UpdatedAt":"2020-04-24T09:49:16.405426+01:00","DeletedAt":null,"name":"1984","author":"George Orwell","rating":5}

$ curl -X DELETE http://localhost:3000/api/v1/book/1
Book Successfully Deleted
```