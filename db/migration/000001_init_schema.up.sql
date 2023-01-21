-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-01-17T15:02:43.271Z

CREATE TYPE StatusRequest AS ENUM (
  'pending',
  'verified',
  'rejected'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "school" varchar NOT NULL,
  "expected_year_of_graduation" smallint NOT NULL,
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
  "rating" decimal(2,1) NOT NULL DEFAULT 0 CHECK (rating >= 0),
  "total_review" int NOT NULL DEFAULT 0 CHECK (total_review >= 0),
  "would_take_again" smallint NOT NULL DEFAULT 0 CHECK (would_take_again >= 0),
  "level_of_difficulty" decimal(2,1) NOT NULL DEFAULT 0 CHECK (level_of_difficulty >= 0.0),
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
  "quality" decimal(2,1) NOT NULL CHECK (quality >= 0.0 AND quality <= 5.0),
  "difficult" decimal(2,1) NOT NULL CHECK (difficult >= 0.0 AND difficult <= 5.0),
  "would_take_again" smallint NOT NULL CHECK (would_take_again >= 0 AND would_take_again <= 1),
  "taken_for_credit" smallint NOT NULL CHECK (taken_for_credit >= 0 AND taken_for_credit <= 2),
  "use_textbooks" smallint NOT NULL CHECK (would_take_again >= 0 AND would_take_again <= 2),
  "attendance_mandatory" smallint NOT NULL CHECK (attendance_mandatory >= 0 AND attendance_mandatory <= 2),
  "grade" varchar NOT NULL,
  -- "tags" varchar[] NOT NULL,
  "review" varchar NOT NULL,
  "up_vote" int NOT NULL DEFAULT 0,
  "down_vote" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "edited_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "status" StatusRequest NOT NULL DEFAULT 'pending',
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',

  "professor_id" bigint NOT NULL,
  "course_code" varchar NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "professor_rating_tags" (
  "tag_name" varchar,
  "professor_rating_id" bigint,
  PRIMARY KEY ("tag_name", "professor_rating_id")
);

CREATE TABLE "tags" (
  "name" varchar PRIMARY KEY
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
  "city" varchar NOT NULL,
  "province" varchar NOT NULL,
  "website" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "status" StatusRequest NOT NULL DEFAULT 'pending',
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "school_ratings" (
  "id" bigserial PRIMARY KEY,
  "reputation" smallint NOT NULL CHECK (reputation >= 1 AND reputation <= 5),
  "location" smallint NOT NULL CHECK (location >= 1 AND location <= 5),
  "opportunities" smallint NOT NULL CHECK (opportunities >= 1 AND opportunities <= 5),
  "facilities" smallint NOT NULL CHECK (facilities >= 1 AND facilities <= 5),
  "internet" smallint NOT NULL CHECK (internet >= 1 AND internet <= 5),
  "food" smallint NOT NULL CHECK (food >= 1 AND food <= 5),
  "clubs" smallint NOT NULL CHECK (clubs >= 1 AND clubs <= 5),
  "social" smallint NOT NULL CHECK (social >= 1 AND social <= 5),
  "happiness" smallint NOT NULL CHECK (happiness>= 1 AND happiness <= 5),
  "safety" smallint NOT NULL CHECK (safety >= 1 AND safety <= 5),
  "review" varchar NOT NULL,
  "up_vote" int NOT NULL DEFAULT 0,
  "down_vote" int NOT NULL DEFAULT 0,
  "overall_rating" decimal(2,1) NOT NULL CHECK (overall_rating >= 0.0),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "edited_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "status" StatusRequest NOT NULL DEFAULT 'pending',
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',

  "user_id" bigint NOT NULL,
  "school_id" bigint NOT NULL
);

CREATE TABLE "report_forms" (
  "id" bigserial PRIMARY KEY,
  "comment" varchar NOT NULL,
  "status" StatusRequest NOT NULL DEFAULT 'pending',
  "request_date" timestamptz NOT NULL DEFAULT (now()),
  "verified_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',

  "professor_rating_id" bigserial NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "correction_forms" (
  "id" bigserial PRIMARY KEY,
  "problem" varchar NOT NULL,
  "correct_info" varchar NOT NULL,
  "email" varchar NOT NULL,
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
      1. 5 distribusi nilai (from quality)
    ';

COMMENT ON TABLE "professor_ratings" IS '
      taken_for_credit (int):
      0 -> unknown
      1 -> false
      2 -> true

      use_textbooks (int):
      0 -> unknown
      1 -> false
      2 -> true

      attendance_mandatory (int):
      0 -> unknown
      1 -> false
      2 -> true

      would_take_again (int):
      0 -> false
      1 -> true
    ';

COMMENT ON TABLE "schools" IS '
      List of derived attribute:
      1. 10 avg field school rating
      2. avg overall quality
    ';

ALTER TABLE "professors" ADD UNIQUE ("first_name", "last_name");

ALTER TABLE "user_save_professors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "user_save_professors" ADD FOREIGN KEY ("professor_id") REFERENCES "professors" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_course_associations" ADD FOREIGN KEY ("professor_id") REFERENCES "professors" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_course_associations" ADD FOREIGN KEY ("course_code") REFERENCES "courses" ("code") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_rating_tags" ADD FOREIGN KEY ("professor_rating_id") REFERENCES "professor_ratings" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "professor_rating_tags" ADD FOREIGN KEY ("tag_name") REFERENCES "tags" ("name") ON DELETE CASCADE ON UPDATE CASCADE;

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

-- Trigger School Rating
CREATE OR REPLACE FUNCTION update_overall_rating()
RETURNS TRIGGER AS $$
BEGIN
  NEW.overall_rating := (
    NEW.reputation +
    NEW.location +
    NEW.opportunities +
    NEW.facilities +
    NEW.internet +
    NEW.food +
    NEW.clubs +
    NEW.social +
    NEW.happiness +
    NEW.safety
  ) * 1.0 / 10;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER tr_insert_update_overall_rating
BEFORE INSERT OR UPDATE ON school_ratings
FOR EACH ROW
EXECUTE FUNCTION update_overall_rating();

-- Trigger Professor Field
CREATE OR REPLACE FUNCTION update_professor_field()
RETURNS TRIGGER AS $$
DECLARE
  avg_rating DECIMAL(2, 1);
  avg_level_of_difficulty DECIMAL(2, 1);
  avg_would_take_again INTEGER;
  avg_total_review INTEGER;
BEGIN

  SELECT
    AVG(PR.quality),
    AVG(PR.difficult),
    AVG(PR.would_take_again)*100,
    COUNT(PR.id)
  INTO
    avg_rating,
    avg_level_of_difficulty,
    avg_would_take_again,
    avg_total_review
  FROM professor_ratings PR
    WHERE PR.professor_id = NEW.professor_id
  GROUP BY PR.professor_id;

  UPDATE professors
  SET
    rating = avg_rating,
    total_review = avg_total_review,
    would_take_again = avg_would_take_again,
    level_of_difficulty = avg_level_of_difficulty
  WHERE
    id = NEW.professor_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER tr_insert_update_professor_field
AFTER INSERT OR UPDATE ON professor_ratings
FOR EACH ROW
EXECUTE FUNCTION update_professor_field();

CREATE TRIGGER tr_delete_professor_field
AFTER DELETE ON professor_ratings
FOR EACH ROW
EXECUTE FUNCTION update_professor_field();
