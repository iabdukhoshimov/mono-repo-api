-- +goose Up
-- +goose StatementBegin
create type item_status as enum ('active', 'inactive');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type item_status;
-- +goose StatementEnd
