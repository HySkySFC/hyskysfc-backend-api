package api

import (
	"log"
	"net/http"
	"encoding/json"
	"database/sql"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/store"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/utils"
)

type PLTDHandler struct {
	pltdStore store.PLTDStore
	logger *log.Logger
}

func NewPLTDHandler(pltdStore store.PLTDStore, logger *log.Logger) *PLTDHandler {
	return &PLTDHandler{
		pltdStore: pltdStore,
		logger: logger,
	}
}

func (ph *PLTDHandler) HandleGetAllPLTD(w http.ResponseWriter, r *http.Request) {
	pltdList, err := ph.pltdStore.GetAllPLTD()
	if err != nil {
		ph.logger.Printf("ERROR: getAllPLTD: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"pltd": pltdList})
}

func (ph *PLTDHandler) HandleGetPLTDByID(w http.ResponseWriter, r *http.Request) {
	pltdID, err := utils.ReadIDParam(r)
	if err != nil {
		ph.logger.Printf("ERROR: readIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid PLTD id"})
		return
	}

	pltd, err := ph.pltdStore.GetPLTDByID(pltdID)
	if err != nil {
		ph.logger.Printf("ERROR: getPLTDByID: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "PLTD not found."})
		return 
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"pltd": pltd})
} 

func (ph *PLTDHandler) HandleCreatePLTD(w http.ResponseWriter, r *http.Request) {
	var pltd store.PLTD
	err := json.NewDecoder(r.Body).Decode(&pltd)
	if err != nil {
		ph.logger.Printf("ERROR: decode PLTD: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	createdPLTD, err := ph.pltdStore.CreatePLTD(&pltd)
	if err != nil {
		ph.logger.Printf("ERROR: create PLTD: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid request payload"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"pltd": createdPLTD})
}

func (ph *PLTDHandler) HandleUpdatePLTDByID(w http.ResponseWriter, r *http.Request) {
	pltdID, err := utils.ReadIDParam(r)
	if err != nil {
		ph.logger.Printf("ERROR: readIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid PLTD id"})
		return
	}

	existingPLTD, err := ph.pltdStore.GetPLTDByID(pltdID)
	if err != nil {
		ph.logger.Printf("ERROR: get PLTD by ID: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	if existingPLTD == nil {
		ph.logger.Printf("ERROR: existingPLTD: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "PLTD Not Found"})
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
		ph.logger.Printf("ERROR: updatePLTD: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"pltd": existingPLTD})
}

func (ph *PLTDHandler) HandleDeletePLTDByID(w http.ResponseWriter, r *http.Request) {
	pltdID, err := utils.ReadIDParam(r)
	if err != nil {
		ph.logger.Printf("ERROR: readIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid PLTD id"})
		return
	}

	err = ph.pltdStore.DeletePLTD(pltdID)
	if err == sql.ErrNoRows {
		ph.logger.Printf("ERROR: deletePLTD: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "PLTD not found"})
		return
	}

	if err != nil {
		ph.logger.Printf("ERROR: deletePLTD: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, utils.Envelope{"pltd": "PLTD successfully deleted"})
}

