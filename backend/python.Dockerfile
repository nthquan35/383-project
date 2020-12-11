#FROM python:3.8
#FROM golang:1.15

FROM golang:latest

RUN go get github.com/mattes/migrate
RUN apt-get update && apt-get install -y python3-pip
#RUN apt-get install python-ctypeslib

#RUN pip install pika==1.1.0
#RUN pip install pyzmq==19.0.1
RUN pip3 install flask
RUN pip3 install flask_cors
RUN pip3 install numpy
#RUN pip3 install ctypes

WORKDIR /app
COPY . .

# for ZeroMQ server
#EXPOSE 5555

#CMD ["go", "build", "-o", "main.so", "-buildmode=c-shared", "main.go", "\&\&", "python3", "server.py"]

CMD go build -o main.so -buildmode=c-shared main.go && python3 server.py