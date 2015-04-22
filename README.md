## Timeout [![Build Status](https://travis-ci.org/yaman/timeout.svg?branch=master)](https://travis-ci.org/yaman/timeout) [![Coverage Status](https://coveralls.io/repos/yaman/timeout/badge.svg?branch=master)](https://coveralls.io/r/yaman/timeout?branch=master) [![Heroku](https://heroku-badge.herokuapp.com/?app=gotimeout)](https://gotimeout.herokuapp.com/anyresource?sleep=1)

timeout is written to satisfy `tcp` and `http` based timeout testing needs. It works as a stub server in replacement of your application's external tcp/http dependency.

Briefly speaking, you can run a tcp/http stub server in seconds.

-------------------------
#### HTTP

- Respond with status `http:yourhost/yourresource?status=x` where x is an integer number between 200 and 599: 
    * `https://gotimeout.herokuapp.com/yourpreciousresource?status=201` 
    * `https://gotimeout.herokuapp.com/yourpreciousresource?status=302` 
    * `https://gotimeout.herokuapp.com/yourpreciousresource?status=401` 

- Sleep before responding `http:yourhost/yourresource?sleep=x` x is in milliseconds:
    * `https://gotimeout.herokuapp.com/yourpreciousresource?sleep=1000` --> wait for 1 second before responding
    * `https://gotimeout.herokuapp.com/yourpreciousresource?sleep=5000` --> wait for 5 second before responding

-------------------------
#### TCP

- Listen but do not respond, from port `:5501`
- Listen and respond with empty string, from port `:5502`
- Listen and respond with empty string after client sends data, `:5503`
- Listen and respond with malformed string, from port `:5504`
- Listen and respond with malformed string after client sends data, from port `:5505`
- Listen and respond with empty string every 5 seconds, from port `:5506`
- Listen and respond with empty string every 30 seconds, from port `:5507`

-------------------------
#### Install
`go get github.com/yaman/timeout`

-------------------------
#### Run
`$GOPATH/bin/timeout -proto=http -port=8080`

`$GOPATH/bin/timeout -proto=tcp`

-------------------------
#### Directly use it from Heroku
Only Http timeout is available through heroku for now. Just change your external url to ;

`https://gotimeout.herokuapp.com/yourpreciousresource?sleep=1`

-------------------------
#### Run with Docker prebuild image
You just need to run following command provided that Docker is up and running on your local machine;

###### for http and exposing port 5000

`sudo docker run -p 5000:5000 -e PROTO=http -e PORT=5000 yaman/timeout`

###### for tcp

`sudo docker run -p 5500-5600:5500-5600 -e PROTO=tcp yaman/timeout`

-------------------------
