begin;
create table if not exists teacher(
      id serial primary key ,
      first_name varchar(50),
    second_name varchar(50),
    third_name varchar(50)
    );

create table if not exists subject_type(
                                      id serial ,
                                      name varchar(50) primary key
    );
create table if not exists variant_type(
                                           id serial ,
                                           name varchar(50) primary key
    );
create table if not exists semester(
                                       id serial,
                                       number integer primary key
);

create table if not exists subject(
                                      id serial  ,
                                      name varchar(50) primary key
    );
create table if not exists instistute(
                                         id serial  ,
                                         number integer primary key
);
create table if not exists active_subject(
                                             id serial primary key ,
                                             name varchar(50) references subject(name) on delete CASCADE,
    semester_number integer references semester(number) on delete CASCADE,
    instistute_num integer references instistute(number) on delete CASCADE,
    type_name varchar(50) references subject_type(name) on delete CASCADE
    );

create table if not exists variant(
      id serial primary key ,
      subject_id integer references active_subject(id) on delete CASCADE,
    name varchar(50),
    num_from integer null,
    grade integer  null ,
    creation_time timestamp not null ,
    type_name varchar(50) references variant_type(name) on delete CASCADE
    );




commit;

