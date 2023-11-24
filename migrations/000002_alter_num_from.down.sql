begin;
ALTER TABLE variant
    RENAME COLUMN num TO num_from;
commit ;