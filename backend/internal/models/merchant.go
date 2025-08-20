package models

//default merchant struct matching db
type Merchant struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	MerchantName string `gorm:"column:merchant_name"`
	PhoneNumber  string `gorm:"column:phone_number"`
	UID          string `gorm:"column:uid"`
	IsAuth       bool   `gorm:"column:isauth"`
	IsAdmin      bool   `gorm:"column:isadmin"`
}

//new merchant creation request
type MerchantRequest struct {
	MerchantName string `gorm:"column:merchant_name"`
	PhoneNumber  string `gorm:"column:phone_number"`
	UID          string `gorm:"column:uid"`
	IsAdmin      bool   `gorm:"column:isadmin"`
}

//merchant returned to frontend for safety
type MerchantInfo struct {
	MerchantName string
	PhoneNumber string
}