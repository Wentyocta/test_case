package models

import (
	"errors"
	"time"
	"github.com/jinzhu/gorm"
)

type Outlet struct {
    ID      uint32    `gorm:"primary_key;auto_increment" json:"id"`
    Merchant     Merchant    `json:"merchant"`
	OutletName    string    `gorm:"size:255;not null;unique" json:"outlet_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint32 `json:"created_by"`
    UpdatedBy uint32 `json:"updated_by"`
}

func (o *Outlet) FindOutletByMerchantID(db *gorm.DB, uid uint32) (*Outlet, error) {
	var err error
	err = db.Debug().Model(Outlet{}).Where("merchant_id = ?", uid).Take(&o).Error
	if err != nil {
		return &Outlet{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Outlet{}, errors.New("Outlet Not Found")
	}
	return o, err
}
