package common

import (
	"github.com/shopspring/decimal"
)

// Product holds information on an item for sale.
type Product struct {
	Name             string           `json:"name"`
	DisplayImage     *string          `json:"displayImage"`
	Thumbnail        *string          `json:"thumbnail"`
	Price            *decimal.Decimal `json:"price"`
	Description      *string          `json:"description"`
	ShortDescription *string          `json:"shortDescription"`
	QtyInStock       int              `json:"qtyInStock"`
}
