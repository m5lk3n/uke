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

Below are examples of how to query chords. Up to 4 are printed.

#### Get HTML

Browse to
- http://localhost:8080/api/v1/chord/C
- http://localhost:8080/api/v1/chord/F/key
- http://localhost:8080/api/v1/chords/C-Am-G
- http://localhost:8080/api/v1/chords/C-Am-G/key

or

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/api/v1/chord/C"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/api/v1/chord/F/key"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/api/v1/chords/C-Am-G"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/api/v1/chords/C-Am-G/key"
```

#### Get JSON

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/api/v1/chord/C"
```

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/api/v1/chord/F/key"
```

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/api/v1/chords/C-Am-G"
```

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/api/v1/chords/C-Am-G/key"
```

#### Get text

```
curl localhost:8080/api/v1/chord/C
```

```
curl localhost:8080/api/v1/chord/F/key
```

```
curl localhost:8080/api/v1/chords/C-Am-G
```

```
curl localhost:8080/api/v1/chords/C-Am-G/key
```

## To do

- Host
- Support chordNames on CLI, document chordNames