CREATE TABLE IF NOT EXISTS threads (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    user_id integer 
        CONSTRAINT threads_user_id_fk
            REFERENCES users
            ON UPDATE CASCADE 
            ON DELETE SET NULL,
    channel_id integer
        CONSTRAINT threads_channel_id_fk
            REFERENCES channels
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
)