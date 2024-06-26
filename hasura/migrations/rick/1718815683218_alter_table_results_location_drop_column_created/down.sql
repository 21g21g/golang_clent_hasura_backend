alter table "results"."location" alter column "created" drop not null;
alter table "results"."location" add column "created" timestamptz;
