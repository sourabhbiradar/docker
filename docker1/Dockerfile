FROM golang:1.22-alpine3.18
WORKDIR /basicDocker
COPY . /basicDocker
RUN go build -o gindocker /basicDocker
EXPOSE 7011
CMD [ "./gindocker" ]
 
