package model

type ProcessResponse struct {
	ID string `json:"id"`
}

// json response for points end point
type PointsResponse struct {
	Points int `json:"points"`
}
