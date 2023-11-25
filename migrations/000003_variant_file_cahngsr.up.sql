begin;
alter table variant add column if not exists file_path varchar(255);
alter table variant add column if not exists file_id varchar(255);
commit;
