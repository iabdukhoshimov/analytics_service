-- Create movie view table
CREATE TABLE "movie_view" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "ip_address" varchar NOT NULL,
  "user_id" varchar,
  "movie_slug" varchar NOT NULL,
  "owner_id" varchar NOT NULL,
  "season_number" int,
  "episode_number" int,
  "viewed_at" timestamp NOT NULL,
  "created_at" timestamp NOT NULL,
  "active" boolean NOT NULL,
  "source_place" varchar NOT NULL,
  "audio_lang" varchar,
  "subtitle_lang" varchar,
  "country" varchar,
  "region" varchar,
  "platform" varchar,
  "city" varchar,
  "lon" float,
  "lat" float,
  "category" varchar NOT NULL
);

-- create movie durations
CREATE TABLE "movie_durations" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "movie_slug" varchar NOT NULL,
  "owner_id" varchar NOT NULL,
  "season_number" int,
  "episode_number" int,
  "start_at" int,
  "end_at" int,
  "audio_lang" varchar,
  "subtitle_lang" varchar,
  "ip_address" varchar,
  "user_id" varchar,
  "platform" varchar,
  "created_at" timestamp NOT NULL,
  "category" varchar NOT NULL
);

CREATE TABLE "movie_impressions" (
  "id" varchar UNIQUE PRIMARY KEY NOT NULL,
  "movie_slug" varchar NOT NULL,
  "owner_id" varchar NOT NULL,
  "user_id" varchar,
  "ip_address" varchar,
  "source_place" varchar NOT NULL,
  "country" varchar,
  "region" varchar,
  "platform" varchar,
  "city" varchar,
  "lon" float,
  "lat" float,
  "created_at" timestamp NOT NULL,
  "category" varchar NOT NULL
);
