alter table "results"."character" alter column "created" drop not null;
alter table "results"."character" add column "created" timestamptz;
