FROM mariadb:latest
RUN touch /var/log/mysql/mysqld.log
ADD ./initdb.d/ /docker-entrypoint-initdb.d
ADD  ./my.cnf /etc/mysql/my.cnf
RUN chmod 644 /etc/mysql/my.cnf
