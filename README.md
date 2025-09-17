// Assisted by watsonx Code Assistant

# Increment Service

This is a simple Go service that increments a counter upon receiving HTTP POST requests.

## Running the Service

1. Ensure you have Go installed.
2. Clone this repository.
3. Navigate to the project directory: `cd increment-service`.
4. Run the service: `go run main.go`.
5. The service will start on `http://localhost:8080`.

## Usage

Send a POST request to `http://localhost:8080/increment` with a JSON body containing a number. For example, using `curl`:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"value": 5}' http://localhost:8080/increment
