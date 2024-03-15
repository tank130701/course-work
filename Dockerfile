FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# install go-migrate
RUN GO111MODULE=on go get -v github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# build go app
RUN go mod download
RUN cd cmd
RUN go build -o todo-app github.com/tank130701/course-work/todo-app/back-end/cmd
RUN cd ..

CMD migrate -path /schema/postgres -database postgres://postgres:mypassword@localhost:5436/database up && ./todo-app
