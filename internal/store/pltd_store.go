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
	GetAllPLTD() ([]*PLTD, error)
	CreatePLTD(*PLTD) (*PLTD, error)
	GetPLTDByID(id int64) (*PLTD, error)
	UpdatePLTD(*PLTD) error
	DeletePLTD(id int64) error
}

func (pg *PostgresPLTDStore) GetAllPLTD() ([]*PLTD, error) {
	query := `
	SELECT (id, name, status, efisiensi, batas_beban)
	FROM mesin_pltd
	`

	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pltdList []*PLTD

	for rows.Next() {
		var pltd PLTD
		var statusStr string
		var efisiensiByte []byte

		err := rows.Scan(&pltd.ID, &pltd.Name, &statusStr, &efisiensiByte, &pltd.BatasBeban)
		if err != nil {
			return nil, err
		}

		switch statusStr {
		case "tersedia":
			pltd.Status = Tersedia
		case "gangguan":
			pltd.Status = Gangguan
		case "pemeliharaan":
			pltd.Status = Pemeliharaan
		default:
			return nil, fmt.Errorf("Invalid status value: %s", statusStr)
		}

		err = json.Unmarshal(efisiensiByte, &pltd.Efisiensi)
		if err != nil {
			return nil, err
		}

		pltdList = append(pltdList, &pltd)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pltdList, nil
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

	query := `
	SELECT (id, name, status, efisiensi, batas_beban)
	FROM mesin_pltd
	WHERE id=$1
	`

	err := pg.db.QueryRow(query, id).Scan(&pltd.ID, &pltd.Name, &pltd.Status, &pltd.Efisiensi, &pltd.BatasBeban)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return pltd, nil
}

func (pg *PostgresPLTDStore) UpdatePLTD(pltd *PLTD) error {
	tx, err := pg.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
	UPDATE mesin_pltd
	SET name = $1, status = $2, efisiensi = $3, batas_beban = $4
	WHERE id=$5
	`

	result, err := tx.Exec(query, pltd.Name, pltd.Status, pltd.Efisiensi, pltd.BatasBeban, pltd.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return tx.Commit()
}

func (pg *PostgresPLTDStore) DeletePLTD(id int64) error {
	query := `
	DELETE FROM mesin_pltd
	WHERE id=$1
	`
	result, err := pg.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

