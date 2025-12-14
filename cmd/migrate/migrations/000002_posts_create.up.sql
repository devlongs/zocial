CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    user_id bigint NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NOW() NOT NULL
);