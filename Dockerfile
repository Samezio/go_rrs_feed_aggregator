FROM postgres
ENV POSTGRES_DB=rssagg
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
#RUN psql -U postgres -c "CREATE DATABASE ${POSTGRES_DB}"
#RUN psql -U postgres -c "CREATE ROLE ${POSTGRES_USER} WITH PASSWORD '${POSTGRES_PASSWORD}'"
#RUN psql -U postgres -d ${POSTGRES_DB} -c "GRANT ALL PRIVILEGES ON DATABASE ${POSTGRES_DB} TO ${POSTGRES_USER}"
EXPOSE 5432