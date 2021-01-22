package Models

type Prize struct {
	ID          int    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Openid      string `gorm:"column:openid;type:varchar(200)"`
	UserId      string `gorm:"column:user_id;type:varchar(20)"`
	UserName    string `gorm:"column:user_name;type:varchar(75)"`
	PrizeName   string `gorm:"column:prize_name;type:varchar(255)"`
	GiftCode    string `gorm:"column:gift_code;type:varchar(50)"`
	Ext         string `gorm:"column:ext;type:varchar(1000)"`
	CreatedAt   int    `gorm:"column:created_at;type:int"`
	CreatedDate int    `gorm:"column:created_date;type:date"`
}

func NewPrize() *Prize {
	return &Prize{}
}
