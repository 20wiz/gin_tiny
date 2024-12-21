# Gin Tiny

## Description
Gin Tiny is a lightweight web API built using the Gin framework in Go. It provides endpoints to create and retrieve models with timestamp functionality.

## Features
- **Create Model:** Submit a JSON payload to create a new model.
- **Retrieve Model:** Fetch the created model data.

## Installation

1. **Clone the Repository**
    ```bash
    git clone https://github.com/20wiz/gin_tiny.git
    ```

2. **Navigate to the Project Directory**
    ```bash
    cd gin_tiny
    ```

3. **Install Dependencies**
    ```bash
    go get
    ```

## Usage

Run the application using the following command:
```bash
go run main.go
```

The API will be available at http://localhost:8080.

## API Endpoints
### Create Model
- URL: /model
- Method: POST
- Body: JSON representation of the model.
- Response: Confirmation message with the created model data.
### Get Model
- URL: /model
- Method: GET
- Response: Retrieved model data.
## Example
### Create a Model
```
curl -X POST http://localhost:8080/model \
-H "Content-Type: application/json" \
-d '{"message":"Sample message"}'
```
### Retrieve a Model
``` 
curl http://localhost:8080/model
```
### License
This project is licensed under the MIT License. 