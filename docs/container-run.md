# Container Run

Build the application image:

```sh
docker build -t invaders:local .
```

Run it locally:

```sh
docker run --rm -p 8080:8080 invaders:local
curl http://127.0.0.1:8080/healthz
curl http://127.0.0.1:8080/ | grep '<title>'
```

The runtime image copies `client/` into `/app/client`, so the container serves the same primary UI from `/` that local `cd server && go run .` serves.
