package model

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// points calculation logic
func CalculatePoints(receipt Receipt) int {
	points := 0

	//One point for every alphanumeric character in the retailer name.
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			points++
		}
	}
	// 	50 points if the total is a round dollar amount with no cents.
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	total, _ := strconv.ParseFloat(receipt.Total, 64)

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	//based on description length
	for _, itm := range receipt.Items {
		//removing spaces between description using trimspace function
		td := strings.TrimSpace(itm.ShortDescription)
		if len(td)%3 == 0 {
			price, _ := strconv.ParseFloat(itm.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}
	// 6 points if the day in the purchase date is odd.
	pD, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if pD.Day()%2 != 0 {
		points += 6
	}
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	pt, _ := time.Parse("15:04", receipt.PurchaseTime)
	if pt.Hour() >= 14 && pt.Hour() < 16 {
		points += 10
	}

	return points
}
