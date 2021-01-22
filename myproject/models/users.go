package Models

type Users struct {
	ID       int    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Openid   string `gorm:"column:openid;type:varchar(200)"`
	UserId   string `gorm:"column:user_id;type:varchar(20)"`
	NickName string `gorm:"column:nick_name;type:varchar(200)"`
	UserName string `gorm:"column:user_name;type:varchar(200)"`
	RoleId   string `gorm:"column:role_id;type:varchar(30)"`
	ServerId string `gorm:"column:server_id;type:varchar(20)"`
	//AdId int `gorm:"column:ad_id;type:int"`
	//ChannelId int `gorm:"column:channel_id;type:int"`
	Num       int `gorm:"column:num;type:tinyint(2)"`
	UseNum    int `gorm:"column:use_num;type:smallint(4)"`
	Money     int `gorm:"column:money;type:int"`
	CreatedAt int `gorm:"column:created_at;type:int"`
}

func NewUsers() *Users {
	return &Users{}
}
