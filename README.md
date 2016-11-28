# go-task
API CRUD Task build with golang.

## Configuration
Modify app.conf
### conf/app.conf
``` conf
db_host=127.0.0.1
db_user=
db_password=
db_name=
```

## Installation
``` bash
go run main.go
```
will serve http://localhost:8080

## API DOCS
I'm using Swagger for API DOCS
http://localhost:8080/swagger/

## TO DO
- JWT Authentication
- Assign Task to Person
- Send Email to assignee