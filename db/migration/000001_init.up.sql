CREATE TABLE
    "users" (
        "id" bigserial PRIMARY KEY,
        "username" varchar UNIQUE NOT NULL,
        "name" varchar NOT NULL,
        "email" varchar UNIQUE NOT NULL,
        "password" varchar NOT NULL,
        "profile_picture" varchar,
        "created_at" timestamp NOT NULL DEFAULT (now()),
        "updated_at" timestamp
    );

CREATE TABLE
    "hobbies" (
        "id" bigserial PRIMARY KEY,
        "user_id" bigserial NOT NULL,
        "name" varchar NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now()),
        "updated_at" timestamp
    );

CREATE INDEX ON "hobbies" ("user_id");

ALTER TABLE "hobbies"
ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("id");