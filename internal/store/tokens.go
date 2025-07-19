package store

import (
	"time"
	"database/sql"
	"github.com/HySkySFC/hyskysfc-backend-api/internal/tokens"
)

type PostgresTokenStore struct {
	db *sql.DB
}

func NewPostgresTokenStore(db *sql.DB) *PostgresTokenStore {
	return &PostgresTokenStore{
		db: db,
	}
}

type TokenStore interface {
	Insert(token *tokens.Token) error
	CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error)
	DeleteAllTokensForUser(userID int, scope string) error
}

func (t *PostgresTokenStore) CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error) {
	token, err := tokens.GenerateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = t.Insert(token)
	return token, err
}

func (t *PostgresTokenStore) Insert(token *tokens.Token) error {
	query := `
	INSERT INTO tokens (hash, user_id, expire, scope)
	VALUES ($1, $2, $3, $4)
	`

	_, err := t.db.Exec(query, token.Hash, token.UserID, token.Expire, token.Scope)
	return err
}

func (t *PostgresTokenStore) DeleteAllTokensForUser(userID int, scope string) error {
	query := `
	DELETE FROM tokens
	WHERE scope = $1 AND user_id = $2
	`
	_, err := t.db.Exec(query, scope, userID)
	return err
}
