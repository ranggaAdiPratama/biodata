CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY, "username" varchar UNIQUE NOT NULL, "name" varchar NOT NULL, "email" varchar UNIQUE NOT NULL, "password" varchar NOT NULL, "profile_picture" varchar NULL, "created_at" timestamp NOT NULL DEFAULT(now()), "updated_at" timestamp NULL
);

CREATE TABLE "hobbies" (
    "id" bigserial PRIMARY KEY, "user_id" bigserial NOT NULL, "name" varchar NOT NULL, "created_at" timestamp NOT NULL DEFAULT(now()), "updated_at" timestamp
);

CREATE INDEX ON "hobbies" ("user_id");

ALTER TABLE "hobbies"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

INSERT INTO
    "users" (
        "id", "username", "name", "email", "password", "profile_picture", "created_at", "updated_at"
    )
VALUES (
        1, 'rangga', 'Rangga Adi Pratama', 'masterrangga@gmail.com', '$2a$10$nYJnqTID3BLwcEZICroXnOj37Lt1gbEVppPlvnNdSsc2CFbNpkbJ2', NULL, '2024-01-21 13:46:27.463394', NULL
    ),
    (
        2, 'mitsuha', 'Mitsuha Miyamizu', 'mitsuha@gmail.com', '$2a$10$ibz7oLoKG4BpHEHOn8JAQeuk/GDlgTikYxIh1YpvvOXg7eQABvlWC', NULL, '2024-01-25 06:39:24.665504', NULL
    )