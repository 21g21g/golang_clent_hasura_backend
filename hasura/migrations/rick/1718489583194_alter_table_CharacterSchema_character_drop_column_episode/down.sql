alter table "CharacterSchema"."character" alter column "episode" drop not null;
alter table "CharacterSchema"."character" add column "episode" text;
