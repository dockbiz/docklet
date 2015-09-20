FROM google/golang:latest

RUN cd src/ && go get && go build -o ../bin/docklet
EXPOSE 8080