# Number Classification API

## API Specification

- Endpoint: `GET <live-proudction-url>/api/classify-number`
- Required Query Parameter: `number`
  - Query Parameter Type: `integer`
- Example Request: `GET <live-proudction-url>/api/classify-number?number=371`
- Example Response

  ```json
  {
      "number": 371,
      "is_prime": false,
      "is_perfect": false,
      "properties": ["armstrong", "odd"],
      "digit_sum": 11,
      "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
  }
  ```

## Errors

- Example Request: `GET <live-proudction-url>/api/classify-number?number=alphabet`
- Example Error Response for invalid input

  ```json
  {
      "number": "alphabet",
      "error": true
  }
  ```

## Development

### Prerequisites

- Go version 1.23.2 or above

1. Clone or fork the repo

    ```bash
    git clone https://github.com/nanafox/number-classifier-api-go.git 
    ```

2. Change directory into the repo

    ```bash
      cd classify-number-api-go
    ```

3. Run the local server

    ```bash
    go run main.go
    ```
