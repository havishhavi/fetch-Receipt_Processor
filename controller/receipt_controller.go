package controller

import (
	"net/http"
	"receipt_processor_example/model"
	"receipt_processor_example/view"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	Receipts = make(map[string]model.Receipt)
	Points   = make(map[string]int)
	Mu       sync.Mutex
)

// handlers
func ProcessReceipt(c *gin.Context) {
	var receipt model.Receipt
	if err := c.ShouldBindBodyWithJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//// Generate a new uuid and store the receipt object in our in-memory cache.
	id := uuid.New().String()
	points := model.CalculatePoints(receipt)

	Mu.Lock()
	Receipts[id] = receipt
	Points[id] = points
	Mu.Unlock()
	//c.JSON(http.StatusOK, model.ProcessResponse{ID: id})
	view.JSONResponse(c, model.ProcessResponse{ID: id})

}

// get points function based on id
func GetPoints(c *gin.Context) {
	id := c.Param("id")

	Mu.Lock()
	points, exists := Points[id]
	Mu.Unlock()

	if !exists {
		view.JSONError(c, http.StatusNotFound, "receipt not found")
		//c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	//c.JSON(http.StatusOK, model.PointsResponse{Points: points})
	view.JSONResponse(c, model.PointsResponse{Points: points})
}
