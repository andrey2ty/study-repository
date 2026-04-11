package dockerpractice

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PeopleDto struct {
	FullName string `json:"fullname"`
	Position string `json:"position"`
}

func HttpHandleAddPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	peopled := PeopleDto{}

	if err := json.NewDecoder(r.Body).Decode(&peopled); err != nil {
		return
	}

	defer r.Body.Close()

	people := NewPeople(peopled.FullName, peopled.Position)

	PeopleM[people.ID] = people

}

func HttpHandleGetPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	data, _ := json.MarshalIndent(PeopleM, "", "    ")

	w.Write(data)
}

func HttpHandleDeletePeople(w http.ResponseWriter, r *http.Request) {

	idstr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idstr)

	v, ok := PeopleM[id]
	if ok != true {
		return
	}

	delete(PeopleM, id)

	data, _ := json.Marshal(v)

	w.Write(data)
}
