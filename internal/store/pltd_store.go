package store

import (
	"fmt"
	"encoding/json"
	"database/sql"
)

type StatusMesin int

const (
	Tersedia StatusMesin = iota
	Gangguan
	Pemeliharaan
)

func (sm StatusMesin) String() string {
	return [...]string{"tersedia", "gangguan", "pemeliharaan"}[sm];
}

func (s *StatusMesin) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}

	switch statusStr {
	case "tersedia":
		*s = Tersedia
	case "gangguan":
		*s = Gangguan
	case "pemeliharaan":
		*s = Pemeliharaan
	default:
		return fmt.Errorf("unknown status: %s", statusStr)
	}
	return nil
}

func (s StatusMesin) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

type PLTD struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Status StatusMesin `json:"status"`
	Efisiensi map[string]float64 `json:"efisiensi"`
	BatasBeban int `json:"batas_beban"`
}

type PostgresPLTDStore struct {
	db *sql.DB
}

func NewPostgresPLTDStore(db *sql.DB) *PostgresPLTDStore {
	return &PostgresPLTDStore{db: db}
}

type PLTDStore interface {
	CreatePLTD(*PLTD) (*PLTD, error)
	GetPLTDByID(id int64) (*PLTD, error)
}

func (pg *PostgresPLTDStore) CreatePLTD(pltd *PLTD) (*PLTD, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
	INSERT INTO mesin_pltd (name, status, efisiensi, batas_beban)
	VALUES($1, $2, $3, $4)
	RETURNING id
	`

	err = tx.QueryRow(query, pltd.Name, pltd.Status.String(), pltd.Efisiensi, pltd.BatasBeban).Scan(&pltd.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return pltd, nil
}

func (pg *PostgresPLTDStore) GetPLTDByID(id int64) (*PLTD, error) {
	pltd := &PLTD{}
	return pltd, nil
}
