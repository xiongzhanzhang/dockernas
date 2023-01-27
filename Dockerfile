FROM alpine:3.17

WORKDIR /home/dockernas
COPY apps ./apps
COPY dockernas ./
RUN mkdir frontend
COPY frontend/dist ./frontend/dist

ENV DOCKERNAS_RUN_IN_CONTAINER true
CMD ["/home/dockernas/dockernas"]