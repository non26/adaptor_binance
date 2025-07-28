package domain

type NewOrder struct {
	Symbol           string `json:"symbol"`
	Side             string `json:"side"`
	PositionSide     string `json:"positionSide"`
	Type             string `json:"type"`
	NewClientOrderId string `json:"newClientOrderId"`
}
