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
\c library_db
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

PUT
```
curl -XPUT localhost:8080/book/1 \
    -d '{
            "book":{
                "title":"Formaggio",
                "author":"Geronimo Stilton",
                "average_rating":4.5
            }
    }'
```

DELETE
```
curl -XDELETE localhost:8080/book/1
```

### Docker
```
docker build --tag library-app .
docker run --rm -it -p 8080:8080 --name my-library-app library-app
```

### docker-compose
```
docker-compose -f docker-compose.test.yml build
docker-compose -f docker-compose.test.yml up
docker-compose -f docker-compose.test.yml down
```

For testing:
```
sudo docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
sudo docker-compose -f docker-compose.test.yml down
```
or:
```
sudo make test 
```

## To Do

- [x] Fix http mux in main.go
- [ ] Response must contain a field "success"
- [ ] Add logging middleware
- [ ] Add tests for book service
- [x] Move all to containers

## What I learned

### go-kit
- if I return an error different from `nil` in functions like `Make...Endpoint()`,
  then the `DecodeResponse` function is not called. Thus, I return a `nil` error
  and the response struct is filled with the error string.