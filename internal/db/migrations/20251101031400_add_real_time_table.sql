-- +goose Up
create table if not exists quotes (
    id serial primary key,
    symbol varchar(255) not null,
    company_name varchar(255) not null,
    stock_type varchar not null,
    exchange varchar not null,
    primary_data jsonb not null,
    secondary_data jsonb,
    market_status varchar not null,
    assetClass varchar not null,
    key_stats jsonb not null
);

-- +goose Down
drop table if exists quotes;
