package entity

type StockUpdateRequest struct {
	ProductID	*string		`json:"product_id"`
	Stock 		*int			`json:"stock"`
}