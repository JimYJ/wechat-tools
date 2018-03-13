# wechat-tools
wechat-tools is common tools (wechat-tools 是一些常用的微信工具集)

include:(目前包含的工具)
* wechat mini programs AES decrypt tools - 小程序AES密文解密(用于微信群ID，用户信息解密等)
* get wechat Access Token - 获取公众号普通Access Token
* get wechat Jsapi Ticket - 获取公众号Jsapi Ticket

next include:(即将支持的功能)
* add cache for Access Token & Ticket - 添加Access Token和Ticket的缓存功能，防止反复请求接口

## How to get

```
go get github.com/JimYJ/wechat-tools
```

## Usage

**import:**

```go
import "github.com/JimYJ/wechat-tools"
```

wechat mini programs AES decrypt:

```go
results, err := wechatools.Decrypt(encryptedData, iv, sessionKey, APPID)
```

get wechat Access Token:
```go
accessToken, err := GetAccessToken(config.APPID, config.SECRETKEY)
```

get wechat Jsapi Ticket:
```go
ticket, err := GetJsapiTicket(accessToken)
```
