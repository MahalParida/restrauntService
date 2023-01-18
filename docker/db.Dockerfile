FROM mysql:latest

COPY ./db/migration/*.sql /docker-entrypoint-initdb.d/