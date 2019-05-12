xhost local:root
docker run -e DISPLAY=$DISPLAY -v $(pwd):/main -v /tmp/.X11-unix/:/tmp/.X11-unix -u 0 test

