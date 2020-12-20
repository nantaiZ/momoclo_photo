package domain

type Photo struct {
	PhotoID   int    `gorm:"PRIMARY_KEY"`
	Series    int    `gorm:"TYPE:INT UNSIGNED; NOT NULL"`
	Code      int    `gorm:"TYPE:INT UNSIGNED; NOT NULL"`
	URL       string `gorm:"TYPE:VARCHAR(255)"`
	HasMomota bool   `gorm:"TYPE:BOOLEAN; NOT NULL"`
	HasTamai  bool   `gorm:"TYPE:BOOLEAN; NOT NULL"`
	HasSasaki bool   `gorm:"TYPE:BOOLEAN; NOT NULL"`
	HasTakagi bool   `gorm:"TYPE:BOOLEAN; NOT NULL"`
}
