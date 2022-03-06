# UkeAPI

Print basic Ukulele chords as text, JSON or HTML.

## CLI usage

```
go run main.go -chord C
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

```

```
go run main.go -chord F -key
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
1 = index finger, 2 = middle finger, 3 = ring finger, 4 = pinky

```

## API usage

```
# start local server on port 8080 (default)
go run main.go -serve
```

```
# option: if a different port is needed, start as follows (*adapt examples below accordingly*)
PORT=8081 go run main.go -serve
```

### Get C chord

- Get HTML

Browse to e.g. http://localhost:8080/api/v1/C or http://localhost:8080/api/v1/F/key

or

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/api/v1/C"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/api/v1/F/key"
```

- Get JSON

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/api/v1/C"
```

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/api/v1/F/key"
```


- Get text

```
curl localhost:8080/api/v1/C
```

```
curl localhost:8080/api/v1/F/key
```

## To do

- Host
- Support sequence of chords, e.g. AmGC