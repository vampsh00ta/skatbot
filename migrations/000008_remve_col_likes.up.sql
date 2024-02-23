begin;
alter table variant drop column if exists likes ;
alter table variant drop column if exists dislikes ;
commit;