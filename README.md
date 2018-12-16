# Optimus

This is the middleware HTTP server to communicate between the IOT device and Controller.

* Install

```go get bitbucket.org/nihanthd/optimus/...```

* Development

    1. Clone ```git clone https://bitbucket.org/nihanthd/optimus.git```
    2. To get the dependencies ```dep ensure```
    3. To build ```cd cmd/optimus && go build```
    4. To run ```./optimus -c ../../config.yaml```
```go get github.com/nihanthd/optimus/...```

* Development

    1. Clone ```git clone https://github.com/nihanthd/optimus.git```
    2. To get the dependencies ```dep ensure```
    3. To build ```cd cmd/optimus && go build```
    4. To run ```./optimus -c ../../config.yaml```
    
## Acknowledgement

1. Uber fx (https://github.com/uber-go/fx)
2. Echo (https://github.com/labstack/echo)
3. go-rpio (https://github.com/stianeikeland/go-rpio)