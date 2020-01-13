ARG FDB_VERSION=6.2.11
FROM golang:1.13.4-stretch
ARG FDB_VERSION

WORKDIR /tmp

RUN wget "https://www.foundationdb.org/downloads/${FDB_VERSION}/ubuntu/installers/foundationdb-clients_${FDB_VERSION}-1_amd64.deb"
RUN dpkg -i foundationdb-clients_${FDB_VERSION}-1_amd64.deb

WORKDIR /go/src/app
COPY . .


RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/start.bash"]