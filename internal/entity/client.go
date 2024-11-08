package entity

type Client struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;unique"`
	Phone string `json:"phone" gorm:"not null"`
	CPF   string `json:"cpf" gorm:"not null;unique;size:11"`
}
