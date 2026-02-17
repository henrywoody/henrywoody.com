\set db_name `echo $DB_NAME`
\set db_user `echo $DB_USER`
\set db_pass `echo $DB_PASS`

CREATE DATABASE :db_name;

CREATE ROLE :db_user WITH LOGIN PASSWORD :'db_pass';

GRANT ALL PRIVILEGES ON DATABASE :db_name TO :db_user;