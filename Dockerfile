FROM golang:1.23.0

WORKDIR /app

RUN apt update -y
RUN apt upgrade -y

RUN apt install make -y

CMD ["sleep", "infinity"]

