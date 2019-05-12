# [Cyclic cellular automaton!](https://en.wikipedia.org/wiki/Cyclic_cellular_automaton)
![](https://i.imgur.com/uzaJ57u.png)

## Docker setup:
### Installera en X11 server
Eftersom det är en grafisk applikation så krävs det att det finns en X11 server installerad på host-maskinen.
Tyvärr så gör detta att applikationen inte är helt platformsoberoende trots att den körs via docker, och den har bara testats på en linuxmaskin då ingen Mac funnits tillgänglig. 
Alla 3:e parts bibliotek som används finns dock tillgängliga för alla platformar så det enklast är nog att bygga och köra lokalt.

#### Mac: 
1. installera https://www.xquartz.org/
2. slå på [‘Allow connections from network clients’](https://blogs.oracle.com/oraclewebcentersuite/running-gui-applications-on-native-docker-containers-for-mac)

#### Linux:
En X11 server borde finnas installerat på alla vanligare linux installationer, men annars så kan man installera med

`sudo apt-get install xorg`

### Bygg docker image
Kör `docker build -t test .` i rotkatalogen

### Starta programmet
Starta genoma att köra `./run_docker.sh` som super user (Krävs för X11 sockets). Detta slussar vidare bild-output till X servern på det lokala systemet.

## Bygg Lokalt:
1. Installera [Ebiten med dependencies](https://ebiten.org/install.html) för ditt operativsystem
2. starta applikationen genom att köra `go run /src/main/main.go` från rotkatalogen

## Inställningar:
För att få upp all tillgängliga inställningar kan man köra `go run /src/main/main.go -h`
```
  -colors int
        total number of colors used in the simulation (default 16)
  -height int
        screen height (default 480)
  -speed int
        number of updates per second to the simulation (cap at 60) (default 10)
  -threshold int
        required number of preceeding neighbours required before transforming (default 1)
  -width int
        screen width (default 640)
```
