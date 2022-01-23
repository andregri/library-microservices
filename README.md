# Library Microservices

## Microservices
1. Books service
2. Customers service
3. Borrowing service 

## Dataset
https://www.kaggle.com/jealousleopard/goodreadsbooks

## Tech Stack
- Go
    - go-kit for the microservices architecture
    - gorm to interact with the db
- PostgreSQL
- Docker

## Test

### Create DB
Create a PostgreSQL server with Docker:
```
sudo docker run --rm --name library_db -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
sudo docker exec -it library_db sh
```

Inside PostgreSQL container:
```
su postgres
psql
create database library_db;
```

### Run microservices
```
go run main.go
```

Make same requests:

POST
```
curl -XPOST localhost:8080/book \
    -d '{
            "book":{
                "title":"Atomic Habits",
                "author":"Geronimo Stilton",
                "average_rating":4.5
            }
    }'

{"id":1,"error":null}
```

GET
```
curl -XGET localhost:8080/book/1

{"book":{"id":1,"title":"Atomic Habits","author":"Geronimo Stilton","average_rating":4.5,"isbn":"","isbn13":"","language_code":"","num_pages":0,"ratings_count":0,"text_reviews_count":0,"publication_date":"0001-01-01T00:49:56+00:49","publisher":""},"error":null}
```

## To Do

- [ ] Fix http mux in main.go
- [ ] Add tests for book service
- [ ] Move all to containers

## What I learned

### go-kit
- if I return an error different from `nil` in functions like `Make...Endpoint()`,
  then the `DecodeResponse` function is not called. Thus, I return a `nil` error
  and the response struct is filled with the error string.