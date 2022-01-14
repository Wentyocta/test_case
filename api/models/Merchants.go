package models

import (
	"errors"
	"time"
	"github.com/jinzhu/gorm"
)

type Merchant struct {
    ID      uint32    `gorm:"primary_key;auto_increment" json:"id"`
    User      User    `json:"user"`
	MerchantName    string    `gorm:"size:255;not null;unique" json:"merchant_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint32 `json:"created_by"`
    UpdatedBy uint32 `json:"updated_by"`
}

func (m *Merchant) FindMerchantByUserID(db *gorm.DB, uid uint32) (*Merchant, error) {
	var err error
	err = db.Debug().Model(Merchant{}).Where("user_id = ?", uid).Take(&m).Error
	if err != nil {
		return &Merchant{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Merchant{}, errors.New("Merchant Not Found")
	}
	return m, err
}