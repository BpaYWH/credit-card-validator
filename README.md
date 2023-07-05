# Credit Card Validator - Go

This credit card validator is a HTTP server written in Go.

The validator performs the following steps:

1. It first validates the card number using the Luhn algorithm.
2. If the card number passes the Luhn algorithm validation, it proceeds to check the Issuer Identification Number (IIN) of the card.
3. If the card number fails the Luhn algorithm validation, the server responds to the request with 
```json 
{
   "valid": false, 
   "card-type": "Unknown Card Type"
}
```

## Usage

### Prerequisites

To run this application, you need to have Go installed on your machine.

### Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/BpaYWH/credit-card-validator.git
   ```

2. Navigate to the project directory:

   ```shell
   cd credit-card-validator
   ```

### Running the Server

To run the server, you have two options:

1. Using `go run`:
   ```shell
   go run .
   ```

2. Building and starting the server:
   ```shell
   go build
   ./credit-card-validator
   ```
   The `go build` command will compile the application, and then you can start the server by running the binary file generated.

The server will start listening on port 3333.

### Validating a Card Number

To validate a card number, send a GET request to `http://localhost:3333/validation` with the following request body:

```json
{
  "card-number": 45320151128336
}
```

Replace `45320151128336` with the card number you want to validate.

### Response

The server will respond with a JSON object containing the validation result:

```json
{
  "valid": true,
  "card-type": "Visa"
}
```

If the card number fails Luhn validation, the response will be:

```json
{
   "valid": false,
   "card-type": "Unknown Card Type"
}
```

If the server cannot identify the card issuer, the response will be:
```json
{
    "valid": true,
    "card-type": "Unknown Card Type"
}
```
