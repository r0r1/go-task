# go-task
API CRUD Task build with golang.

## Configuration
rename config.example to config

## Installation
``` bash
go get github.com/rorikurniadi/go-task

go build

./go-task
```
will serve http://localhost:8080


# Overview
- 422 Http status code for validation errors
- /api/v1 as the base endpoint
- Requests and responses are in JSON

## API DOCS

### POST /register

*Request* :

        {
            "name": "John Doe",
            "email": "john@doe.com",
            "password": "foobarbaz",
        }

*Response* :

        {
            "ID": 10,
            "CreatedAt": "2016-12-25T23:11:21.872065994+07:00",
            "UpdatedAt": "2016-12-25T23:11:21.872065994+07:00",
            "DeletedAt": null,
            "name": "John Doe",
            "email": "john@doe.com",
            "password": "$2a$10$EJAR0Ppbe2OkzMDskarAYOhYKcCSpPecA9LX/0WWm/a.HBSY.UZdW",
            "contact": "",
            "address": "",
            "Task": null
        }

### POST /login

*Request* :

        {
            "username": "john@doe.com",
            "password": "foobarbaz",
        }

*Response* :

        {
            "expire": "2016-12-26T00:12:19+07:00",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODI2ODU5MzksImlkIjoicm9yaS5rdXJuaWFkaUB5YWhvby5jb20iLCJvcmlnX2lhdCI6MTQ4MjY4MjMzOX0.OzRTRuCh9BnNmEJEEqjwGzJiQQaUafl7lx0swPvZ0rE"
        }

## List all the tasks
`Authorization: Bearer <token>`

### GET /tasks

*Response* :

        [
            {
              "ID": 3,
              "CreatedAt": "2016-12-25T16:03:17Z",
              "UpdatedAt": "2016-12-25T16:03:17Z",
              "DeletedAt": null,
              "name": "Task 3",
              "user_id": 0,
              "parent": 0,
              "priority": 2,
              "status": "complete",
              "description": ""
            },
            {
              "ID": 2,
              "CreatedAt": "2016-12-25T15:19:31Z",
              "UpdatedAt": "2016-12-25T15:19:31Z",
              "DeletedAt": null,
              "name": "Task 2",
              "user_id": 0,
              "parent": 0,
              "priority": 5,
              "status": "progress",
              "description": ""
            },
            {
              "ID": 1,
              "CreatedAt": "2016-12-25T15:17:10Z",
              "UpdatedAt": "2016-12-25T16:06:03Z",
              "DeletedAt": null,
              "name": "Task 1",
              "user_id": 0,
              "parent": 0,
              "priority": 4,
              "status": "pending",
              "description": ""
            }
        ]

### GET /tasks/:id

*Response* :

        {
            "ID": 3,
            "CreatedAt": "2016-12-25T16:03:17Z",
            "UpdatedAt": "2016-12-25T16:03:17Z",
            "DeletedAt": null,
            "name": "Task 3",
            "user_id": 0,
            "parent": 0,
            "priority": 2,
            "status": "complete",
            "description": ""
        }

### POST /tasks

*Request* :

        {
            "name": "Task 4",
            "status": "complete",
            "priority": 5
        }

*Response* :

        {
            "id": 4
        }

### PUT /tasks/:id

*Request* :

        {
            "name": "Task 4",
            "status": "progress",
            "priority": 3
        }

*Response* :

        {
            "ID": 4,
            "CreatedAt": "2016-12-25T16:03:17Z",
            "UpdatedAt": "2016-12-25T16:03:17Z",
            "DeletedAt": null,
            "name": "Task 4",
            "user_id": 0,
            "parent": 0,
            "priority": 3,
            "status": "progress",
            "description": ""
        }

### DELETE /tasks/:id

*Response* :

        {
            "message": "Delete task has been successful."
        }

## TO DO
- Assign Task to Person
- Send Email to assignee