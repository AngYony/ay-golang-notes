package model

type Category struct {
	// 关于表设计时，是否为null，尽量选择不为null
	BaseModel
	Name             string    `gorm:"type:varchar(20);not null"`
	ParentCategoryID int32     `gorm:"type:int"` // 指向自己的上级Id
	ParentCategory   *Category // 自己指向自己，必须使用指针
	Level            int32     `gorm:"type:int;not null;default:1"`
	IsTab            bool      `gorm:"default:false;not null"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"` // 图片以url形式存在
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"` // 点击图片跳转的url
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"` // 是否上架
	ShipFree bool `gorm:"default:false;not null"` // 是否免运费
	IsNew    bool `gorm:"default:false;not null"` // 是否新品
	IsHot    bool `gorm:"default:false;not null"` // 是否热卖商品

	Name            string   `gorm:"type:varchar(50);not null"`   // 商品名称
	GoodsSn         string   `gorm:"type:varchar(50);not null"`   // 商品编号
	ClickNum        int32    `gorm:"type:int;default:0;not null"` // 商品点击数
	SoldNum         int32    `gorm:"type:int;default:0;not null"` // 销量
	FavNum          int32    `gorm:"type:int;default:0;not null"` // 收藏数
	MarketPrice     float32  `gorm:"not null"`                    // 商品价格，float32可以自动映射，因此不需要在gorm中指定
	ShopPrice       float32  `gorm:"not null"`                    // 售价
	GoodsBrief      string   `gorm:"type:varchar(100);not null"`  // 商品简介
	Images          GormList `gorm:"type:varchar(1000);not null"` // 商品图片
	DescImages      GormList `gorm:"type:varchar(1000);not null"` // 商品详情图片
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`  // 封面图
}
