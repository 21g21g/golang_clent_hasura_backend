alter table "results"."episode" add column "created_at" timestamptz
 null default now();
