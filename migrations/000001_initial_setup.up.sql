CREATE TABLE IF NOT EXISTS Drivers(
    id varchar(256),
    fullname varchar(64),
    phone varchar(16),
    deleted_at timestamp default null
);

CREATE TABLE IF NOT EXISTS Clients(
    id varchar(256),
    fullname varchar(64),
    phone varchar(16),
    deleted_at timestamp default null
);

CREATE TABLE IF NOT EXISTS Orders(
    id varchar(256),
    driver_id varchar(256),
    client_id varchar(256),
    status varchar(12),
    created_at timestamp default current_timestamp,
    updated_at timestamp default null,
    deleted_at timestamp default null
);