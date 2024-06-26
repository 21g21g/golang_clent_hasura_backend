alter table "results"."character_episode" alter column "character_id" set default nextval('results.character_episode_character_id_seq'::regclass);
alter table "results"."character_episode" add constraint "character_episode_character_id_key" unique (character_id);
alter table "results"."character_episode"
  add constraint "character_episode_character_id_fkey"
  foreign key (character_id)
  references "results"."character"
  (id) on update restrict on delete restrict;
alter table "results"."character_episode" alter column "character_id" drop not null;
alter table "results"."character_episode" add column "character_id" int4;
