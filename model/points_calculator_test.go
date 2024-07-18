package model

import "testing"

func TestCalculatePoints(t *testing.T) {

	receipt := Receipt{

		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{"Gatorade", "2.25"},
			{"Gatorade", "2.25"},
			{"Gatorade", "2.25"},
			{"Gatorade", "2.25"},
		},

		Total: "9.00",
	}

	expPoints := 109
	points := CalculatePoints(receipt)

	if points != expPoints {
		t.Errorf("expected %d points, got %d points", expPoints, points)
	}

}

// Total Points: 109
// Breakdown:
//     50 points - total is a round dollar amount
//     25 points - total is a multiple of 0.25
//     14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
//                 note: '&' is not alphanumeric
//     10 points - 2:33pm is between 2:00pm and 4:00pm
//     10 points - 4 items (2 pairs @ 5 points each)
//   + ---------
//   = 109 points
