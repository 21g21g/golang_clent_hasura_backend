alter table "results"."character"
  add constraint "character_location_id_fkey"
  foreign key ("location_id")
  references "results"."location"
  ("id") on update restrict on delete restrict;
