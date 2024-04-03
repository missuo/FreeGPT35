[![Docker Pulls][1]](https://hub.docker.com/r/missuo/freegpt35)

[1]: https://img.shields.io/docker/pulls/missuo/freegpt35?logo=docker

Utilize the unlimited free **GPT-3.5-Turbo** API service provided by the login-free ChatGPT Web. Forked from [https://github.com/skzhengkai/free-chatgpt-api](https://github.com/skzhengkai/free-chatgpt-api).

## Deploy
### Node

```bash
npm install
node app.js
```
### Docker

```bash
docker run -p 3040:3040 ghcr.io/missuo/freegpt35
```

```bash
docker run -p 3040:3040 missuo/freegpt35
```

### Docker Compose

```bash
mkdir freegpt35 && cd freegpt35
wget -O compose.yaml https://raw.githubusercontent.com/missuo/FreeGPT35/main/compose.yaml
docker compose up -d
```

## Request Example

**You don't have to pass Authorization, of course, you can also pass any string randomly.**

```bash
curl http://127.0.0.1:3040/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer any_string_you_like" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [
      {
        "role": "user",
        "content": "Hello!"
      }
    ],
    "stream": true
    }'
```
