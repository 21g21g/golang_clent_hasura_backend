alter table "results"."episode" alter column "characters" drop not null;
alter table "results"."episode" add column "characters" _text;
