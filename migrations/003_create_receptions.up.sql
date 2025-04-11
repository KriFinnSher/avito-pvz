CREATE TABLE IF NOT EXISTS receptions (
                            id UUID PRIMARY KEY,
                            date_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            pvz_id UUID NOT NULL REFERENCES pvz(id) ON DELETE CASCADE,
                            status TEXT NOT NULL CHECK (status IN ('in_progress', 'close'))
);
CREATE INDEX IF NOT EXISTS idx_receptions_date_time ON receptions(date_time);
CREATE INDEX IF NOT EXISTS idx_receptions_pvz_id ON receptions(pvz_id);