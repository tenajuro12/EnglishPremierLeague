CREATE TABLE matches (
                         id SERIAL PRIMARY KEY,              -- Unique integer ID for the match
                         home_team VARCHAR(100) NOT NULL,   -- Home team name
                         away_team VARCHAR(100) NOT NULL,   -- Away team name
                         date TIMESTAMP NOT NULL,           -- Match date
                         status VARCHAR(50) NOT NULL,      -- Match status
                         home_score INT DEFAULT 0,         -- Home team score, default to 0
                         away_score INT DEFAULT 0,        -- Away team score, default to 0
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp when the match record was created
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp when the match record was last updated
);

CREATE INDEX idx_matches_date ON matches(date);
