alter table "results"."character" add column "created_at" timestamptz
 null default now();
