begin;
    create table if not exists variant_likes(
        username varchar(255),
        variant_id integer references variant(id),
        l  integer,
        d integer,
        primary key(username,variant_id)
    );

commit;