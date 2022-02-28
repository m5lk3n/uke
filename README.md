# UkeAPI

**this fun project is work in progress!**

Print basic Ukulele chords like so e.g.:

```
C
+==+==+==+
|  |  |  |
+--+--+--+
|  |  |  |
+--+--+--+
|  |  |  3
+--+--+--+
|  |  |  |
+--+--+--+
1 = index finger, 2 = middle finger, 3 = ring finger, 4 = pinky
```

## CLI usage

```
go run main.go -chord F
F
+==+==+==+
|  |  1  |
+--+--+--+
2  |  |  |
+--+--+--+
|  |  |  |
+--+--+--+
|  |  |  |
+--+--+--+

```

## API usage

```
go run main.go -serve
```

### get C chord

- get HTML

Browse to http://localhost:8080/C or

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/C"
```

- get JSON

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/C"
```

- get text

```
curl localhost:8080/C
```

## to do

- clean up uke interface
- support key in interface
- document/Swagger
- add LICENSE
- add Makefile?
- host?