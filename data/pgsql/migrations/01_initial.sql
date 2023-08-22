-- +migrate Up 
create table urls (
    url_id serial primary key,
    full_url text not null unique
);

-- +migrate Down
drop table urls cascade;