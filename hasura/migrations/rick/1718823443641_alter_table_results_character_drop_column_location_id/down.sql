alter table "results"."character" alter column "location_id" set default nextval('results.character_location_id_seq'::regclass);
alter table "results"."character"
  add constraint "character_location_id_fkey"
  foreign key (location_id)
  references "results"."location"
  (id) on update restrict on delete restrict;
alter table "results"."character" alter column "location_id" drop not null;
alter table "results"."character" add column "location_id" int4;
