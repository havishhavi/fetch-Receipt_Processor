# Fetch Receipt Processor

A web service that fulfills the documented API for processing receipts and calculating points based on defined rules.

## Structure

This project follows the Model-View-Controller (MVC) architecture, where the software components are located in the following directories:

-receipt_processor/
        main.go :Entry Point of the application
        controller/ : handles HTTP request, interaction with models
                receipt_controller : handlers for processing receipts and retrieving points. Interacts wit the clients.
        model/ : has logic and data structures
                receipt.go : structure firo for receipts and items.
                response : structure for respnse
                points_calculator: calculates points based on the receipt rules.
                points_calculator_test: one test case provide to perform unit testing.
        view/ : formats response data
                receipt_view : sends JSON response
        Readme.md : has all the instructions to run the file on local (use make file)

        
     Fetch-Receipt_Processor/
├── main.go : Entry point of the application
├── controllers/ : Handles HTTP requests, interaction with models
│ ├── receipt_controller.go : Handlers for processing receipts and retrieving points. Interacts with the clients.
├── models/ : Contains logic and data structures
│ ├── receipt.go : Structure for receipts and items.
│ ├── response.go : Structure for responses.
│ ├── points_calculator.go : Calculates points based on the receipt rules.
│ └── points_calculator_test.go : Unit test case for the points calculator.
├── views/ : Formats response data
│ └── receipt_view.go : Sends JSON response.
├── go.mod : Dependency management file.
├── Makefile : Makefile for running and testing the application.
└── README.md : Instructions to run the application locally.

## Instructions to Run the Application

1. **Ensure Go is installed.**

2. **Clone the repository and navigate to the project directory:**

   ```bash
   git clone [<repository_url>](https://github.com/havishhavi/fetch-Receipt_Processor.git)
   cd fetch-Receipt_Processor
   
3. **Download the dependencies:**
   
   ```bash
   go mod tidy
   
5. **Run the Project**
 Using the Makefile:
```bash
 make run

If you get an error that make is not recognized, run the code directly:
```bash
 go run main.go
# or
```bash
 go run . 


 ##Instructions to test application.
 test the project (by using make command)

 ```bash
  make test

 if you get any error as make is not recognized ,
    directly test the code
    ```bash go test ./model -v


## Testing with Postman

Run the code as mentioned above.

### Create a POST Request

1. Open Postman and create a POST request:
   - **URL:** `http://localhost:8084/receipts/process`
   - Set the body to RAW JSON and provide the receipt details.
   - Send the request.
   - You will receive a response with an ID.

### Create a GET Request

2. Create a GET request:
   - **URL:** `http://localhost:8084/receipts/{id}/points`
   - Replace `{id}` with the ID received from the POST request.
   - Example: `http://localhost:8084/receipts/b6a1482c-a040-4659-9858-be32a1927f3f/points`
   - Send the request.
   - You will receive the points based on the criteria.

### Note

The project runs on port 8084 in localhost.

