begin;

    create table if not exists stats(
        id serial primary key ,
        likes integer references variant(id),
        dislikes integer references variant(id)

    );
commit;