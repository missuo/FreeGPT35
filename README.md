# FreeGPT35
Utilize the unlimited free **GPT-3.5-Turbo** API service provided by the login-free ChatGPT Web. Forked from [https://github.com/skzhengkai/free-chatgpt-api](https://github.com/skzhengkai/free-chatgpt-api).

## Usage
### Node

```bash
npm install
node app.js
```
### Docker

```bash
docker run -p 3040:3040 missuo/freegpt35
```

### Docker Compose

```bash
mkdir freegpt35 && cd freegpt35
wget -O compose.yaml https://raw.githubusercontent.com/missuo/FreeGPT35/main/compose.yaml
docker compose up -d
```
