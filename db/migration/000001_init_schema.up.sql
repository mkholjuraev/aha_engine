CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar,
  "created_at" timestamp DEFAULT (now()),
  "country_code" int,
  "password" varchar,
  "login" varchar,
  "telephone" varchar,
  "socialLinks" varchar,
  "notifications" bigserial
);

CREATE TABLE "writer" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial,
  "biograph" varchar,
  "distinct_likes" int,
  "distinct_views" int
);

CREATE TABLE "followers" (
  "id" bigserial,
  "user_id" bigserial,
  "follower" bigserial
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "created_at" timestamp DEFAULT (now()),
  "description" varchar,
  "writer" bigserial,
  "content" varchar,
  "views" int,
  "likes" int,
  "shares" int
);

CREATE TABLE "notifications" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "message" varchar,
  "posts" bigserial
);

CREATE TABLE "chat" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "message" varchar,
  "sender" bigserial,
  "receiver" bigserial
);

ALTER TABLE "users" ADD FOREIGN KEY ("notifications") REFERENCES "notifications" ("id");

ALTER TABLE "writer" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("follower") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("writer") REFERENCES "writer" ("id");

ALTER TABLE "notifications" ADD FOREIGN KEY ("posts") REFERENCES "posts" ("id");

ALTER TABLE "chat" ADD FOREIGN KEY ("sender") REFERENCES "users" ("id");

ALTER TABLE "chat" ADD FOREIGN KEY ("receiver") REFERENCES "users" ("id");