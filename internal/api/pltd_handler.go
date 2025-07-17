package api

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type PLTDHandler struct {}

func NewPLTDHandler() *PLTDHandler {
	return &PLTDHandler{}
}

func (ph *PLTDHandler) HandleGetPLTDByID(w http.ResponseWriter, r *http.Request) {
	paramsPLTDID := chi.URLParam(r, "id")
	if paramsPLTDID == "" {
		http.NotFound(w, r)
		return
	}

	pltdID, err := strconv.ParseInt(paramsPLTDID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "This data pltd with id %d\n", pltdID)
} 

func (ph *PLTDHandler) HandleCreatePLTD(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a PLTD mesin")
}
