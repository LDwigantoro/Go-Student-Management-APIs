# Collage Student Management APIs Using Clean Architecture

The Clean Architecture approach to software design emphasizes separating the different concerns of an application, such as the business logic, user interface, and data access, into distinct layers or components. In Go, this can be achieved by using interfaces to define the contracts between different layers and using dependency injection to create loose coupling between them.

This repository is an example of a Go project that uses clean architecture, which is a simple API for managing collage student. The project is organized into the following layers:

- **Entities**. The entities package contains the domain model for the application, including the Mahasiswa struct that represents a single collage student.
- **Usecases**. The usecases package contains the business logic for the application, including the operations that can be performed on student data.
- **Interface**. The interface package contains the data access logic for the application, including the the methods for storing and retrieving todo items from a database.
- **Frameworks and Drivers**. This package contains the user interface for the application, including the Mahasiswa handler that defines the HTTP endpoints for the API.

![alt text](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

In this repository, we have four layers:

1. Entities
2. Repository
3. Usecases
4. Handler

Each of these layers is isolated from the others and only depends on the layers below it in the architecture. For example, the Mahasiswa Handler depends on the Mahasiswa Usecases but not directly on the Mahasiswa Repository. This makes it easy to test the different components of the application in isolation and also allows for flexibility in how the application is implemented.

## Running the Application
```bash
# Clone Repository
$ git clone https://github.com/LDwigantoro/Go-Student-Management-APIs.git

# go to directory
$ cd Go-Student-Management-APIs

# Create database with name "postgres" in postresql
# Make .env file and fill DB_URL with your database url

# Install dependencies
$ go mod download

# Run the application
$ go run main.go

```
<!-- ## Running With Docker Compose

```docker
# Build docker image
$ docker build --tag go-student-apis .

# Check if the containers are running
$ docker image ls

# Stop
make stop

# Clean
make clean
``` -->

## APIs Request

| Endpoint         | HTTP Method |    Description    |
| ---------------- | :---------: | :---------------: |
| `/mahasiswa`     |   `POST`    | `Create Mahasiswa` |
| `/mahasiswa`     |    `GET`    |  `All List of Mahasiswa`  |
| `/mahasiswa/:id` |     `GET`   | `Get Mahasiswa by id` |
| `/mahasiswa/:id` |   `PUT`    | `Update Mahasiswa` |
| `/mahasiswa/:id` |    `DELETE`    | `Delete mahasiswa`  |

## Test endpoints API 

`Request`
```bash
curl -i --request POST 'http://localhost:9090/mahasiswa' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "leonardo",
    "last_name": "dwigantoro"
}'
```

`Response`
```json
{
    "status": "200",
    "message": "Perintah berhasil dilakukan",
    "data": {
        "id": 6,
        "first_name": "leonardo",
        "last_name": "dwigantoro",
        "created_at": "2022-12-14T15:03:18.1242851+07:00",
        "updated_at": "2022-12-14T15:03:18.1242851+07:00",
        "deleted_at": null
}
```