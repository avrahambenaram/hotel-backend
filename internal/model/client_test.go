package model_test

import (
	"testing"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/model"
	inmemory "github.com/avrahambenaram/hotel-backend/internal/repository/implementation/in-memory"
	"github.com/stretchr/testify/assert"
)

type ClientSuite struct {
	client1      entity.Client
	client2      entity.Client
	clientToSave entity.Client

	clientModel *model.ClientModel
}

func SetupClientSuite() *ClientSuite {
	clientRepository := new(inmemory.ClientRepository)
	clientModel := model.NewClientModel(clientRepository)

	client1 := entity.Client{
		ID:    0,
		Name:  "Benedita Liz Martins",
		Email: "benedita-martins72@akto.com.br",
		Phone: "+554738834841",
		CPF:   "90484210718",
	}
	client2 := entity.Client{
		ID:    1,
		Name:  "Lorenzo Fernando Jorge Ferreira",
		Email: "lorenzo_fernando_ferreira@inpa.gov.br",
		Phone: "+559238308888",
		CPF:   "61534210326",
	}
	clientToSave := entity.Client{
		ID:    2,
		Name:  "Calebe Emanuel Sérgio Ferreira",
		Email: "calebe.emanuel.ferreira@franciscofilho.adv.br",
		Phone: "+557935608525",
		CPF:   "12319852670",
	}
	clientRepository.Save(client1)
	clientRepository.Save(client2)

	suite := &ClientSuite{
		client1,
		client2,
		clientToSave,
		clientModel,
	}

	return suite
}

func TestGetAllClients(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()

	clients := suite.clientModel.FindAll()

	assert.Len(clients, 2)
}

func TestGetByClientIDNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()

	_, err := suite.clientModel.FindByID(2)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestGetByClientIDSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()

	client, _ := suite.clientModel.FindByID(suite.client1.ID)

	assert.Equal(suite.client1.CPF, client.CPF)
}

func TestGetByCPFNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()

	_, err := suite.clientModel.FindByCPF("12345678912")

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestGetByCPFSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()

	client, _ := suite.clientModel.FindByCPF(suite.client2.CPF)

	assert.Equal(suite.client2.ID, client.ID)
}

func TestUpdateClientNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := suite.client1
	update.ID = 2

	_, err := suite.clientModel.Update(update)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestUpdateClientCPFNotValidLength(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := suite.client1
	update.CPF = "123456"

	_, err := suite.clientModel.Update(update)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestUpdateClientCPFNotValid(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := suite.client1
	update.CPF = "12345678912"

	_, err := suite.clientModel.Update(update)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestUpdateClientCPFAlreadyUsed(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := suite.client1
	update.CPF = suite.client2.CPF

	_, err := suite.clientModel.Update(update)

	if assert.NotNil(err) {
		assert.Equal(409, err.Status)
	}
}

func TestUpdateClientEmailNotValid(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := suite.client1
	update.Email = "emailnotvalid"

	_, err := suite.clientModel.Update(update)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestUpdateClientEmailAlreadyUsed(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := suite.client1
	update.Email = suite.client2.Email

	_, err := suite.clientModel.Update(update)

	if assert.NotNil(err) {
		assert.Equal(409, err.Status)
	}
}

func TestUpdateClientPhoneNotValid(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := suite.client1
	update.Email = "01234"

	_, err := suite.clientModel.Update(update)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestUpdateClientSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	update := entity.Client{
		ID:    suite.client1.ID,
		Name:  "Camila Antônia Lopes",
		Email: "camila_antonia_lopes@aliancacadeiras.com.br",
		Phone: "+556828570924",
		CPF:   "53243790506",
	}

	clientUpdated, _ := suite.clientModel.Update(update)

	assert.Equal(update.Name, clientUpdated.Name)
	assert.Equal(update.Email, clientUpdated.Email)
	assert.Equal(update.Phone, clientUpdated.Phone)
	assert.Equal(update.CPF, clientUpdated.CPF)
}

func TestSaveClientCPFAlreadyUsed(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	save := suite.clientToSave
	save.CPF = suite.client1.CPF

	_, err := suite.clientModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(409, err.Status)
	}
}

func TestSaveClientEmailAlreadyUsed(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	save := suite.clientToSave
	save.Email = suite.client1.Email

	_, err := suite.clientModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(409, err.Status)
	}
}

func TestSaveClientCPFNotValidLength(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	save := suite.clientToSave
	save.CPF = "123456"

	_, err := suite.clientModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestSaveClientCPFNotValid(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	save := suite.clientToSave
	save.CPF = "12345678912"

	_, err := suite.clientModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestSaveClientEmailNotValid(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	save := suite.clientToSave
	save.Email = "emailnotvalid"

	_, err := suite.clientModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestSaveClientPhoneNotValid(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	save := suite.clientToSave
	save.Email = "01234"

	_, err := suite.clientModel.Save(save)

	if assert.NotNil(err) {
		assert.Equal(403, err.Status)
	}
}

func TestSaveClientSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()
	save := suite.clientToSave

	clientSaved, _ := suite.clientModel.Save(save)

	assert.Equal(save.ID, clientSaved.ID)
	assert.Equal(save.Name, clientSaved.Name)
	assert.Equal(save.Email, clientSaved.Email)
	assert.Equal(save.Phone, clientSaved.Phone)
	assert.Equal(save.CPF, clientSaved.CPF)
}

func TestDeleteClientNotFound(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()

	err := suite.clientModel.Delete(2)

	if assert.NotNil(err) {
		assert.Equal(404, err.Status)
	}
}

func TestDeleteClientSuccess(t *testing.T) {
	assert := assert.New(t)
	suite := SetupClientSuite()

	err := suite.clientModel.Delete(suite.client1.ID)

	assert.Nil(err)
}
