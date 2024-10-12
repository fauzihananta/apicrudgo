package internal

import (
	"encoding/json"
	"net/http"
	"projectlocal/central"
	"projectlocal/gorm"
	"projectlocal/middleware"
	"projectlocal/models"
	"projectlocal/response"

	"github.com/google/uuid"
)

type SKURequest struct {
	Sku []string `json:"sku"`
}

func GetProductWithSKU(w http.ResponseWriter, r *http.Request) {
	var req SKURequest
	var responses []response.DataProductResponse

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	for _, sku := range req.Sku {
		dataproduct := central.SelectProductWithSKU(sku)
		if len(dataproduct) == 0 {
			responses = append(responses, response.DataProductResponse{
				Sku:     sku,
				Message: "Data not found",
			})
			continue
		} else {
			for i := 0; i < len(dataproduct); i++ {
				responses = append(responses, response.DataProductResponse{
					ID:        dataproduct[i].ID,
					Sku:       dataproduct[i].Sku,
					Name:      dataproduct[i].Name,
					CreatedBy: dataproduct[i].CreatedBy,
					CraetedAt: dataproduct[i].CraetedAt,
				})
			}
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(responses)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var (
		prd models.Product
		ctx = r.Context()
	)

	userID, _ := middleware.GetUserID(ctx)
	createdby, _ := uuid.Parse(userID)

	err := json.NewDecoder(r.Body).Decode(&prd)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	err = gorm.CreateProduct(prd, createdby)
	if err != nil {
		http.Error(w, "Failed to create product "+err.Error(), http.StatusInternalServerError)
		return
	}

	responses := response.InsertResponse{
		SkuInsert:     prd.SKU,
		MessageInsert: "Product created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses)
}

func UpdateProductStatus(w http.ResponseWriter, r *http.Request) {
	var (
		prd models.Product
		ctx = r.Context()
	)

	err := json.NewDecoder(r.Body).Decode(&prd)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	userID, _ := middleware.GetUserID(ctx)
	updatedby, _ := uuid.Parse(userID)
	IDSku := central.SelectProductIDBySKU(prd.SKU)

	err = gorm.UpdateProductStatus(prd.Status, IDSku, updatedby)
	if err != nil {
		http.Error(w, "Failed to create product "+err.Error(), http.StatusInternalServerError)
		return
	}

	responses := response.UpdateResponse{
		SkuUpdate:     prd.SKU,
		StatusUpdate:  prd.Status,
		UpdatedBy:     updatedby,
		MessageUpdate: "Product udpated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var (
		prd       models.Product
		responses []response.DeleteResponse
	)

	err := json.NewDecoder(r.Body).Decode(&prd)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	IDSku := central.SelectProductIDBySKU(prd.SKU)

	if IDSku.String() == "00000000-0000-0000-0000-000000000000" {
		responses = append(responses, response.DeleteResponse{
			SkuDelete:     prd.SKU,
			MessageDelete: "Product not Found",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(responses)
		return
	}

	err = gorm.DeleteProduct(IDSku)
	if err != nil {
		http.Error(w, "Failed to create product "+err.Error(), http.StatusInternalServerError)
		return
	}

	responses = append(responses, response.DeleteResponse{
		SkuDelete:     prd.SKU,
		MessageDelete: "Product deleted successfully",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses)
}
