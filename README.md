
# Telegram Webhook

A middleware API to convert webhook from supported services to Telegram messages

## Installation

Build with Go

```bash
  go build
  ./telegram-webhook
```
A server will start on 8080 port.

#### Enviroment variables:

You can create a `.env` on project root folder
```dotenv
PORT=8081 # Choose a diferent port
TELEGRAM_BOT_DEBUG=true # Show debug information from Telegram Bot
```
## API Reference

### Bitbucket WebHook

```http
  POST /bitbucket
```

| Query Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `bot_token` | `string` | **Required**. Your Telegram Bot Token |
| `chat_id` | `string` | **Required**. Where to send messages |

#### Supported triggers:
- Pipelines