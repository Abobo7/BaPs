![introduce](https://socialify.git.ci/gucooing/BaPs/image?description=1&font=Source+Code+Pro&forks=1&issues=1&language=1&name=1&owner=1&pattern=Plus&pulls=1&stargazers=1&theme=Light)

# BaPs 🎮

> 🌟 由于是无状态设计,所以对内存的要求会略高
  
> 📅 当前支持版本：Japan 1.54.327262

## 📍Discord

![Discord](https://img.shields.io/badge/Join-Discord-blue?logo=discord&logoSize=auto
)

---
## 🚀 已实现功能
```
- 登录  
- 新手教程  
- 队伍管理  
- 抽卡  
- 剧情 待测试  
- 账号基础管理  
- MomoTalk  
- 邮件 全局/私人 收发管理  
- 角色养成管理  
- 背包管理  
- 副本 - 悬赏通缉 / 特别依赖 / 学院交流会 / 综合战术考试  
- 可恢复品自动恢复  
- 咖啡厅  
- 好友管理  
- 课程表  
- 社团  
- 战斗援助  
- 总力战  
- 彩奈登录奖励  
- 制约解除决战  
- 大决战  
- 商店
```
---
## 🛠️ 使用方法

#### 前置准备 (此步骤非常重要！！！)

1. 前往APKpure / Google Play 下载《ブルーアーカイブ》并安装游戏
2. 前往[Releases](https://github.com/gucooing/BaPs/releases/latest)下载最新的发行版本并拷贝到运行目录（请根据自己的系统进行下载）
3. 拷贝仓库的data文件夹到运行目录
3. 下载[Releases](https://github.com/gucooing/BaPs/releases/latest)中的Excel.bin文件,并替换到data文件夹中
4. 直接运行一次将会自动生成config.json文件,打开并编辑config.json文件
5. 运行

>若Excel.bin找不到请前往源代码中data文件夹下载
---

### 🐳 Docker部署
```bash
docker run -d \
  -p 5000:5000 \
  -v /data/baps/config.json:/usr/ba/config.json \
  -v /data/baps/sqlite/BaPs.db:/usr/ba/BaPs.db \
  -v /data/baps/sqlite/Rank.db:/usr/ba/Rank.db \
  ghcr.io/gucooing/baps:latest
```
---

## ⚙️ 配置说明
```
{
  "LogLevel": "info",
  "ResourcesPath": "./resources", // 发行版无用
  "DataPath": "./data",
  "GucooingApiKey": "123456", // 使用api时验证身份的key
  "AutoRegistration": true, // 是否自动注册
  "Tutorial": false, // 是否开启教程-不完善
  "HttpNet": {
    "InnerAddr": "0.0.0.0", // 监听地址
    "InnerPort": "5000", // 监听端口
    "OuterAddr": "10.0.0.3", // 外网地址
    "OuterPort": "5000", // 外网端口
    "Tls": false, // 是否启用ssl
    "CertFile": "./data/cert.pem",
    "KeyFile":   "./data/key.pem"
  },
  "GateWay": {
    "MaxPlayerNum": 0, // 最大在线玩家数
    "MaxCachePlayerTime": 720, // 最大玩家缓存时间
    "BlackCmd": {}, // 发行版无用
    "IsLogMsgPlayer": true // 发行版无用
  },
  "DB": {
    "dbType": "sqlite", // 使用的数据库类型,支持sqlite和mysql
    "dsn": "BaPs.db" // 数据库地址,如果是mysql请填写mysql url
  },
  "RaidRankDB": {
    "dbType": "sqlite", // 使用的数据库类型,支持sqlite和mysql
    "dsn": "Rank.db" // 数据库地址,如果是mysql请填写mysql url
  },
  "Irc": { // 可使用通用irc服务器地址
    "HostAddress": "127.0.0.1", // 社团聊天服务器irc地址
    "Port": 16666, // 社团聊天服务器irc端口
    "Password": "mx123" // 社团聊天服务器irc密码
  }
}
```
---

## 🌐 代理设置
转代以下地址:其中 http://127.0.0.1:5000 为服务器地址
```plaintext
https://ba-jp-sdk.bluearchive.jp  →  http://127.0.0.1:5000
https://yostar-serverinfo.bluearchiveyostar.com → http://127.0.0.1:5000
```
---

## ⛓️ 代理脚本
```python
# KitanoSakura
# 脚本还没完善，请使用WireGuard进行代理

from mitmproxy import http

# 定义重定向规则
redirects = {
    "https://ba-jp-sdk.bluearchive.jp": "http://127.0.0.1:5000",
    "https://prod-gateway.bluearchiveyostar.com:5100/api/gateway": "http://127.0.0.1:5000/getEnterTicket/gateway",
    "https://prod-game.bluearchiveyostar.com:5000/api/gateway": "http://127.0.0.1:5000/api/gateway",
    "https://prod-logcollector.bluearchiveyostar.com:5300": "http://127.0.0.1:5000/game/log",
}

def request(flow: http.HTTPFlow) -> None:
    # 判断请求的URL是否在重定向规则中
    for original_url, redirected_url in redirects.items():
        if flow.request.pretty_url.startswith(original_url):
            # 如果匹配，修改请求的URL为本地地址
            flow.request.url = flow.request.pretty_url.replace(original_url, redirected_url)
            print(f"Redirecting {original_url} to {redirected_url}")
            break
```

### 如何使用？

具体可前往以下docs查看
- [MitmProxy代理方案](Mitmproxy_Readme_ZH.md)

---

## ⌨️ GM工具
前往下方仓库下载GM工具以方便使用

- [BlueArchiveGM](https://github.com/PrimeStudentCouncil/BlueArchiveGM/releases/latest)

在线版GM菜单
注：内网环境不可使用在线版

- [BlueArchiveGM Web](https://gm.bluearchive.cc)

默认连接地址：http://127.0.0.1:5000 
默认密钥：123456 (可前往config.json进行更改)

---
## 🤝 参与贡献
我们欢迎所有想帮助我们的人加入，可通过以下方式进行帮助我们：
- 🐛 提交Issue报告问题
- 💡 提交Pull Request改进代码
- 📖 完善项目文档
- 🚀 加入Discord频道为我们提供建议
---

## ⚠️ 注意事项
1. 由于版权原因，dev使用的resources我们不会公开
2. 由于版权原因，部分源代码将不会被公开，但我们可以保证非公开部分代码无任何恶意内容
3. 玩家数据并不会实时保存到数据库中,如果有最新数据的需求,可通过api进行访问玩家数据
4. Api的使用过于复杂，没时间写docs自己研究，或自行下载GM工具调用

---
## 🤜 感谢名单

- 感谢 [zset](https://github.com/liyiheng/zset) 以此为基础实现排行榜