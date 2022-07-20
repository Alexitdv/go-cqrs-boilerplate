# Readme 

## Development environment
You have to install Docker. Pls use OSx or Linux to working with it. 
### Databases 
`docker run --name postgres -p 5455:5432 -e POSTGRES_USER=pgRoot -e POSTGRES_PASSWORD=pgpass -e PGDATA=/var/lib/postgresql/data/pgdata -v ~/var/lib/pglatest:/var/lib/postgresql/data -d postgres`

## Migrations
go get -u github.com/pressly/goose/cmd/goose or brew install goose
goose -dir ./migrations/sql create some_name sql
goose -dir ./migrations/sql postgres "user=boilerplate password=boilerplatepwd dbname=cqrs-boilerplate port=5455 sslmode=disable" up
goose -dir ./migrations/sql postgres "user=boilerplate password=boilerplatepwd dbname=cqrs-boilerplate port=5455 sslmode=disable" down

## Minikube
https://github.com/kubernetes/minikube/releases/
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v1.18.1/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/

### Notes
https://cloud.google.com/endpoints/docs/grpc/transcoding
https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api