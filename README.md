# URL Shortener Service
This is a simple URL shortener service built with Go and Fiber.
## Running the Service
### Prerequisites
- Go installed (https://golang.org/doc/install)
- Fiber package (github.com/gofiber/fiber/v2) installed

### Steps:
1. Clone the repository:
```
git clone https://github.com/Aashish-32/URL_shortner_with_Map.git
```
2. Navigate to the project directory
```
cd URL_shortner_with_Map
```
3.Run the service:
```
go run main.go
```
The service will start running on http://localhost:8080.

## Running Tests
To run tests for this service, follow these steps:

1. Make sure you have the required packages installed by running:
```
go mod tidy
```
2. Run the tests:
```
 go test ./handlers/
```

## API Endpoints
- ### POST /shorten - Shorten a URL
  - Request Body
  ```
  {
    "long_url": "https://example.com"
  }
  ```
  - Response:
  ```
  "short_url": "http://localhost:8080/abc",
    "long_url": "https://example.com"
  
  ```
  




## Project Structure


