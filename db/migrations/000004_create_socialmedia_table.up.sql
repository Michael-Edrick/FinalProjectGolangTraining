begin;

create table if not exists public.socialmedia (
    scId serial primary key, 
    name varchar(100) not null,
    social_media_url varchar(100) not null,
    user_id int not null references users(userId)
);

commit;