# UkeAPI

Print basic Ukulele chords as text, JSON or HTML.

## CLI usage

With the CLI version, you can get a single chord with an optional key.

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

With the API version, you can get a single chord or multiple chords, both with an optional key.

```
# start local server on port 8080 (default)
go run main.go -serve
```

```
# option: if a different port is needed, start as follows (*adapt examples below accordingly*)
PORT=8081 go run main.go -serve
```

### Get chord(s)

- Get HTML

Browse to e.g. http://localhost:8080/api/v1/C or http://localhost:8080/api/v1/F/key or http://localhost:8080/api/v1/C-Am-G

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

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/api/v1/C-Am-G"
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

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/api/v1/C-Am-G"
```

- Get text

```
curl localhost:8080/api/v1/C
```

```
curl localhost:8080/api/v1/F/key
```

```
curl localhost:8080/api/v1/C-Am-G
```

## To do

- Host
- Improve support for multiple chords (vertical output, name per chord)