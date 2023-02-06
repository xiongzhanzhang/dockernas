FROM golang:alpine AS gobuilder
ENV GOPROXY https://goproxy.cn
RUN mkdir /app
WORKDIR /app
COPY internal ./internal
COPY dockernas.go ./
COPY go.sum ./
COPY go.mod ./
RUN go build ./dockernas.go


FROM alpine:3.17

WORKDIR /home/dockernas
COPY apps ./apps
COPY --from=gobuilder /app/dockernas ./
RUN mkdir frontend
COPY frontend/dist ./frontend/dist

ENV DOCKERNAS_RUN_IN_CONTAINER true
CMD ["/home/dockernas/dockernas"]