Fetch Receipt Processor
-A webservice that fulfils the documented API
 --Structure
 -i followed Model-View-Controller(MVC) architecture structure, where the software components are located into the following directories:
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

 Instructions to run application.
 - make sure Go is installed. Run the following commands
     // clone the repo and go to receipt_processor directory
        cd Receipt_Processor
     //download the dependencies
     - go mod tidy
     //run the project (by using make command)
     - make run
     // if you get any error as make is not recognized 
        //directly run the code
    - go run main.go or go run .

 Instructions to test application.
 //test the project (by using make command)
 - make test
 // if you get any error as make is not recognized 
    //directly test the code
    - go test ./model -v


 //testing with PostMan
    Once you run the code as mentioned
    Open  postman and create a POST request
        URL: `http://localhost:8084/receipts/process` 
        Set the body to RAW JSON and Give the receipt details and send the request.
    You will receive the ID

    Now create a GET request
        Url: `http://localhost:8084/receipts/{id}/points` 
            //past the ID in Get URL and send
            Eg: `http://localhost:8084/receipts/b6a1482c-a040-4659-9858-be32a1927f3f/points` 
    You will receive the points based on the criteria


    I have provided images in screenshots of Results.

 Note: the project opens the 8084 port in localhost.

