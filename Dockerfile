FROM postgres:12-alpine

COPY ./secrets ./secrets

EXPOSE 5432