alter table "results"."character" alter column "episode" drop not null;
alter table "results"."character" add column "episode" _text;
