source ./.env
source setenv.sh
#docker run -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable" up
