## Timeout
[![Build Status](https://travis-ci.org/yaman/timeout.svg?branch=master)](https://travis-ci.org/yaman/timeout) [![Coverage Status](https://coveralls.io/repos/yaman/timeout/badge.svg?branch=master)](https://coveralls.io/r/yaman/timeout?branch=master) [![Heroku](https://heroku-badge.herokuapp.com/?app=gotimeout)](https://gotimeout.herokuapp.com/anyresource?sleep=1)
===========================================

timeout is written to satisfy `tcp` and `http` based timeout testing needs. It works as a stub server in replacement of your application's external tcp/http dependency.

Briefly speaking, you can run a tcp/http stub server in seconds.


#### TCP


-------------------------
#### HTTP


-------------------------
#### Install

go get github.com/yaman/timeout

-------------------------
#### Run

$GOPATH/bin/timeout -proto=http -port=8080

$GOPATH/bin/timeout -proto=tcp

-------------------------
#### Directly use it from Heroku
Only Http timeout is available through heroku for now. Just change your external url to ;

`https://gotimeout.herokuapp.com/yourpreciousresource?sleep=1`

-------------------------


