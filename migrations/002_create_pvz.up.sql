CREATE TABLE IF NOT EXISTS pvz (
                     id UUID PRIMARY KEY,
                     registration_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                     city TEXT NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань'))
);
CREATE INDEX IF NOT EXISTS idx_pvz_registration_date ON pvz(registration_date);