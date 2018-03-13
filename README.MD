# wechat-tools
wechat-tools is common tools (wechat-tools 是一些常用的微信工具集)

include:(目前包含的工具)
* wechat mini programs AES decrypt tools - 小程序AES密文解密(用于微信群ID，用户信息解密等)


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

