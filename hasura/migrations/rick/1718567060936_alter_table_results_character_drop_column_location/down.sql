alter table "results"."character" alter column "location" drop not null;
alter table "results"."character" add column "location" jsonb;
