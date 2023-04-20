FROM golang:1.20.3-bullseye@sha256:595c9af0430dd84bad33020e7e9e328af4bd1a1aabd46a03b5bf6f252cdbecf3 AS dev

RUN go install github.com/cosmtrek/air@v1.43.0

WORKDIR /things/services/api
COPY ./services/api/go.mod ./services/api/go.sum /things/services/api/
RUN go mod download

COPY ./services/api/ /things/services/api/


ENTRYPOINT [ "air" ]
CMD [ "" ]

FROM golang:1.20.3-alpine3.17@sha256:08e9c086194875334d606765bd60aa064abd3c215abfbcf5737619110d48d114 AS runtime
