package repositories

import (
	"mersinden-stockapp/internal/models"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	ReadID(id int) ([]models.Product, error)
	ReadAll() ([]models.Product, error)
	CreateItem(merchantId int, productReq models.ProductRequest) error
	UpdateItem(id int, productReq models.ProductRequest) error
	DeleteItem(id int) error
}

type ProductRepository struct {
	db *gorm.DB
}

// constructor
func NewProductRepository(db *gorm.DB) *ProductRepository {
	repo := new(ProductRepository)
	repo.db = db
	return repo
}

// read commands
func (r *ProductRepository) ReadID(merchantId int) ([]models.Product, error) {
	var products []models.Product
	res := r.db.Where("merchant_id = ?", merchantId).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}

func (r *ProductRepository) ReadAll() ([]models.Product, error) {
	var products []models.Product
	res := r.db.Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}

// create
func (r *ProductRepository) CreateItem(merchantId int, productReq models.ProductRequest) error {
	//cast productRequest into a product
	product := models.Product{
		ProductName: productReq.ProductName,
		ProductDescription: productReq.ProductDescription,
		MerchantID: merchantId,
		Price: productReq.Price,
		Stock: productReq.Stock,
	}
	res := r.db.Create(&product)
	return res.Error
}

// update
func (r *ProductRepository) UpdateItem(id int, productReq models.ProductRequest) error {
	res := r.db.Model(&models.Product{}).Where("id = ?", id)
	res.Updates(models.Product{
		ProductName:        productReq.ProductName,
		ProductDescription: productReq.ProductDescription,
		Price:              productReq.Price,
		Stock:              productReq.Stock,
	})
	return res.Error
}

// delete
func (r *ProductRepository) DeleteItem(id int) error {
	res := r.db.Delete(&models.Product{}, id)
	return res.Error
}
