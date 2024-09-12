FROM golang:1.23.0

WORKDIR /app

RUN apt update -y
RUN apt upgrade -y

RUN apt install make -y

EXPOSE 8080

RUN echo "alias project='source /app/go.sh'" >> ~/.bashrc
RUN source ~/.bashrc

CMD ["sleep", "infinity"]

