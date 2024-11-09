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
		"POST /add",
		middleware.ParseBody(
			http.HandlerFunc(clientController.addClient),
		),
	)

	return clientController
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
