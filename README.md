# [Cyclic cellular automaton!](https://en.wikipedia.org/wiki/Cyclic_cellular_automaton)

## Docker setup:
### Installera en X11 server
(Eftersom det är en grafisk applikation så krävs det att det finns en X11 server installerad på host-maskinen)

#### Mac: 
1. installera https://www.xquartz.org/
2. slå på [‘Allow connections from network clients’](https://blogs.oracle.com/oraclewebcentersuite/running-gui-applications-on-native-docker-containers-for-mac)

#### Linux:
En X11 server borde finnas installerat på alla vanligare linux installationer, men annars så kan man installera med

`sudo apt-get install xorg`

### Bygg docker image
Kör `docker build -t test .` i rotkatalogen

### Starta programmet
Starta genoma att köra `./run_docker.sh`. Detta slussar vidare bild-output till X servern på det lokala systemet.

## Bygg Lokalt:
1. Installera [Ebiten med dependencies](https://ebiten.org/install.html) för ditt operativsystem
2. starta applikationen genom att köra `go run /src/main/main.go` från rotkatalogen
