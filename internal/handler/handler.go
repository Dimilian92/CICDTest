package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandle() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It's allive!!!")

}
