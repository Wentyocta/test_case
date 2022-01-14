package models

import (
	"errors"
	"time"
	"github.com/jinzhu/gorm"
)

type Transaction struct {
    ID      uint32    `gorm:"primary_key;auto_increment" json:"id"`
    Merchant     Merchant    `json:"merchant"`
	Outlet    Outlet    `json:"outlet"`
	BillTotal      uint32    `gorm:"primary_key;auto_increment" json:"bill_total"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint32 `json:"created_by"`
    UpdatedBy uint32 `json:"updated_by"`
}

func (t *Transaction) FindTransactionByOutletID(db *gorm.DB, merchant_id uint32, outlet_id uint32) (*Transaction, error) {
	var err error
	err = db.Debug().Model(Transaction{}).Where("merchant_id = ? && outlet_id = ?", merchant_id, outlet_id).Take(&t).Error
	if err != nil {
		return &Transaction{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Transaction{}, errors.New("Transaction Not Found")
	}
	return t, err
}