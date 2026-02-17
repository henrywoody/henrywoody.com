BEGIN;

CREATE TABLE asteroids_scores (
    score INTEGER NOT NULL,
    name VARCHAR(3) NOT NULL CHECK(upper(name) = name),
    created_date DATE NOT NULL DEFAULT CURRENT_DATE
);

COMMIT;
