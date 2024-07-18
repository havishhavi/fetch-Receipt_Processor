package main

import (
	"receipt_processor_example/controller"

	"github.com/gin-gonic/gin"
)

// // item struct to represent different item names on the list
// type Item struct {
// 	ShortDescription string `json:"shortDescrption"`
// 	Price            string `json:"price"`
// }

// // receipt that has all the req fields
// type Receipt struct {
// 	Retailer     string `json:"retailer"`
// 	PurchaseDate string `json:"purchaseDate"`
// 	PurchaseTime string `json:"purchasetime"`

// 	Total string `json:"total"`
// 	Items []Item `json:"items"`
// }

// type ProcessResponse struct {
// 	ID string `json:"id"`
// }

// // json response for points end point
// type PointsResponse struct {
// 	Points int `json:"points"`
// }

// points calculation logic
// func CalculatePoints(receipt Receipt) int {
// 	points := 0

// 	//One point for every alphanumeric character in the retailer name.
// 	for _, c := range receipt.Retailer {
// 		if unicode.IsLetter(c) || unicode.IsDigit(c) {
// 			points++
// 		}
// 	}
// 	// 	50 points if the total is a round dollar amount with no cents.
// 	if strings.HasSuffix(receipt.Total, ".00") {
// 		points += 50
// 	}

// 	// 25 points if the total is a multiple of 0.25.
// 	total, _ := strconv.ParseFloat(receipt.Total, 64)

// 	if math.Mod(total, 0.25) == 0 {
// 		points += 25
// 	}

// 	// 5 points for every two items on the receipt.
// 	points += (len(receipt.Items) / 2) * 5
// 	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// 	//based on description length
// 	for _, itm := range receipt.Items {
// 		//removing spaces between description using trimspace function
// 		td := strings.TrimSpace(itm.ShortDescription)
// 		if len(td)%3 == 0 {
// 			price, _ := strconv.ParseFloat(itm.Price, 64)
// 			points += int(math.Ceil(price * 0.2))
// 		}
// 	}
// 	// 6 points if the day in the purchase date is odd.
// 	pD, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
// 	if pD.Day()%2 != 0 {
// 		points += 6
// 	}
// 	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
// 	pt, _ := time.Parse("15:04", receipt.PurchaseTime)
// 	if pt.Hour() >= 14 && pt.Hour() < 16 {
// 		points += 10
// 	}

// 	return points
// }

// // handlers
// func processReceipt(c *gin.Context) {
// 	var receipt Receipt
// 	if err := c.ShouldBindBodyWithJSON(&receipt); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	id := uuid.New().String()
// 	points := CalculatePoints(receipt)

// 	Mu.Lock()
// 	Receipts[id] = receipt
// 	Points[id] = points
// 	Mu.Unlock()
// 	c.JSON(http.StatusOK, ProcessResponse{ID: id})

// }

// get points function
// func getPoints(c *gin.Context) {
// 	id := c.Param("id")

// 	Mu.Lock()
// 	points, exists := Points[id]
// 	Mu.Unlock()

// 	if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, PointsResponse{Points: points})
// }

func main() {
	r := gin.Default()

	r.POST("/receipts/process", controller.ProcessReceipt)
	r.GET("/receipts/:id/points", controller.GetPoints)

	r.Run("localhost:8084")
}
