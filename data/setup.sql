drop table if exists setquantities;
drop table if exists sets cascade;
drop table if exists lifts cascade;
drop table if exists workouts cascade;
drop table if exists plans cascade;
drop table if exists workout_lifts;

create table workouts (
    id serial primary key,
    name varchar(255),
    week_no integer,
    date date,
    created_at timestamp not null
);

create table lifts (
    id serial primary key,
    name varchar(255),
    max real,
    created_at timestamp not null
);

create table sets (
    id serial primary key,
    done boolean,
    created_at timestamp not null,
    lift_id integer references lifts(id)
);

create table setquantities (
    id serial primary key,
    rep_type varchar(255) not null,
    quantity integer,
    weight real, 
    planned_ratio int,
    ratio_type varchar(255),
    set_id integer references sets(id),
    created_at timestamp not null 
);

create table workout_lifts (
    id serial primary key,
    workout_id integer references workouts(id),
    lift_id integer references lifts(id)
);