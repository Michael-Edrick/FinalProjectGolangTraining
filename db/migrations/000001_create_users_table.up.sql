begin;

create table if not exists public.users (
    userId serial primary key, 
    username varchar (100) unique not null,
    email varchar (100) unique not null,
    password varchar(100) not null,
    age int not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

commit;