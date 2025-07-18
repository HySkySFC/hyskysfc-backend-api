package api

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"database/sql"
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

func (ph *PLTDHandler) HandleGetAllPLTD(w http.ResponseWriter, r *http.Request) {
	pltdList, err := ph.pltdStore.GetAllPLTD()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch pltd", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pltdList)
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

	pltd, err := ph.pltdStore.GetPLTDByID(pltdID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to read the pltd", http.StatusNotFound)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pltd)
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

func (ph *PLTDHandler) HandleUpdatePLTDByID(w http.ResponseWriter, r *http.Request) {
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

	existingPLTD, err := ph.pltdStore.GetPLTDByID(pltdID)
	if err != nil {
		http.Error(w, "Failed to fetch pltd", http.StatusInternalServerError)
		return
	}

	if existingPLTD == nil {
		http.NotFound(w, r)
		return
	}

	var updatePLTDRequest struct {
		Name *string `json:"name"`
		Status *string `json:"status"`
		Efisiensi *map[string]float64 `json:"efisiensi"`
		BatasBeban *int `json:"batas_beban"`
	}

	err = json.NewDecoder(r.Body).Decode(&updatePLTDRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatePLTDRequest.Name != nil {
		existingPLTD.Name = *updatePLTDRequest.Name
	}

	if updatePLTDRequest.Status != nil {
		existingPLTD.Status = *updatePLTDRequest.Status
	}
	
	if updatePLTDRequest.Efisiensi != nil {
		existingPLTD.Efisiensi = *updatePLTDRequest.Efisiensi
	}

	if updatePLTDRequest.BatasBeban != nil {
		existingPLTD.BatasBeban = *updatePLTDRequest.BatasBeban
	}

	err = ph.pltdStore.UpdatePLTD(existingPLTD)
	if err != nil {
		fmt.Println("update workout error", err)
		http.Error(w, "failed to update the pltd", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingPLTD)
}

func (ph *PLTDHandler) HandleDeletePLTDByID(w http.ResponseWriter, r *http.Request) {
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

	err = ph.pltdStore.DeletePLTD(pltdID)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		http.Error(w, "PLTD not found", http.StatusNotFound)
		return
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "error deleting workout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

