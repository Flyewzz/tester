alter table users alter column login type varchar(25) using login::varchar(25);

alter table users alter column email type varchar(254) using email::varchar(254);

alter table users alter column name type varchar(80) using name::varchar(80);