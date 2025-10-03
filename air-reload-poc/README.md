# Air Reload POC

## How to Run

```bash
docker-compose up --build -d
curl localhost:8080
# {"text":"Hello, Docker + Go + Air + Gorilla/Mux + UUID! 🚀"}
```

Change some thing in the code, then, call the url again.
Inspect if your change is in there.

```bash
curl localhost:8080
```
