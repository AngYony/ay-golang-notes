package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type GormList []string

// 外部值转换为存储到数据库中的值
func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 将数据库中的值取出转换为外部值
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID        int32     `gorm:"primary_key;type:int"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}
