alter table "results"."character" alter column "origin_id" set default nextval('results.character_origin_id_seq'::regclass);
alter table "results"."character"
  add constraint "character_origin_id_fkey"
  foreign key (origin_id)
  references "results"."location"
  (id) on update restrict on delete restrict;
alter table "results"."character" alter column "origin_id" drop not null;
alter table "results"."character" add column "origin_id" int4;
