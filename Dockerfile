FROM golang:1.23.1-alpine

WORKDIR /app

RUN apk update && apk upgrade

EXPOSE 8080

RUN echo "alias project='source /app/go.sh'" >> ~/.bashrc
RUN echo "source ~/.bashrc" >> ~/.bash_profile

CMD ["sleep", "infinity"]

