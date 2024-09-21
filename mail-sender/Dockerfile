FROM ubuntu

WORKDIR /mail-sender
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates
COPY config/config.yaml config/
COPY mail-sender /mail-sender

EXPOSE 8003

CMD ["./mail-sender"]