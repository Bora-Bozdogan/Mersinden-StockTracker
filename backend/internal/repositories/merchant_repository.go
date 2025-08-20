package repositories

import (
	"mersinden-stockapp/internal/models"

	"gorm.io/gorm"
)

type MerchantRepositoryInterface interface {
	ReadID(id int) (*models.Merchant, error)
	ReadUID(uid string) (*models.Merchant, error)
	ReadAll() ([]models.Merchant, error)
	DeleteItem(id int) error
	CreateMerchant(merchantReq models.MerchantRequest) error
	UpdateMerchant(uid string, merchantInfo models.MerchantInfo) error
}

type MerchantRepository struct {
	db *gorm.DB
}

// constructor
func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	repo := new(MerchantRepository)
	repo.db = db
	return repo
}

// read commands
func (r *MerchantRepository) ReadID(id int) (*models.Merchant, error) {
	res := &models.Merchant{}
	err := r.db.First(res, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return res, nil
}

func (r *MerchantRepository) ReadUID(uid string) (*models.Merchant, error) {
	res := &models.Merchant{}
	err := r.db.First(res, "uid = ?", uid)
	if err.Error != nil {
		return nil, err.Error
	}
	return res, nil
}

func (r *MerchantRepository) ReadAll() ([]models.Merchant, error) {
	var merchants []models.Merchant
	res := r.db.Find(&merchants)
	if res.Error != nil {
		return nil, res.Error
	}
	return merchants, nil
}

// create
func (r *MerchantRepository) CreateMerchant(merchantReq models.MerchantRequest) error {
	//cast merchantReq into a merchant
	merchant := models.Merchant{
		MerchantName: merchantReq.MerchantName,
		PhoneNumber:  merchantReq.PhoneNumber,
		UID:          merchantReq.UID,
		IsAdmin:      merchantReq.IsAdmin,
	}
	res := r.db.Create(&merchant)
	return res.Error
}

// update
func (r *MerchantRepository) UpdateMerchant(uid string, merchantInfo models.MerchantInfo) error {
	res := r.db.Model(&models.Merchant{}).Where("uid = ?", uid)
	res.Updates(models.Merchant{
		MerchantName: merchantInfo.MerchantName,
		PhoneNumber:  merchantInfo.PhoneNumber,
	})
	return res.Error
}

// delete
func (r *MerchantRepository) DeleteItem(id int) error {
	res := r.db.Delete(&models.Merchant{}, id)
	return res.Error
}
