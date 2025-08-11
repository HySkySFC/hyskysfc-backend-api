package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HySkySFC/hyskysfc-backend-api/internal/store"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/utils"
)

type PLTSHandler struct {
	pltsStore store.PLTSStore
	logger *log.Logger
}

func NewPLTSHandler(pltsStore store.PLTSStore, logger *log.Logger) *PLTSHandler {
	return &PLTSHandler{
		pltsStore: pltsStore,
		logger: logger,
	}
}

func (ph *PLTSHandler) HandleGetAllPLTS(w http.ResponseWriter, r *http.Request) {
	newData, err := ph.pltsStore.GetAllPLTS()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"data": newData})
}

func (ph *PLTSHandler) HandleReplaceAllPLTS(w http.ResponseWriter, r *http.Request) {
	var payload []*store.PLTS
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid payload"})
		return
	}

	if err := ph.pltsStore.ReplaceAllPLTS(payload); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	newData, err := ph.pltsStore.GetAllPLTS()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"data": newData})
}

