package response

import "github.com/google/uuid"

type DataProductResponse struct {
	ID        uuid.UUID `json:"id"`
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"createdby"`
	CraetedAt string    `json:"createdat"`
	Message   string    `json:"message"`
}

type InsertResponse struct {
	SkuInsert     string `json:"sku"`
	MessageInsert string `json:"message"`
}

type UpdateResponse struct {
	SkuUpdate     string    `json:"sku"`
	StatusUpdate  string    `json:"status"`
	UpdatedBy     uuid.UUID `json:"updatedby"`
	MessageUpdate string    `json:"message"`
}

type DeleteResponse struct {
	SkuDelete     string `json:"sku"`
	MessageDelete string `json:"message"`
}
