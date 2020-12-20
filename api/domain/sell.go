package domain

type Sell struct {
	TwitterID int `gorm:"PRIMARY_KEY; AUTO_INCREMENT:false"`
	PhotoID   int `gorm:"PRIMARY_KEY; AUTO_INCREMENT:false"`
	Num       int `gorm:"TYPE:INT UNSIGNED; NOT NULL"`
}
