alter table "results"."character"
  add constraint "character_origin_id_fkey"
  foreign key ("origin_id")
  references "results"."location"
  ("id") on update restrict on delete restrict;
