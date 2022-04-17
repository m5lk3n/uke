# UkeAPI

Prints basic Ukulele chords as text, JSON or HTML.

This README mainly describes the interface usage; more information about build and deployment can be found in the [Makefile](Makefile) as well as at the end of this README.

See also "Hosted Version" below to try things out online.

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

To obtain the list of supported chord names:

```
go run main.go -chordNames
A
A7
Am
Bm
C
C7
Cm
D
Dm
E
Em
F
G
G7
Gbm
Gm
```

CLI parameters can be combined.

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

Below are examples of how to query chords. A single chord is printed for all supported formats; up to 4 are printed in HTML mode.

#### As HTML

Browse to
- http://localhost:8080/ukeapi/v1/chord/C
- http://localhost:8080/ukeapi/v1/chord/F/key
- http://localhost:8080/ukeapi/v1/chords/C-Am-G
- http://localhost:8080/ukeapi/v1/chords/C-Am-G/key

or

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/ukeapi/v1/chord/C"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/ukeapi/v1/chord/F/key"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/ukeapi/v1/chords/C-Am-G"
```

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/ukeapi/v1/chords/C-Am-G/key"
```

#### As JSON

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/ukeapi/v1/chord/C"
```

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/ukeapi/v1/chord/F/key"
```

#### As text

```
curl localhost:8080/ukeapi/v1/chord/C
```

```
curl localhost:8080/ukeapi/v1/chord/F/key
```

### Get chord(s)

#### As HTML

Browse to
- http://localhost:8080/ukeapi/v1/chordNames

or

```
curl -X GET \
  -H "Accept: text/html" \
  "http://localhost:8080/ukeapi/v1/chordNames"
```

#### As JSON

```
curl -X GET \
  -H "Accept: application/json" \
  "http://localhost:8080/ukeapi/v1/chordNames"
```

#### As text

```
curl localhost:8080/ukeapi/v1/chordNames
```

## Hosted Version

The hosted version is here:
- https://lttl.dev/ukeapi/v1/version
- https://lttl.dev/ukeapi/v1/healthy
- https://lttl.dev/ukeapi/v1/ready
- https://lttl.dev/ukeapi/v1/chordNames
- https://lttl.dev/ukeapi/v1/chord/C
- https://lttl.dev/ukeapi/v1/chord/F/key
- https://lttl.dev/ukeapi/v1/chords/C-Am-G
- ...

## Deployment Configuration - *Exemplified!*

**Disclaimer: No guarantee, no warranty, no support! This is not production grade!**

Prerequisites on the target Linux system:
- Create a local user and group, both called `ukeapi`, e.g. by `adduser --system --no-create-home --group ukeapi`
- `mkdir /opt/ukeapi` and change ownership to `ukeapi` for both, user and group: `chown ukeapi:ukeapi /opt/ukeapi/`
- Install nginx and point my.tld to it
- Create a public key pair so that SSH works for `ukeapi` to the target machine 

On the source system (where this repo resides and which can SSH into the target): 

```
DEPLOY_TARGET=ukeapi@my.tld:/opt/ukeapi make deploy
```

On the my.tld target host:

For an nginx deployment, add the following snippet to `/etc/nginx/sites-available`:

```
         location /ukeapi {
           proxy_pass  http://127.0.0.1:9000;
           proxy_set_header Host $host;
           proxy_set_header X-Real-IP $remote_addr;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
         }
```

Start UkeAPI:
```
cd /opt/ukeapi
GIN_MODE=release PORT=9000 /opt/ukeapi/ukeapi -serve >> /opt/ukeapi/ukeapi.log &
```

Restart nginx: `systemctl restart nginx.service`

Check nginx status: `systemctl status nginx.service`