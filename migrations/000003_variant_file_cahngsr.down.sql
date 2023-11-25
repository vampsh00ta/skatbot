begin;
alter table variant drop if  exists column file_path ;
alter table variant drop if  exists column file_id ;
commit;