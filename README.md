# ratonera - make transportation great again
============================================

In this project we want to create a transportation management system

![alt tag](https://github.com/brunomorishita/ratonera/blob/master/res/raton.jpg)

### Installation
    go get github.com/gorilla/websocket
    go get gopkg.in/mgo.v2
    go get github.com/brunomorishita/ratonera
    cd $GOPATH/github.com/gorilla/brunomorishita/ratonera/server
    go build server.go

### What we have until now
  * A server written in go that stores in a mongo db gps, accelerometer, timestamp data from an android smartphone
  * An android app that send this kind of information

### What we want to achieve
  * A server that can extract useful information for helping truck drivers
  * A server that can extract useful information for helping shipping companies
  * Be a millionaire
