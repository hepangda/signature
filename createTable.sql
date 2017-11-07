create database signature;
use signature;

create table sign (
    id SERIAL NOT NULL,
    user_id TEXT NOT NULL,
    user_name TEXT NOT NULL,
    state TEXT NOT NULL,
    sign_time TIMESTAMP NOT NULL,
    PRIMARY KEY(id)
);