package store

import (
	"database/sql"
	"time"
)

type PLTS struct {
	Time time.Time `json:"time"`
	Weight float64 `json:"weight"`
}

type PostgresPLTSStore struct {
	db *sql. DB
}

func NewPostgresPLTSStore(db *sql.DB) *PostgresPLTSStore {
	return &PostgresPLTSStore{
		db: db,
	}
}

type PLTSStore interface {
	GetAllPLTS() ([]*PLTS, error)
	ReplaceAllPLTS([]*PLTS) error
}

func (pg *PostgresPLTSStore) GetAllPLTS() ([]*PLTS, error) {
	rows, err := pg.db.Query(`SELECT time, weight FROM plts ORDER BY time ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*PLTS, 0)
	for rows.Next() {
		p := &PLTS{}
		if err := rows.Scan(&p.Time, &p.Weight); err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return list, rows.Err()
}

func (pg *PostgresPLTSStore) ReplaceAllPLTS(pltsList []*PLTS) error {
	tx, err := pg.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`DELETE FROM plts`); err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO plts (time, weight) VALUES ($1, $2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range pltsList {
		if _, err := stmt.Exec(p.Time, p.Weight); err != nil {
			return err
		}
	}

	return tx.Commit()
}


