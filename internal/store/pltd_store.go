package store

import (
	"fmt"
	"encoding/json"
	"database/sql"
)

type PLTD struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
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
	SELECT id, name, status, efisiensi, batas_beban
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
		var efisiensiByte []byte

		err := rows.Scan(&pltd.ID, &pltd.Name, &pltd.Status, &efisiensiByte, &pltd.BatasBeban)
		if err != nil {
			return nil, err
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

	err = tx.QueryRow(query, pltd.Name, pltd.Status, pltd.Efisiensi, pltd.BatasBeban).Scan(&pltd.ID)
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

	var efisiensiBytes []byte
	query := `
	SELECT id, name, status, efisiensi, batas_beban
	FROM mesin_pltd
	WHERE id=$1
	`

	err := pg.db.QueryRow(query, id).Scan(&pltd.ID, &pltd.Name, &pltd.Status, &efisiensiBytes, &pltd.BatasBeban)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	// Parse efisiensi dari JSON []byte ke map
	if len(efisiensiBytes) > 0 {
		if err := json.Unmarshal(efisiensiBytes, &pltd.Efisiensi); err != nil {
			return nil, fmt.Errorf("gagal unmarshal efisiensi: %w", err)
		}
	} else {
		pltd.Efisiensi = make(map[string]float64) // jika null
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

