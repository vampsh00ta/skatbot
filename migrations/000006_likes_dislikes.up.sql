begin;
alter table variant add column if not exists likes integer;
alter table variant add column if not exists dislikes integer;
commit;