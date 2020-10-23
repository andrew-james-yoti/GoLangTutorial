# Webserver

- [Tutorial](https://www.freecodecamp.org/news/how-to-build-a-web-app-with-go-gin-and-react-cffdc473576/)
- [Modules](https://blog.golang.org/using-go-modules)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

## Usage

- GET    /api/jokes
```bash
curl http://localhost:8080/api/jokes
```

- POST   /api/jokes/like/:jokeID
```bash
curl -X POST http://localhost:8080/api/jokes/like/1
```