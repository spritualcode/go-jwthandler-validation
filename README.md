# go-jwthandler-validation example

This is an example to demonstrate go-jwthandler configuration for third party API calls. As mentioned, I use go-jwthandler(github.com/dgrijalva/jwt-go).

In order to run this example...

1) Go to  ..handler/TestingHandler.go

Also make sure, you pass correct private key to validate accessToken as part of request header.

2) Please make sure, you have glide installed

3) Run $ glide update

4) Run $ go build

5) $ ./go-jwthandler-validation

6) open browser localhost:8080/example

I use gorilla Mux for the routing.

If you have any question, feel free to contact me.....
