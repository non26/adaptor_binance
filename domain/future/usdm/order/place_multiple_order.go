package domain

type PlaceMultipleOrder struct {
	BatchOrders []NewOrder `json:"batchOrders"`
}
