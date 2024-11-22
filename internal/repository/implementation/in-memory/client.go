package inmemory

import (
	"slices"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/exception"
)

type ClientRepository struct {
	clients []entity.Client
}

func (c ClientRepository) FindAll() []entity.Client {
	return c.clients
}

func (c ClientRepository) FindByID(id uint) (entity.Client, *exception.Exception) {
	for _, c := range c.clients {
		if c.ID == id {
			return c, nil
		}
	}
	return entity.Client{}, exception.New("Cliente não encontrado", 404)
}

func (c ClientRepository) FindByCPF(cpf string) (entity.Client, *exception.Exception) {
	for _, c := range c.clients {
		if c.CPF == cpf {
			return c, nil
		}
	}
	return entity.Client{}, exception.New("Cliente não encontrado", 404)
}

func (c ClientRepository) FindByEmail(email string) (entity.Client, *exception.Exception) {
	for _, c := range c.clients {
		if c.Email == email {
			return c, nil
		}
	}
	return entity.Client{}, exception.New("Cliente não encontrado", 404)
}

func (c *ClientRepository) Update(client entity.Client) *exception.Exception {
	for i, current := range c.clients {
		if current.ID == client.ID {
			c.clients[i] = current
			return nil
		}
	}

	return exception.New("Cliente não encontrado", 404)
}

func (c *ClientRepository) Save(client entity.Client) *exception.Exception {
	c.clients = append(c.clients, client)
	return nil
}

func (c *ClientRepository) Delete(id uint) *exception.Exception {
	for i, current := range c.clients {
		if current.ID == id {
			c.clients = slices.Delete(c.clients, i, i+1)
			return nil
		}
	}
	return exception.New("Cliente não encontrado", 404)
}
