-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table IF NOT EXISTS users
(
	id uuid constraint users_pk primary key DEFAULT uuid_generate_v4(),
	first_name varchar(100) not null,
    last_name varchar(100) not null,
    role varchar(100) not null,
	email varchar(100) not null UNIQUE,
	password_hash text not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

create unique index users_email_uindex on users (email);

create table IF NOT EXISTS products
(
	id uuid constraint products_pk primary key DEFAULT uuid_generate_v4(),
	slug varchar(256) not null,
    name varchar(256) not null,
    description text not null,
	price bigint not null,
    discount bigint not null,
    available_quantity bigint not null,
	currency varchar(256) not null,
    status varchar(256) not null,
	created_at timestamp default current_timestamp not null,
    deleted_at timestamp default null,
	updated_at timestamp default null
);

create table IF NOT EXISTS orders
(
	id uuid constraint orders_pk primary key DEFAULT uuid_generate_v4(),
	user_id uuid not null,
    tracking_code varchar(256) not null,
    status varchar(100) not null,
	fee bigint not null,
    history jsonb not null,
	currency varchar(256) not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

create table IF NOT EXISTS order_records
(
	id uuid constraint order_records_pk primary key DEFAULT uuid_generate_v4(),
	product_id uuid not null,
    order_id uuid not null,
	quantity bigint not null,
    amount bigint not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "order_records" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP Table users;
DROP Table products;
DROP Table orders;
DROP Table order_records;
-- +goose StatementEnd
