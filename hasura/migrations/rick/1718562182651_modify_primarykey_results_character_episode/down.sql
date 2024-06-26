alter table "results"."character_episode"
    add constraint "character_episode_pkey"
    primary key ("character_id", "episode_id");
