ARG APP_NAME=transactions-api

# build stage
FROM golang:alpine3.12 as builder
ARG APP_NAME
RUN apk add --no-cache git
ADD . /src
WORKDIR /src
ENV GOOS=linux
ENV GARCH=amd64
ENV CGO_ENABLED=0
RUN go build -v -a -installsuffix cgo -o ${APP_NAME} cmd/*.go

# final stage
FROM alpine
ARG APP_NAME
COPY --from=builder /src/${APP_NAME} .
COPY --from=builder /src/application.yml .
ENV TZ America/Sao_Paulo
CMD /${APP_NAME}
