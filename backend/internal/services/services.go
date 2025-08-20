package services

import (
	"errors"
	"mersinden-stockapp/internal/models"
	"mersinden-stockapp/internal/repositories"

	"gorm.io/gorm"
)

type ServicesInterface interface {
	GetItems(uid string) ([]models.Product, error)
	CreateItem(merchantId int, productReq models.ProductRequest) error
	UpdateItem(id int, productReq models.ProductRequest) error
	DeleteItem(id int) error
	GetMerchantID(id int) (*models.Merchant, error)
	GetMerchantUID(uid string) (*models.Merchant, error)
	CreateMerchant(merchantReq models.MerchantRequest) error
	UpdateMerchant(uid string, merchantInfo models.MerchantInfo) error
}

type ServicesStruct struct {
	productRepo  repositories.ProductRepositoryInterface
	merchantRepo repositories.MerchantRepositoryInterface
}

func NewServicesStruct(productRepo repositories.ProductRepositoryInterface, merchantRepo repositories.MerchantRepositoryInterface) ServicesStruct {
	return ServicesStruct{productRepo: productRepo, merchantRepo: merchantRepo}
}

func (s ServicesStruct) GetItems(uid string) ([]models.Product, error) {
	//get isAdmin by uid by calling merchantReadUID, save to isAdmin
	merchant, err := s.merchantRepo.ReadUID(uid)

	if err != nil {
		return nil, err
	}

	if merchant != nil && merchant.IsAdmin {
		//merchant admin, return all items
		items, err := s.productRepo.ReadAll()
		return items, err
	} else {
		//merchant not an admin, get only by id
		items, err := s.productRepo.ReadID(merchant.ID)
		return items, err

	}
}

func (s ServicesStruct) CreateItem(merchantId int, productReq models.ProductRequest) error {
	err := s.productRepo.CreateItem(merchantId, productReq)
	return err
}

func (s ServicesStruct) UpdateItem(id int, productReq models.ProductRequest) error {
	err := s.productRepo.UpdateItem(id, productReq)
	return err
}

func (s ServicesStruct) DeleteItem(id int) error {
	err := s.productRepo.DeleteItem(id)
	return err
}

func (s ServicesStruct) GetMerchantID(id int) (*models.Merchant, error) {
	merchant, err := s.merchantRepo.ReadID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} 
	}
	return merchant, err
}

func (s ServicesStruct) GetMerchantUID(uid string) (*models.Merchant, error) {
	merchant, err := s.merchantRepo.ReadUID(uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} 
	}
	return merchant, err
}

func (s ServicesStruct) CreateMerchant(merchantReq models.MerchantRequest) error {
	return s.merchantRepo.CreateMerchant(merchantReq)
}

func (s ServicesStruct) UpdateMerchant(uid string, merchantInfo models.MerchantInfo) error {
	return s.merchantRepo.UpdateMerchant(uid, merchantInfo)
}
