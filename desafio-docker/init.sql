CREATE USER docker_usr WITH ENCRYPTED PASSWORD 'docker_pwd';
CREATE DATABASE curso_docker;
GRANT ALL PRIVILEGES ON DATABASE curso_docker TO docker_usr;
