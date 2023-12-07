# TodoServer README

TodoServer is a simple Golang application that provides a RESTful API for managing todos. The application uses Docker Compose to manage its dependencies, including a MySQL database.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Running the Application](#running-the-application)
- [Testing the Application](#testing-the-application)
- [Building the Application](#building-the-application)
- [Stopping the Application](#stopping-the-application)

## Getting Started

### Prerequisites

- **Docker:** [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose:** [Install Docker Compose](https://docs.docker.com/compose/install/)

### Running the Application

1. **Clone the repository:**

    ```bash
    git clone https://github.com/kumaresan1983/todoserver.git
    cd todoserver
    ```

2. **Create a `local.env` file:**

    Create a file named `local.env` in the project root with the following content:

    ```env
    JWT_SECRET=my_ultra_secure_secret
    TOKEN_EXPIRED_IN=60m
    TOKEN_MAXAGE=60

    GOOGLE_OAUTH_CLIENT_ID= 532663944661-apps.googleusercontent.com
    GOOGLE_OAUTH_CLIENT_SECRET= GOCSPX-uj4qAnql6Q4x
    GOOGLE_OAUTH_REDIRECT_URL=http://localhost:8080/v1/api/auth/google/callback
    DB_USERNAME=test
    DB_PASSWORD=test
    DB_HOST=mysql
    DB_PORT=3306
    DB_NAME=golang
    ```

3. **Initialize and run the application with Docker Compose:**

    ```bash
    docker-compose up --build
    ```

4. **Access the application:**

    The application should be accessible at [http://localhost:8080](http://localhost:8080).

## Testing the Application

To test the application, you can use `curl` commands. Access tokens can be obtained by logging in via Google.

### Sample curl commands:

```bash
# Get todos
curl --location --request GET 'http://localhost:8080/v1/api/todo?page=1&page_size=10&sort_by=completed&sort_order=desc' \
--header 'Content-Type: application/json' \
--header 'Authorization: YOUR_ACCESS_TOKEN'

# Create a todo
curl --location --request PUT 'http://localhost:8080/v1/api/todo' \
--header 'Content-Type: application/json' \
--header 'Authorization: YOUR_ACCESS_TOKEN' \
--data '{
    "title": "AR Rahman",
    "content": "90s Song"
}'

# Delete a todo
curl --location --request DELETE 'http://localhost:8080/v1/api/todo/TODO_ID' \
--header 'Content-Type: application/json' \
--header 'Authorization: YOUR_ACCESS_TOKEN'

# Mark a todo as complete
curl --location --request PATCH 'http://localhost:8080/v1/api/todos/TODO_ID/complete' \
--header 'Content-Type: application/json' \
--header 'Authorization: YOUR_ACCESS_TOKEN' \
--data '{
    "completed": true
}'

5. **Stop the application with Docker Compose:**

    ```bash
    docker-compose down
    ```

