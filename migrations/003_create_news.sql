CREATE TABLE news
(
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    slug        VARCHAR(255) NOT NULL UNIQUE,
    summary     TEXT         NOT NULL,
    content     TEXT         NOT NULL,
    image_url   TEXT,
    status      VARCHAR(20)  NOT NULL DEFAULT 'draft' CHECK ( status IN ('draft', 'published') ),
    category_id BIGINT REFERENCES categories (id),
    author_id   BIGINT REFERENCES users (id),
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);