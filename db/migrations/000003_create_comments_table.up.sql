begin;

create table if not exists public.comments (
    commentId serial primary key, 
    user_id int not null references users(userId),
    photo_id int not null references photos(photoId),
    message varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

commit;