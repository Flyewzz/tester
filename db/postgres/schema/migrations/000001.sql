create table users
(
  id       serial       not null
    constraint users_pk
      primary key,
  login    varchar(20)  not null,
  email    varchar(200) not null,
  password varchar(64)  not null,
  name     varchar(100)
);

alter table users
  owner to postgres;

create unique index users_email_uindex
  on users (email);

create unique index users_nick_uindex
  on users (login);

create table tasks
(
  id          serial  not null
    constraint tasks_1_pk
      primary key,
  text        text    not null,
  ram         integer not null,
  hdd         integer,
  time        integer not null,
  samples     text,
  limitations text
);

alter table tasks
  owner to postgres;

create table attempts
(
  id      serial                  not null
    constraint attempts_pk
      primary key,
  user_id integer                 not null
    constraint attempts_users_id_fk
      references users,
  task_id integer
    constraint attempts_tasks_id_fk
      references tasks,
  status  varchar(10)             not null,
  time    timestamp default now() not null
);

alter table attempts
  owner to postgres;