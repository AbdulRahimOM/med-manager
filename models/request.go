package models

type MedicineRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	TypeID       int     `json:"typeId"`
	Price        float64 `json:"price"`
	MinStock     int     `json:"min_stock"`
	OptimalStock int     `json:"optimal_stock"`
}

type StockUpdateRequest struct {
	ID           int            `json:"-"`
	Type         string         `json:"type"`
	StockChanges []StockChanges `json:"stock_changes"`
}

type UpdateStockUpdateRequest struct {
	ID           int            `json:"-"`
	StockChanges []StockChanges `json:"stock_changes"`
}

type StockChanges struct {
	MedicineID int `json:"medicine_id"`
	Quantity   int `json:"quantity"`
}
