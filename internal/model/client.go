package model

import (
	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
	"github.com/avrahambenaram/hotel-backend/internal/repository"
	"github.com/avrahambenaram/hotel-backend/internal/service"
)

type ClientModel struct {
	clientRepository repository.ClientRepository
}

func NewClientModel(clientRepository repository.ClientRepository) *ClientModel {
	return &ClientModel{
		clientRepository,
	}
}

func (c *ClientModel) FindAll() []entity.Client {
	return c.clientRepository.FindAll()
}

func (c *ClientModel) FindByID(id uint) (entity.Client, *exception.Exception) {
	return c.clientRepository.FindByID(id)
}

func (c *ClientModel) FindByCPF(cpf string) (entity.Client, *exception.Exception) {
	return c.clientRepository.FindByCPF(cpf)
}

func (c *ClientModel) Update(client entity.Client) (entity.Client, *exception.Exception) {
	_, err := c.FindByID(client.ID)
	if err != nil {
		return client, err
	}

	errFields := service.Validate.Struct(client)
	if errFields != nil {
		return client, exception.New("Campo(s) inválidos", 403)
	}

	errUpdate := c.clientRepository.Update(client)
	if errUpdate != nil {
		return client, errUpdate
	}

	return client, nil
}

func (c *ClientModel) Save(client entity.Client) (entity.Client, *exception.Exception) {
	errFields := service.Validate.Struct(client)
	if errFields != nil {
		return client, exception.New("Campo(s) inválidos", 403)
	}

	clientExists, _ := c.clientRepository.FindByCPF(client.CPF)
	if clientExists.CPF == client.CPF {
		return client, exception.New("CPF de cliente já cadastrado", 409)
	}

	err := c.clientRepository.Save(client)
	if err != nil {
		return client, err
	}

	clientSaved, _ := c.clientRepository.FindByCPF(client.CPF)

	return clientSaved, nil
}

func (c *ClientModel) Delete(id uint) *exception.Exception {
	_, err := c.FindByID(id)
	if err != nil {
		return err
	}

	errDelete := c.clientRepository.Delete(id)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
