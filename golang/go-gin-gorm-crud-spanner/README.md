
# Go Gin GORM CRUD API

This is a simple CRUD API built with the Go Gin GORM framework. The API manages user data with an Spanner database, allowing for basic user operations such as Create, Read, Update, and Delete (CRUD).

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
  - [Create User](#create-user)
  - [Get All Users](#get-all-users)
  - [Get User by ID](#get-user-by-id)
  - [Update User](#update-user)
  - [Delete User](#delete-user)

## Features
- Create, Read, Update, and Delete users.
- GCP Spanner database for fast operations.
- RESTful API built with the Go Gin GORM framework.

## Installation

1. **Create and query a database using the Google Cloud console:**

   - [Create and query a database using the Google Cloud console](https://cloud.google.com/spanner/docs/create-query-database-console)

2. **Google Cloud SDK:**

   - [Install the Google Cloud CLI](https://cloud.google.com/sdk/docs/install-sdk)

3. **Clone the repository:**

   ```bash
   git clone github.com/gitish/polyglot_training.git
   cd golang/go-gin-gorm-crud-spanner
   ```

4. **Install dependencies:**

   Make sure you have Go installed, then run:

   ```bash
   go mod tidy
   ```

5. **Configure DB:**

   Provide your Spanner configuration properties like project id, instance id, database id in config package.

6. **Run the application:**

   ```bash
   go run golang/go-gin-gorm-crud-spanner/main.go
   ```

   The server will start on `http://localhost:8081`.

## Usage

You can use [Postman](https://www.postman.com/) or any HTTP client to interact with the API. Below are the available endpoints:

## API Endpoints

### Create User
- **Endpoint:** `POST /users`
- **Description:** Creates a new user.
- **Request Body:**
  ```json
  {
      "id": 1,
      "name": "John Doe",
      "email": "johndoe@example.com"
  }
  ```
- **Response:**
  ```json
  {
      "id": 1,
      "name": "John Doe",
      "email": "johndoe@example.com"
  }
  ```

### Get All Users
- **Endpoint:** `GET /users`
- **Description:** Retrieves a list of all users.
- **Response:**
  ```json
  [
      {
          "id": 1,
          "name": "John Doe",
          "email": "johndoe@example.com"
      }
  ]
  ```

### Get User by ID
- **Endpoint:** `GET /users/:id`
- **Description:** Retrieves a user by their ID.
- **Example URL:** `GET /users/1`
- **Response:**
  ```json
  {
      "id": 1,
      "name": "John Doe",
      "email": "johndoe@example.com"
  }
  ```

### Update User
- **Endpoint:** `PUT /users/:id`
- **Description:** Updates an existing user by ID.
- **Example URL:** `PUT /users/1`
- **Request Body:**
  ```json
  {
      "id": 1,
      "name": "Jane Doe",
      "email": "janedoe_new@example.com"
  }
  ```
- **Response:**
  ```json
  {
      "id": 1,
      "name": "Jane Doe",
      "email": "janedoe_new@example.com"
  }
  ```

### Delete User
- **Endpoint:** `DELETE /users/:id`
- **Description:** Deletes a user by their ID.
- **Example URL:** `DELETE /users/1`
- **Response:**
  ```json
  {
      "message": "User deleted successfully"
  }
  ```

## Acknowledgments
- [Go Gin](https://github.com/gin-gonic/gin) - The web framework used to build this API.
- [GORM](https://gorm.io/) - The ORM framework used to build this API.

---

# Middleware
Middleware is like a helper that sits between your web server and your application. It’s a piece of code that can do things like check if a user is logged in, log requests, or modify data before the main part of your app handles it. You can think of it as a checkpoint where you can add extra rules or actions for incoming requests before they reach your main application.
