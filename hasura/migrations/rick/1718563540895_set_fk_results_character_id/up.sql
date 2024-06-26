alter table "results"."character"
  add constraint "character_id_fkey"
  foreign key ("id")
  references "results"."location"
  ("id") on update restrict on delete restrict;
