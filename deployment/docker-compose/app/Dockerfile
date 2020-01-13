# ARG needs to be defined for both FROM instructions,
# see https://github.com/moby/moby/issues/34129
ARG FDB_VERSION
FROM foundationdb/foundationdb:${FDB_VERSION} as fdb
FROM golang:1.13.4-stretch
ARG FDB_VERSION

WORKDIR /tmp

RUN apt update
# dnsutils is needed to have dig installed to create cluster file
RUN apt install -y dnsutils 

RUN wget "https://www.foundationdb.org/downloads/${FDB_VERSION}/ubuntu/installers/foundationdb-clients_${FDB_VERSION}-1_amd64.deb"
RUN dpkg -i foundationdb-clients_${FDB_VERSION}-1_amd64.deb

COPY --from=fdb /var/fdb/scripts/create_cluster_file.bash /

WORKDIR /go/src/app
COPY . .

COPY start.bash /start.bash

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/start.bash"]