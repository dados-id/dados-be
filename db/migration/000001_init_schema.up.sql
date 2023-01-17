-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-01-17T15:02:43.271Z

CREATE TYPE "StatusRequest" AS ENUM (
  'pending',
  'verified',
  'rejected'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "school" varchar,
  "expected_year_of_graduation" smallint,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_save_professors" (
  "professor_id" bigint,
  "user_id" bigint,
  PRIMARY KEY ("professor_id", "user_id")
);

CREATE TABLE "professors" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "rating" smallint NOT NULL,
  "total_review" int NOT NULL,
  "would_take_again" smallint NOT NULL,
  "level_of_difficulty" decimal(2,1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "status" StatusRequest NOT NULL DEFAULT 'pending',
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "faculty_id" bigint NOT NULL,
  "school_id" bigint NOT NULL
);

CREATE TABLE "professor_course_associations" (
  "professor_id" bigint,
  "course_code" varchar,
  PRIMARY KEY ("professor_id", "course_code")
);

CREATE TABLE "professor_ratings" (
  "id" bigserial PRIMARY KEY,
  "quality" decimal(2,1) NOT NULL,
  "difficult" decimal(2,1) NOT NULL,
  "would_take_again" smallint NOT NULL,
  "taken_for_credit" boolean,
  "use_textbooks" boolean,
  "attendance_mandatory" smallint NOT NULL,
  "grade" varchar,
  "tags" varchar[],
  "review" varchar NOT NULL,
  "up_vote" int NOT NULL DEFAULT 0,
  "down_vote" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "edited_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "professor_id" bigint NOT NULL,
  "course_code" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "verified" boolean DEFAULT false
);

CREATE TABLE "professor_rating_tags" (
  "tag_id" bigint,
  "professor_id" bigint,
  PRIMARY KEY ("tag_id", "professor_id")
);

CREATE TABLE "tags" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "courses" (
  "code" varchar PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "faculties" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "school_faculty_associations" (
  "faculty_id" bigint,
  "school_id" bigint,
  PRIMARY KEY ("faculty_id", "school_id")
);

CREATE TABLE "schools" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "nick_name" varchar[] NOT NULL,
  "country" varchar NOT NULL,
  "province" varchar NOT NULL,
  "website" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "status" StatusRequest NOT NULL DEFAULT 'pending',
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "school_ratings" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "school_id" bigint NOT NULL,
  "reputation" smallint NOT NULL,
  "location" smallint NOT NULL,
  "opportunities" smallint NOT NULL,
  "facilities" smallint NOT NULL,
  "internet" smallint NOT NULL,
  "food" smallint NOT NULL,
  "clubs" smallint NOT NULL,
  "social" smallint NOT NULL,
  "happiness" smallint NOT NULL,
  "safety" smallint NOT NULL,
  "review" varchar NOT NULL,
  "up_vote" int NOT NULL DEFAULT 0,
  "down_vote" int NOT NULL DEFAULT 0,
  "overall_rating" decimal(2,1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "edited_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "verified" boolean DEFAULT false
);

CREATE TABLE "report_forms" (
  "id" bigserial PRIMARY KEY,
  "comment" varchar NOT NULL,
  "status" StatusRequest NOT NULL,
  "request_date" timestamptz NOT NULL DEFAULT (now()),
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "professor_rating_id" bigserial NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "correction_forms" (
  "id" bigserial PRIMARY KEY,
  "problem" varchar,
  "correct_info" varchar,
  "email" varchar,
  "status" StatusRequest NOT NULL DEFAULT 'pending',
  "request_date" timestamptz NOT NULL DEFAULT (now()),
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "user_id" bigint NOT NULL
);

CREATE INDEX ON "professors" ("first_name", "last_name");

CREATE INDEX ON "professors" ("school_id");

CREATE INDEX ON "professors" ("faculty_id");

CREATE INDEX ON "professors" ("faculty_id", "school_id");

CREATE INDEX ON "professor_ratings" ("course_code");

CREATE INDEX ON "professor_ratings" ("professor_id");

CREATE INDEX ON "schools" ("name", "nick_name");

CREATE INDEX ON "school_ratings" ("user_id");

CREATE INDEX ON "school_ratings" ("school_id");

COMMENT ON TABLE "professors" IS '
      List of derived attribute:
      1. top tags
      2. 5 distribusi nilai (from quality)
    ';

COMMENT ON TABLE "professor_ratings" IS '
      attendance_mandatory:
      1. true
      2. false
      3. unknown

      would_take_again:
      0. false
      1. true
    ';

COMMENT ON TABLE "schools" IS '
      List of derived attribute:
      1. 10 avg field school rating
      2. avg Overall Quality
    ';

ALTER TABLE "user_save_professors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_save_professors" ADD FOREIGN KEY ("professor_id") REFERENCES "professors" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_course_associations" ADD FOREIGN KEY ("professor_id") REFERENCES "professors" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_course_associations" ADD FOREIGN KEY ("course_code") REFERENCES "courses" ("code") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_rating_tags" ADD FOREIGN KEY ("professor_id") REFERENCES "professor_ratings" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_rating_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "school_faculty_associations" ADD FOREIGN KEY ("faculty_id") REFERENCES "faculties" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "school_faculty_associations" ADD FOREIGN KEY ("school_id") REFERENCES "schools" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "school_ratings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "correction_forms" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "report_forms" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_ratings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professors" ADD FOREIGN KEY ("faculty_id") REFERENCES "faculties" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "report_forms" ADD FOREIGN KEY ("professor_rating_id") REFERENCES "professor_ratings" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_ratings" ADD FOREIGN KEY ("course_code") REFERENCES "courses" ("code") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professors" ADD FOREIGN KEY ("school_id") REFERENCES "schools" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_ratings" ADD FOREIGN KEY ("professor_id") REFERENCES "professors" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "school_ratings" ADD FOREIGN KEY ("school_id") REFERENCES "schools" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
