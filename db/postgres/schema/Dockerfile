FROM postgres
RUN apt-get update && \
    apt-get install -y postgis && \
    apt-get install -y postgresql-11-postgis-2.5
COPY migrations/ /docker-entrypoint-initdb.d
VOLUME [ "/backups" ]
EXPOSE 5432