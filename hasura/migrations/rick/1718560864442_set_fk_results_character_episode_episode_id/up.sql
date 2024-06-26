alter table "results"."character_episode"
  add constraint "character_episode_episode_id_fkey"
  foreign key ("episode_id")
  references "results"."episode"
  ("id") on update restrict on delete restrict;
