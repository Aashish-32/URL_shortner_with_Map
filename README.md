# URL Shortener Service
This is a simple URL shortener service built with Go and Fiber using base62 encoding.
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
3. Make sure you have the required packages installed by running:
```
go mod tidy
```
4.Run the service:
```
go run main.go
```
The service will start running on http://localhost:8080.

## Running Tests
To run tests for this service, follow these steps:


1. Run the tests:
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
  "short_url": "http://localhost:8080/abc"
  "long_url": "https://encyclopedia2.thefreedictionary.com"
  
  
  ```

- ### GET /:url - Resolve a shortened URL
  - Redirects to the original long URL if found.




## Project Structure
- URL-shortener_with_Map/
    - main.go
    - handlers/
        - shortner.go
        - resolver.go
        - handlers_test.go
    - helpers/
        - encode.go
    - routes/
        - routes.go



