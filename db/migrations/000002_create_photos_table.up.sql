begin;

create table if not exists public.photos (
    photoId serial primary key, 
    title varchar (100) not null,
    caption varchar (100) not null,
    photo_url varchar(100) not null,
    user_id int not null references users(userId),
    created_at timestamp not null,
    updated_at timestamp not null
);

commit;