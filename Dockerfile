FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GOPATH /app  
RUN apt-get update && apt-get install libc6-dev libglu1-mesa-dev libgles2-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libasound2-dev git -y
RUN go get github.com/hajimehoshi/ebiten/...
CMD cd /app/src/main && go build -o main . && ./main
