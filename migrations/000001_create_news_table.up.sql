CREATE TABLE news_articles (
                               id SERIAL PRIMARY KEY,
                               title VARCHAR(255) NOT NULL,
                               content TEXT NOT NULL,
                               category VARCHAR(100) NOT NULL,
                               author VARCHAR(100) NOT NULL,
                               created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                               updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
