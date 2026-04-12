package dockerpractice

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type HttpHandler struct {
	ctx  context.Context
	conn *pgx.Conn
}

func NewHttpHandler(ctx context.Context, conn *pgx.Conn) HttpHandler {
	return HttpHandler{
		ctx:  ctx,
		conn: conn,
	}
}

type PeopleDto struct {
	FullName string `json:"fullname"`
	Position string `json:"position"`
}

func (h *HttpHandler) HttpHandleAddPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	peopled := PeopleDto{}

	if err := json.NewDecoder(r.Body).Decode(&peopled); err != nil {
		return
	}

	defer r.Body.Close()

	people := NewPeople(peopled.FullName, peopled.Position)

	peoplemodel := NewPeopleModel(people.ID, people.FullName, people.Position)

	InsertPeople(h.conn, h.ctx, peoplemodel)

}

func (h *HttpHandler) HttpHandleGetPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	peoplearr, err := SelectPeople(h.conn, h.ctx)

	if err != nil {
		return
	}

	data, _ := json.MarshalIndent(peoplearr, "", "    ")

	w.Write(data)
}

func (h *HttpHandler) HttpHandleDeletePeople(w http.ResponseWriter, r *http.Request) {

	idstr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idstr)

	if err := DeletePeople(h.conn, h.ctx, id); err != nil {
		return
	}

	w.Write([]byte("успешно"))
}
