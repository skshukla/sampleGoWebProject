##### Run the below commands to run a fresh PostGres DB in Docker

```
docker kill pg_sachin

docker  run --rm --name sachin_pg_c -e POSTGRES_USER=sachin -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=mydb -p 5432:5432 -d postgres:9.6.14

docker exec -it sachin_pg_c psql -h localhost --user sachin --db mydb

create table if not exists employees(
 id bigserial primary key,
 name varchar(100)
)
;
```