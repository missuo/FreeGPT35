[![Docker Pulls][1]](https://hub.docker.com/r/missuo/freegpt35)

[1]: https://img.shields.io/docker/pulls/missuo/freegpt35?logo=docker

Utilize the unlimited free **GPT-3.5-Turbo** API service provided by the login-free ChatGPT Web.


> [!IMPORTANT]  
> **If you are unable to use this project normally, it is most likely due to issues with your IP. Your IP has triggered Cloudflare's shield, or has already been banned. Please try to change your IP or switch servers on your own. 如果您无法正常使用此项目，很可能是由于您的 IP 存在问题。您的 IP 已触发了 Cloudflare 的盾，或已被 ban 掉。请尝试自行更改您的 IP 或切换服务器。**

## Please READ the following content carefully!
- Please do not use the IP provided by proxy providers, otherwise you probably won't be able to use it. 请不要使用机场的 IP，不然你大概率无法使用。
- Do not make frequent requests, such as using **immersive translate**. 不要频繁请求，例如使用沉浸式翻译。
- Recommended to use US home broadband IP, you are very likely to succeed. 推荐使用美国家宽IP，你很大可能可以成功。
- Don't share and abuse your API. 不要共享和滥用你的 API。

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

#### Only FreeGPT35 Service

```bash
mkdir freegpt35 && cd freegpt35
wget -O compose.yaml https://raw.githubusercontent.com/missuo/FreeGPT35/main/compose.yaml
docker compose up -d
```

#### FreeGPT35 Service with [ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web):

```bash
mkdir freegpt35 && cd freegpt35
wget -O compose.yaml https://raw.githubusercontent.com/missuo/FreeGPT35/main/compose_with_next_chat.yaml
docker compose up -d
```

After deployment, you can directly access `http://[IP]:3040/v1/chat/completions` to use the API. Or use `http://[IP]:3000` to directly use ChatGPT-Next-Web.

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

## Compatibility

You can use it in any app, such as OpenCat, Next-Chat, Lobe-Chat, Bob, etc. Feel free to fill in an **API Key** with any string, for example, `gptyyds`.

### Bob
![Bob](./img/bob.png)

## Credits
- Forked From: [https://github.com/skzhengkai/free-chatgpt-api](https://github.com/skzhengkai/free-chatgpt-api)
- Original Author: [https://github.com/PawanOsman/ChatGPT](https://github.com/PawanOsman/ChatGPT)

## License
MIT License
