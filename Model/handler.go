package Model

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"wb-l0/handlers"
)

type handler struct {
	cache map[string]Order
}

func NewHandler(cache map[string]Order) handlers.Handler {
	return &handler{
		cache: cache,
	}
}
func (h *handler) Register(router *httprouter.Router) {
	router.GET("/", h.GetAllId)
	router.GET("/orders", h.GetById)
}

func (h *handler) GetAllId(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	tmpl, err := template.ParseFiles("templates/find_by_orderId.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := r.URL.Query().Get("id")
	val, exists := h.cache[id]
	if exists == true {
		jsonData, _ := json.Marshal(val)
		w.Write(jsonData)
	} else {
		w.Write([]byte("This is order doesn't exists"))
	}

}
