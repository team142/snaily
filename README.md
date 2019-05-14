# snaily



# Note #
The Go program proxies the Angular app.

## Installing

1. Start the database
```bash
    ./start-pg.sh
```
2. Run setup.sql against postgres::snaily@localhost:5000/madast 

## Running the application

1. Start the angular application:
```bash
    cd snaily-web
    npm i
    ng serve
```

2. Start the go app
```bash
    go run application.go
```


