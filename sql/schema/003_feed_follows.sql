-- +goose Up
CREATE TABLE feed_follows(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE ON UPDATE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE(user_id, feed_id),

    CONSTRAINT fk_ff_users
    FOREIGN KEY(user_id)
    REFERENCES users(id),

    CONSTRAINT fk_ff_feeds
    FOREIGN KEY(feed_id)
    REFERENCES feeds(id)

);

-- +goose Down
DROP TABLE feed_follows;