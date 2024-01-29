#/bin/sh
docker run --name devops_pro_postgresql -p 5432:5432 -e POSTGRES_USER=docker_usr -e POSTGRES_PASSWORD=docker_pwd -e POSTGRES_DB=curso_docker -v $(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql -d postgres
