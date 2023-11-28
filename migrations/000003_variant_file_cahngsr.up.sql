begin;
alter table variant add column if not exists file_path varchar(255);
alter table variant add column if not exists file_id varchar(255);
alter table variant add column if not exists tg_username varchar(255);
alter table variant add column if not exists tg_id varchar(255);
commit;
