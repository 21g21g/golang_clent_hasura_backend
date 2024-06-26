alter table "results"."episode" alter column "created" drop not null;
alter table "results"."episode" add column "created" timestamptz;
