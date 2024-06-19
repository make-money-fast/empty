package model

type User struct {
	Id         int     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username   string  `json:"username" gorm:"column:username;not null"`
	Password   string  `json:"password" gorm:"column:password;not null"`
	Coin       float64 `json:"coin" gorm:"column:coin;not null"` // 余额.
	CreateTime int64   `json:"create_time" gorm:"column:create_time;not null"`
}

func (User) TableName() string {
	return "user"
}

type UserCoinHistory struct {
	Id         int     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Uid        int     `json:"uid" gorm:"column:uid;not null"`
	Reason     string  `json:"reason" gorm:"column:reason;not null"`
	Coin       float64 `json:"coin" gorm:"column:coin;not null"`
	CreateTime int64   `json:"create_time" gorm:"column:create_time;not null"`
}

func (UserCoinHistory) TableName() string {
	return "user_coin_history"
}
