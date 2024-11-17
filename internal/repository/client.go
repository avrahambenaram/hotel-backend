package repository

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type ClientRepository interface {
	FindAll() []entity.Client
	FindByID(id uint) (entity.Client, *exception.Exception)
	FindByCPF(cpf string) (entity.Client, *exception.Exception)
	Update(client entity.Client) *exception.Exception
	Save(client entity.Client) *exception.Exception
	Delete(id uint) *exception.Exception
}
