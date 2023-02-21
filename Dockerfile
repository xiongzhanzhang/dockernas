FROM golang AS gobuilder
ENV GOPROXY https://goproxy.cn
ENV CGO_ENABLED=1
RUN mkdir /app
WORKDIR /app
COPY internal ./internal
COPY dockernas.go ./
COPY go.sum ./
COPY go.mod ./
RUN go build -ldflags '--extldflags "-static -fpic"' ./dockernas.go


FROM alpine

WORKDIR /home/dockernas
COPY apps ./apps
COPY --from=gobuilder /app/dockernas ./
RUN mkdir frontend
COPY frontend/dist ./frontend/dist

ENV DOCKERNAS_RUN_IN_CONTAINER true
CMD ["/home/dockernas/dockernas"]