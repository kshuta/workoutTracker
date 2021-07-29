drop table if exists setquantities;
drop table if exists sets cascade;
drop table if exists lifts cascade;
drop table if exists workouts cascade;
drop table if exists plans cascade;

create table plans(
    id serial primary key,
    name varchar(255),
    duration integer,
    frequency integer,
    created_at timestamp not null
);

create table workouts (
    id serial primary key,
    name varchar(255),
    week_no integer,
    date date,
    plan_id integer references plans(id),
    created_at timestamp not null
);

create table lifts (
    id serial primary key,
    name varchar(255),
    max real,
    workout_id integer references workouts(id),
    created_at timestamp not null
);

create table sets (
    id serial primary key,
    done boolean,
    lift_id integer references lifts(id),
    created_at timestamp not null
);

create table setquantities (
    id serial primary key,
    repType varchar(255) not null,
    quantity integer not null,
    weight real not null,
    planned_ratio int,
    ratio_type varchar(255),
    set_id integer references sets(id),
    created_at timestamp not null 
);
