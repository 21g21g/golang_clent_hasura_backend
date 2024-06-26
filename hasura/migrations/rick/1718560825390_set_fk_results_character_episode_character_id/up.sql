alter table "results"."character_episode"
  add constraint "character_episode_character_id_fkey"
  foreign key ("character_id")
  references "results"."character"
  ("id") on update restrict on delete restrict;
