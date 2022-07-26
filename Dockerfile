FROM alpine


# http port
EXPOSE 80
# grpc port
EXPOSE 50052
# websocket port
EXPOSE 8080
# fileserver port
EXPOSE 8939
# fileserver port
EXPOSE 8501

COPY /toy-gin /app

CMD ["./app"]