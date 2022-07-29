CREATE TABLE IF NOT EXISTS "playlist" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" VARCHAR NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "song" (
    "title" VARCHAR NOT NULL PRIMARY KEY,
    "genre" VARCHAR NOT NULL,
    "artist" VARCHAR NOT NULL,
    "duration" int NOT NULL
);

CREATE TABLE IF NOT EXISTS "queue" (
    "id_playlist" INT NOT NULL,
    "title_song" VARCHAR NOT NULL,
    "position" int NOT NULL,
    CONSTRAINT "fk_id" FOREIGN KEY ("id_playlist") REFERENCES "playlist" ("id"),
    CONSTRAINT "fk_title" FOREIGN KEY ("title_song") REFERENCES "song" ("title")
);