# INFO

Golang websocket服务器 用于机器人的传感器等数据的通信传输

# 用法

## 创建新的传感器

1. message下面创建消息类型
2. sensor下面创建对象
3. server下面创建服务, `server/main`当中添加路由

+ 继承均来自每个包的 `main`