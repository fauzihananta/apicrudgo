package central

import (
	"projectlocal/gorm"
	"projectlocal/response"

	"github.com/google/uuid"
)

func SelectProductWithSKU(sku string) []response.DataProductResponse {
	products, _ := gorm.GetSKU(sku)

	if len(products) == 0 {
		return []response.DataProductResponse{}
	}

	var dataproduct []response.DataProductResponse
	for _, dproduct := range products {
		cratedat := dproduct.CreatedAt.Format("2006-01-02 15:04:05")

		dataproductemp := response.DataProductResponse{
			ID:        dproduct.ID,
			Sku:       dproduct.SKU,
			Name:      dproduct.Name,
			CreatedBy: dproduct.CreatedBy,
			CraetedAt: cratedat,
		}
		dataproduct = append(dataproduct, dataproductemp)
	}
	return dataproduct
}

func SelectProductIDBySKU(sku string) uuid.UUID {
	var IDSku uuid.UUID
	defaultuuid, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")

	dataproduct := SelectProductWithSKU(sku)
	if len(dataproduct) == 0 {
		IDSku = defaultuuid
	} else {
		IDSku = dataproduct[0].ID
	}
	return IDSku
}
