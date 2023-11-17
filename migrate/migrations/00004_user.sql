-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id serial PRIMARY KEY NOT NULL,
    name varchar NOT NULL,
    email varchar NOT NULL,
    created_at timestamp WITH time zone NOT NULL DEFAULT (NOW()),
    updated_at timestamp WITH time zone NOT NULL DEFAULT (NOW()),
    deleted_at timestamp WITH time zone DEFAULT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "user";

-- +goose StatementEnd