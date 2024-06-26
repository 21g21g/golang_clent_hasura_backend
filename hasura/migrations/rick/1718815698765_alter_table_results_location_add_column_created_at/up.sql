alter table "results"."location" add column "created_at" timestamptz
 null default now();
