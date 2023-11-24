begin;
ALTER TABLE variant
    RENAME COLUMN num_from TO num;
commit ;