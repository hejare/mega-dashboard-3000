-- +goose Up
ALTER TABLE consultants ADD profile_imaga_url TEXT;
ALTER TABLE consultants ADD probable_extension_status TEXT;

-- +goose Down
ALTER TABLE consultants DROP COLUMN profile_imaga_url;
ALTER TABLE consultants DROP COLUMN probable_extension_status;