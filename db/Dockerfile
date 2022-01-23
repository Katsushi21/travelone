FROM mysql:8.0.27

RUN touch /var/log/mysql.log
COPY ./my.cnf /etc/mysql/conf.d/
COPY ./init/1_createdb.sql /docker-entrypoint-initdb.d

CMD ["mysqld --character-set-server=utf8 --collation-server=utf8_unicode_ci"]