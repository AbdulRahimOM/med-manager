package response

type StockSummary struct {
	MedicineID           int    `json:"medicine_id"`
	Medicine             string `json:"medicine"`
	Quantity             int    `json:"quantity"`
	MinStock             int    `json:"min_stock"`
	OptimalStock         int    `json:"optimal_stock"`
	DeficiencyToMinStock int    `json:"deficiency_to_min_stock"`
}
