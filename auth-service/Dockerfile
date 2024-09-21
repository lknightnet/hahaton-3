FROM ubuntu

WORKDIR /auth-service

COPY config/config.yaml config/
COPY auth-service /auth-service

EXPOSE 8004

CMD ["./auth-service"]