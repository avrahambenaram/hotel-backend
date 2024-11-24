package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/avrahambenaram/hotel-backend/internal/entity"
	"github.com/avrahambenaram/hotel-backend/internal/middleware"
	"github.com/avrahambenaram/hotel-backend/internal/model"
	"github.com/avrahambenaram/hotel-backend/internal/repository"
)

type RoomController struct {
	Handler   http.Handler
	roomModel *model.RoomModel
}

func NewRoomController(roomModel *model.RoomModel) *RoomController {
	mux := http.NewServeMux()
	roomController := &RoomController{
		middleware.SendJSON(mux),
		roomModel,
	}

	mux.Handle(
		"GET /{$}",
		http.HandlerFunc(roomController.query),
	)
	mux.Handle(
		"GET /id/{ID}",
		middleware.GetId(
			http.HandlerFunc(roomController.findByID),
		),
	)
	mux.Handle(
		"GET /number/{number}",
		http.HandlerFunc(roomController.findByNumber),
	)
	mux.Handle(
		"POST /add",
		middleware.ParseBody(
			http.HandlerFunc(roomController.addRoom),
			entity.HotelRoom{},
		),
	)
	mux.Handle(
		"PUT /update",
		middleware.ParseBody(
			http.HandlerFunc(roomController.update),
			entity.HotelRoom{},
		),
	)
	mux.Handle(
		"DELETE /{ID}",
		middleware.GetId(
			http.HandlerFunc(roomController.delete),
		),
	)

	return roomController
}

func (c *RoomController) query(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	capacityStr := query.Get("capacity")
	typeStr := query.Get("type")
	priceDiaryStr := query.Get("priceDiary")
	var capacity uint
	var roomType uint
	var priceDiary float32

	if capacityStr != "" {
		capacityConverted, _ := strconv.Atoi(capacityStr)
		capacity = uint(capacityConverted)
	}
	if typeStr != "" {
		typeConverted, _ := strconv.Atoi(typeStr)
		roomType = uint(typeConverted)
	}
	if priceDiaryStr != "" {
		priceConverted, _ := strconv.ParseFloat(priceDiaryStr, 32)
		priceDiary = float32(priceConverted)
	}
	if capacity != 0 || roomType != 0 || priceDiary != 0 {
		rooms := c.roomModel.FindByQuery(repository.RoomQuery{
			Capacity:   capacity,
			Type:       roomType,
			PriceDiary: priceDiary,
		})
		ctx := context.WithValue(r.Context(), "json", rooms)
		*r = *r.WithContext(ctx)
		return
	}

	rooms := c.roomModel.FindAll()

	ctx := context.WithValue(r.Context(), "json", rooms)
	*r = *r.WithContext(ctx)
}

func (c *RoomController) findByID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(uint)

	room, err := c.roomModel.FindByID(id)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", room)
	*r = *r.WithContext(ctx)
}

func (c *RoomController) findByNumber(w http.ResponseWriter, r *http.Request) {
	pathNumber := r.PathValue("number")
	number, err := strconv.Atoi(pathNumber)
	if err != nil {
		http.Error(w, "Insert a valid number", http.StatusForbidden)
		return
	}

	room, errFind := c.roomModel.FindByNumber(number)
	if errFind != nil {
		http.Error(w, errFind.Message, errFind.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", room)
	*r = *r.WithContext(ctx)
}

func (c *RoomController) addRoom(w http.ResponseWriter, r *http.Request) {
	room := r.Context().Value("body").(entity.HotelRoom)
	roomCreated, err := c.roomModel.Save(room)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", roomCreated)
	*r = *r.WithContext(ctx)
}

func (c *RoomController) update(w http.ResponseWriter, r *http.Request) {
	room := r.Context().Value("body").(entity.HotelRoom)
	roomUpdated, err := c.roomModel.Update(room)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	ctx := context.WithValue(r.Context(), "json", roomUpdated)
	*r = *r.WithContext(ctx)
}

func (c *RoomController) delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(uint)

	errDelete := c.roomModel.Delete(id)
	if errDelete != nil {
		http.Error(w, errDelete.Message, errDelete.Status)
		return
	}

	w.WriteHeader(204)
}
