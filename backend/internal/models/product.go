package models

import (
	"errors"
	"reflect"
)

type Product struct {
	ID 				   int `gorm:"primaryKey;autoIncrement"`
	ProductName        string `gorm:"column:product_name"`
	ProductDescription string `gorm:"column:product_description"`
	MerchantID	   	   int `gorm:"column:merchant_id"`
	Price              int `gorm:"column:price"`
	Stock              int `gorm:"column:stock"`
}

type ProductRequest struct {
	ProductName        string `gorm:"column:product_name"`
	ProductDescription string `gorm:"column:product_description"`
	Price              int `gorm:"column:price"`
	Stock              int `gorm:"column:stock"`
}

/* Validation function, must cast all items to pointers for it to work */
func (p Product) Validate() error {
	val := reflect.ValueOf(p)

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).IsNil() {
			return errors.New("invalid json")
		}
	}

	return nil
}
