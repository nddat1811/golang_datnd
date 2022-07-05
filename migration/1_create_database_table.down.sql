-- +migrate Down
-- SQL in section 'Up' is executed when this migration is applied

DROP TABLE golang_api.users;
DROP TABLE golang_api.books