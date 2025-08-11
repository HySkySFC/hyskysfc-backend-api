--+goose Up
--+goose StatementBegin
CREATE TABLE IF NOT EXISTS plts (
    time TIMESTAMP NOT NULL UNIQUE,
    weight DOUBLE PRECISION NOT NULL
)
--+goose StatementEnd

--+goose Down
--+goose StatementBegin
DROP TABLE plts;
--+goose StatementEnd
