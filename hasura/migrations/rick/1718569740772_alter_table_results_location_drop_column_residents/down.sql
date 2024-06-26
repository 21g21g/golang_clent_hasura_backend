alter table "results"."location" alter column "residents" drop not null;
alter table "results"."location" add column "residents" _text;
