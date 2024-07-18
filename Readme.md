# Fetch Receipt Processor

A web service that fulfills the documented API for processing receipts and calculating points based on defined rules.

## Structure

This project follows the Model-View-Controller (MVC) architecture, where the software components are located in the following directories:

- fetch-receipt_processor/
- **main.go**: Entry Point of the application.

### controller/
Handles HTTP requests and interaction with models.

- **receipt_controller**: Handlers for processing receipts and retrieving points. Interacts with clients.

### model/
Logic and data structures.

- **receipt.go**: Structure definition for receipts and items.
- **response.go**: Structure for responses.
- **points_calculator.go**: Calculates points based on receipt rules.
- **points_calculator_test.go**: Unit test cases for points calculator.

### view/
Formats response data.

- **receipt_view.go**: Sends JSON responses.

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

    
        go run main.go
  or
    
        go run . 

## Instructions to test application.
   test the project (by using make command)
   
        make test

   if you get any error as make is not recognized ,
     directly test the code
     
        go test ./model -v

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
### Raw json Data, Id and Points
       
           {
            "retailer": "Target",
            "purchaseDate": "2024-07-01",
            "purchaseTime": "18:01",
            "items": [
                {"shortDescription": "Air jordan Highs Retro", "price": "98.17"},
                {"shortDescription": "Yeezy V@2", "price": "200"},
                {"shortDescription": "Rolex Watch 1807", "price": "200"}
        
            ],
            "total": "498.17"
        }
                
        {
            "id": "b6a1482c-a040-4659-9858-be32a1927f3f"
        }
        {
            "points": 117
        }
        
### Note
- Screenshots are in Images.md file
- The project runs on port 8084 in localhost.

