ARG FDB_VERSION=6.3.23

FROM foundationdb/foundationdb:${FDB_VERSION} as fdb
FROM golang:1.13.6-stretch
ARG FDB_VERSION

RUN apt update
# dnsutils is needed to have dig installed to create cluster file
RUN apt install -y dnsutils 

WORKDIR /tmp

RUN mkdir -p /var/lib/foundationdb
RUN wget "https://github.com/apple/foundationdb/releases/download/${FDB_VERSION}/foundationdb-clients_${FDB_VERSION}-1_amd64.deb"
RUN dpkg -i foundationdb-clients_${FDB_VERSION}-1_amd64.deb

COPY --from=fdb /var/fdb/scripts/create_cluster_file.bash /

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/go/bin/fdb-prometheus-exporter"]
