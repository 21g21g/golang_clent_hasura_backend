alter table "results"."character" drop constraint "character_id_fkey",
  add constraint "character_id_fkey"
  foreign key ("id")
  references "results"."location"
  ("id") on update restrict on delete restrict;
