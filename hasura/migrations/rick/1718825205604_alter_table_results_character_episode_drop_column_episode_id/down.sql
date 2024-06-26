alter table "results"."character_episode" alter column "episode_id" set default nextval('results.character_episode_episode_id_seq'::regclass);
alter table "results"."character_episode" add constraint "character_episode_episode_id_key" unique (episode_id);
alter table "results"."character_episode"
  add constraint "character_episode_episode_id_fkey"
  foreign key (episode_id)
  references "results"."episode"
  (id) on update restrict on delete restrict;
alter table "results"."character_episode" alter column "episode_id" drop not null;
alter table "results"."character_episode" add column "episode_id" int4;
