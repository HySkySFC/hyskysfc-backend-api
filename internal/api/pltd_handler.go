package api

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/store"
)

type PLTDHandler struct {
	pltdStore store.PLTDStore
}

func NewPLTDHandler(pltdStore store.PLTDStore) *PLTDHandler {
	return &PLTDHandler{
		pltdStore: pltdStore,
	}
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
	var pltd store.PLTD
	err := json.NewDecoder(r.Body).Decode(&pltd)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create pltd", http.StatusInternalServerError)
		return
	}

	createdPLTD, err := ph.pltdStore.CreatePLTD(&pltd)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create pltd", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPLTD)
}
