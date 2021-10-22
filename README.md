Goddb Sample App
---------------------

* Introduction
* Requirements
* Installation
* Configuration
* Customization


Introduction
------------
this project is a sample app to demonstrate key:value db which is called goddb


Requirements
------------

This module requires no modules outside of Golang's standard library.


Installation
------------

* docker

``` 
docker run -d \
  --restart=always \
  -p 8099:8099 \
  --name goddb-api \
  ghcr.io/mevlanaayas/goddb/goddb
```
* go cli

``` 
go build .
./goddb
```


Configuration
-------------

``` 
{
  "server": {
    "port": 8099
  },
  "app": {
    "path": "tmp",
    "syncInMin": 1
  }
}
```
* port: port that application runs on. do not forget to edit Dockerfile
* timer: time in min that after how much time scheduler will save inmemory data to filesystem 
* path: path for persisting/reading in-memory data to/from filesystem


Customization
-------------

* readwriter

``` 
write custom struct that implements these two methods

Read() (error, []byte)
Write([]byte) error
```
* getputflusher

``` 
write custom struct that implements these five methods

Get(key string) (error, string)
GetAll() (error, map[string]string)
Put(key, value string) error
PutAll(values map[string]string) error
Flush() error
```
```
inject newly created structs to related fields under app.go
```

Maintainers
-----------

* Mevlana - https://github.com/mevlanaayas
