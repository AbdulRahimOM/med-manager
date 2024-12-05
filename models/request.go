package models

type StockUpdateRequest struct {
	StockChanges []StockChanges `json:"stock_changes" validate:"required,dive"`
}

type UpdateStockUpdateRequest struct {
	StockChanges []StockChanges `json:"stock_changes"`
}

type StockChanges struct {
	MedicineID int `json:"medicine_id" validate:"required,gte=1"`
	Quantity   int `json:"quantity" validate:"required,gte=1"`
}
