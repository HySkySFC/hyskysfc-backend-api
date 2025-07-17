--+goose Up
--+goose StatementBegin
CREATE TYPE status_mesin 
AS ENUM ('gangguan', 'pemeliharaan', 'tersedia');

CREATE TABLE IF NOT EXISTS mesin_pltd (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status status_mesin NOT NULL,
    efisiensi JSON NOT NULL,
    batas_beban INTEGER NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
--+goose StatementEnd

--+goose Down
--+goose StatementBegin
DROP TABLE mesin_pltd;
DROP TYPE IF EXISTS status_mesin;
--+goose StatementEnd
