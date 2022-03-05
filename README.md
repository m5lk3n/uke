# UkeAPI

**This fun project is work in progress!**

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

### Get C chord

- Get HTML

Browse to e.g. http://localhost:8080/C or http://localhost:8080/F/key

or

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/C"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/F/key"
```

- Get JSON

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/C"
```

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/F/key"
```


- Get text

```
curl localhost:8080/C
```

```
curl localhost:8080/F/key
```

## To do

- Clean up uke interface
- Document/Swagger
- Add Makefile?
- Host?