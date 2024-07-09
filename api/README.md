# Movie Rating API

## Getting Started

### Running the Go API Locally
1. Start docker by opening the docker desktop application.
2. From a terminal opened to this repo `cd api` and run `go mod vendor`
3. `go run main.go` this command takes a minute to run for the first time

-API will start at localhost:8080

4. Find the Ports tab on the bottom of the codespace next to your Terminal. Use the url found under `Forwarded Port` for 8080 going forward 
   

5. Validate api started correctly by navigating to `http://localhost:8080/api/health` in a browser or run `curl http://localhost:8080/api/health` and confirming response body of **{"health":"OK"}**

Troubleshooting:
- If you encounter any issues building the application before start, try deleting the provided vendor file at /api/vendor and running `go mod tidy` and `go mod vendor`