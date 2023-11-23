begin;
    create table if not exists subject_type(
         id serial primary key ,
         name varchar(50)
    );
    create table if not exists variant_type(
        id serial primary key ,
        name varchar(50)
    );
    create table if not exists subject(
        id serial primary key ,
        name varchar(50),
        semester integer,
        type integer references subject_type(id)
    );
    create table if not exists variant(
        id serial primary key ,
        subject_id integer references subject(id) on delete CASCADE,
        name varchar(50),
        num_from integer,
        grade integer  null ,
        creation_time timestamp not null ,
        type integer references variant_type(id)
    );

commit;
