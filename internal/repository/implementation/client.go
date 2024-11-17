package implementation

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type ClientRepository struct{}

func (c ClientRepository) FindAll() []entity.Client {
	clients := []entity.Client{}
	entity.DB.Find(&clients)
	return clients
}

func (c ClientRepository) FindByID(id uint) (entity.Client, *exception.Exception) {
	client := entity.Client{}
	entity.DB.Where("ID = ?", id).Find(&client)
	if client.CPF == "" {
		return client, exception.New("Cliente não encontrado", 404)
	}
	return client, nil
}

func (c ClientRepository) FindByCPF(cpf string) (entity.Client, *exception.Exception) {
	client := entity.Client{}
	entity.DB.Where("CPF = ?", cpf).Find(&client)
	if client.CPF == "" {
		return client, exception.New("Cliente não encontrado", 404)
	}
	return client, nil
}

func (c ClientRepository) Update(client entity.Client) *exception.Exception {
	_, err := c.FindByID(client.ID)
	if err != nil {
		return err
	}
	entity.DB.Save(&client)
	return nil
}

func (c ClientRepository) Save(client entity.Client) *exception.Exception {
	result := entity.DB.Create(&client)
	if result.RowsAffected != 1 {
		return exception.New("Erro ao criar cliente", 500)
	}
	return nil
}

func (c ClientRepository) Delete(id uint) *exception.Exception {
	result := entity.DB.Delete(&entity.Client{}, id)
	if result.RowsAffected != 1 {
		return exception.New("Erro ao deletar cliente", 500)
	}
	return nil
}
