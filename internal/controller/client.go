package controller

import (
	"context"
	"net/http"

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
		"GET /{$}",
		http.HandlerFunc(clientController.findAll),
	)
	mux.Handle(
		"GET /id/{ID}",
		middleware.GetId(
			http.HandlerFunc(clientController.findByID),
		),
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
	mux.Handle(
		"DELETE /{ID}",
		middleware.GetId(
			http.HandlerFunc(clientController.delete),
		),
	)

	return clientController
}

func (c *ClientController) findAll(w http.ResponseWriter, r *http.Request) {
	clients := c.clientModel.FindAll()

	ctx := context.WithValue(r.Context(), "json", clients)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) findByID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(uint)

	client, err := c.clientModel.FindByID(id)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", client)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) findByCPF(w http.ResponseWriter, r *http.Request) {
	cpf := r.PathValue("CPF")
	client, err := c.clientModel.FindByCPF(cpf)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", client)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) addClient(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("body").(entity.Client)
	clientCreated, err := c.clientModel.Save(client)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", clientCreated)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) update(w http.ResponseWriter, r *http.Request) {
	client := r.Context().Value("body").(entity.Client)
	clientCreated, err := c.clientModel.Update(client)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", clientCreated)
	*r = *r.WithContext(ctx)
}

func (c *ClientController) delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(uint)

	errDelete := c.clientModel.Delete(id)
	if errDelete != nil {
		http.Error(w, errDelete.Message, errDelete.Status)
		return
	}

	w.WriteHeader(204)
}
