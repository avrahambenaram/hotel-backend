package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/middleware"
	"github.com/avrahambenaram/hotel-backend/internal/model"
)

type ClientController struct {
	Handler     http.Handler
	clientModel *model.ClientModel
}

func NewClientController(clientModel *model.ClientModel) *ClientController {
	mux := http.NewServeMux()
	clientController := &ClientController{
		middleware.SendJSON(mux),
		clientModel,
	}

	mux.Handle(
		"GET /id/{ID}",
		http.HandlerFunc(clientController.findByID),
	)
	mux.Handle(
		"GET /cpf/{CPF}",
		http.HandlerFunc(clientController.findByCPF),
	)
	mux.Handle(
		"POST /add",
		middleware.ParseBody(
			http.HandlerFunc(clientController.addClient),
			entity.Client{},
		),
	)
	mux.Handle(
		"PUT /update",
		middleware.ParseBody(
			http.HandlerFunc(clientController.update),
			entity.Client{},
		),
	)

	return clientController
}

func (c *ClientController) findByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("ID")
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		http.Error(w, "ID must be an integer", 403)
	}

	client, err := c.clientModel.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Message, int(err.Status))
		return
	}

	ctx := context.WithValue(r.Context(), "json", client)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) findByCPF(w http.ResponseWriter, r *http.Request) {
	cpf := r.PathValue("CPF")
	client, err := c.clientModel.FindByCPF(cpf)
	if err != nil {
		http.Error(w, err.Message, int(err.Status))
		return
	}

	ctx := context.WithValue(r.Context(), "json", client)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) addClient(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("client").(entity.Client)
	clientCreated, err := c.clientModel.Save(client)
	if err != nil {
		http.Error(w, err.Message, int(err.Status))
		return
	}

	ctx := context.WithValue(r.Context(), "json", clientCreated)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) update(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("client").(entity.Client)
	clientCreated, err := c.clientModel.Update(client)
	if err != nil {
		http.Error(w, err.Message, int(err.Status))
		return
	}

	ctx := context.WithValue(r.Context(), "json", clientCreated)
	*r = *r.WithContext(ctx)
}
