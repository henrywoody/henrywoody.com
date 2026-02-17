BEGIN;

ALTER TABLE asteroids_score RENAME TO asteroids_score__old;

CREATE TABLE asteroids_score (
	"id" serial PRIMARY KEY,
	"name" varchar(3) NOT NULL CHECK(upper(name) = name),
	"score" integer NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT NOW()
);

INSERT INTO asteroids_score ("name", "score", "created_at")
SELECT "name", "score", "created_date"::timestamptz FROM asteroids_score__old;

DROP TABLE asteroids_score__old;

COMMIT;
