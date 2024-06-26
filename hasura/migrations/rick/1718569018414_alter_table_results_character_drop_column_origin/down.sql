alter table "results"."character" alter column "origin" drop not null;
alter table "results"."character" add column "origin" jsonb;
