package entity

type Client struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null" validate:"required"`
	Email string `json:"email" gorm:"not null;unique" validate:"required,email"`
	Phone string `json:"phone" gorm:"not null" validate:"required,e164"`
	CPF   string `json:"cpf" gorm:"not null;unique;size:11" validate:"cpf"`
}
