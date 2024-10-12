package gorm

import (
	"log"
	"projectlocal/config"
	"projectlocal/models"
	"time"

	"github.com/google/uuid"
)

func GetSKU(sku string) ([]models.Product, error) {

	var products []models.Product

	result := config.DBProduct().Where("sku = ?", sku).Find(&products)
	if result.Error != nil {
		log.Println("Data not found for SKU :", result.Error)
		return nil, result.Error
	}

	return products, nil
}

func CreateProduct(prd models.Product, userID uuid.UUID) error {
	id, _ := uuid.NewRandom()
	now := time.Now()
	var updatedat *time.Time = nil

	insertData := models.Product{
		ID:        id,
		SKU:       prd.SKU,
		Name:      prd.Name,
		CreatedBy: userID,
		CreatedAt: now,
		UpdatedAt: updatedat,
		Status:    "draft",
	}
	err := config.DBProduct().Create(&insertData).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateProductStatus(status string, idSKU, userID uuid.UUID) error {
	now := time.Now()

	updateData := models.Product{
		Status:    status,
		UpdatedBy: userID,
		UpdatedAt: &now,
	}

	err := config.DBProduct().Model(&updateData).Where("id =?", idSKU).Updates(updateData).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteProduct(idSKU uuid.UUID) error {
	deleteData := models.Product{}

	err := config.DBProduct().Where("id =?", idSKU).Delete(&deleteData).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
