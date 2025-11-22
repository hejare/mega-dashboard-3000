-- +goose Up
ALTER TABLE consultants ADD profile_image_url TEXT;
ALTER TABLE consultants ADD probable_extension_status TEXT;

-- +goose Down
ALTER TABLE consultants DROP COLUMN profile_image_url;
ALTER TABLE consultants DROP COLUMN probable_extension_status;