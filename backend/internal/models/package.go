package models

/* Package is the struct end user sees, 
 * it is independent of ids, and contains the relevant
 * information about the product, and the merchant in
 * a single struct
 * 
 * it is not written into the database, but only created by 
 * reading from merchant and product, and is displayed to the user
 */

type Package struct {
	ID 				   int `gorm:"primaryKey;autoIncrement"`
	ProductName        string `gorm:"column:product_name"`
	ProductDescription string `gorm:"column:product_description"`
	MerchantID	   string `gorm:"column:merchant_id"`
	Price              int `gorm:"column:price"`
	Stock              int `gorm:"column:stock"`
}