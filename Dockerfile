FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GOPATH /app
RUN apt-get update && apt-get install libc6-dev libglu1-mesa-dev libgles2-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libasound2-dev git -y
RUN go get -v github.com/hajimehoshi/ebiten/...
ENTRYPOINT ["go", "run", "/app/src/main/main.go"]

