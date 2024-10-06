[中文](about:blank)|[English](about:blank)

# 简述：
伴随着物联网技术的逐渐普及，网络的复杂度越来越高，面对海量的设备信息数据，例如智能驾驶，工业设备等，HTTP/HTTPS协议在如此大规模的系统中很难满足响应需求，而MQTT以其轻巧高效、可靠安全、双向通讯等诸多优势，已然在物联网行业中成为主流协议。

如何实现千万级消息并发？随着物联网架构技术的不断迭代，目前常用解决方案的途径为：

![画板](https://cdn.nlark.com/yuque/0/2024/jpeg/29382993/1728182548389-16faf499-cee7-40e9-9aaf-47a636fe46e5.jpeg)

        该项目为EMQX高可用集群的中间件应用，旨在通过简单配置的方式，配合容器，即可实现大规模集群化，并应用在实际生产中。

# 项目功能：
该项目可通过配置文件的方式，即可实现多Mqtt客户端消息接受、分发、过滤、分发重试等，现已实现HTTP/HTTPS服务的分发，后续会迭代加入其他服务协议的分发机制。

# 结构图：
![画板](https://cdn.nlark.com/yuque/0/2024/jpeg/29382993/1728117289099-8151b2a6-e174-4e6f-9544-4e82bc7b2a3e.jpeg)

# 使用方法：
## 配置文件：
| 名称 | 类型 | 必选 | 说明 |
| --- | --- | --- | --- |
| mqtt_brokers | list[MqBrokerConfig] | true | mqtt配置 |
| log_config | struct | true | 日志输出配置 |


### MqBrokerConfig:
| 名称 | 类型 | 必选 | 说明 |
| --- | --- | --- | --- |
| client_id | str | 是 | client唯一标识 |
| username | str | 否 | 连接用户名 |
| password | str | 否 | 连接密码 |
| alive | int | 是 | 连接活跃检测时间 |
| broker_ip | str | 是 | 连接ip |
| broker_port | int | 是 | 连接端口 |
| sub_deal_config | list[topicConfig] | 是 | topic转发处理配置 |


topicConfig：

| 名称 | 类型 | 必选 | 说明 |
| --- | --- | --- | --- |
| app_name | str | 否 | 转发的服务名称 |
| app_id | str | 是 | 服务的唯一id |
| enabled | bool | 是 | 是否启用 |
| callbackMethod | str | 是 | 回调方式：<br/>目前仅支持：HTTP/HTTPS |
| callbaccallbackAddress | list[str] | 是 | 回调地址，支持多地址 |
| subTopic | struct | 是 | 订阅的子topic配置，支持通配符 |
| >>>topic | str | 是 | 订阅的topic，支持通配符 |
| >>>qos | int | byte | 是 | qos等级 |
| excludeTopics | list[str] | 否 | 需要排除的topic |
| retry | int | 是 | 消息处理异常重试次数 |


### log_config:  
| 名称 | 类型 | 必选 | 说明 |
| --- | --- | --- | --- |
| level | str | 是 | 日志记录等级 |
| filename | str | 是 | 日志文件名 |
| maxsize | int | 是 | 日志大小，单位MB |
| max_age | int | 是 | 日志保留的天数，单位天 |
| max_backups | int | 是 | 日志保留的最大数量 |


## 示例：
```json
{
  "mqtt_brokers": [
    {
      "client_id": "go-mqtt-client-1",
      "username": "",
      "password": "",
      "alive": 60,
      "broker_ip": "127.0.0.1",
      "broker_port": 1883,
      "sub_deal_config": {
        "app_name": "test1",
        "app_id": "123",
        "enabled": true,
        "callbackMethod": "HTTP",
        "callbackAddress": [
          "http://127.0.0.1:8000/api/test"
        ],
        "subTopic": {
          "topic": "from/v1/#",
          "qos": 0
        },
        "excludeTopics": [
          "from/v1/aaa/"
        ],
        "retry": 5
      }
    },
    {
      "client_id": "go-mqtt-client-2",
      "username": "",
      "password": "",
      "alive": 60,
      "broker_ip": "127.0.0.1",
      "broker_port": 1883,
      "sub_deal_config": {
        "app_name": "test2",
        "app_id": "123",
        "enabled": true,
        "callbackMethod": "HTTP",
        "callbackAddress": [
          "http://127.0.0.1:8000/api/test2"
        ],
        "subTopic": {
          "topic": "from/v1/#",
          "qos": 0
        },
        "excludeTopics": [
          "from/v1/aaa/"
        ],
        "retry": 5
      }
    }
  ],
  "log_config": {
    "level": "info",
    "filename": "logs/mqtt.log",
    "maxsize": 1,
    "max_age": 7,
    "max_backups": 3
  }
}
```

# 消息格式：
MQTT消息格式分为header和body，消息格式如下：

```json
{
  "header":"2024-10-06 10:57:30.220911642 +0800 CST m=+108.080484125",
  "body":"2024-10-06 10:57:30.220927581 +0800 CST m=+108.080500051"
}
```

# 测试：
需要结合配置文件一起使用

## py服务测试
以下是使用py的fastapi进行简单的分发服务测试

依赖安装：

```python
pip install fastapi
pip intsall uvicorn
```

简单的转发服务器：

```python
import uvicorn
from fastapi import FastAPI

app = FastAPI()

@app.post('/api/test')
async def test(data: dict = {}):
    print("test1:", data)
    return {"code": 200, "msg": data}


@app.post('/api/test2')
async def test2(data: dict = {}):
    print("test2:", data)
    return {"code": 200, "msg": data}


if __name__ == '__main__':
    uvicorn.run("start:app", host="127.0.0.1", port=8000, workers=1)

```

运行：

```python
uvicorn main:app --reload
```

# 功能与迭代：
## 支持：
已实现多Mqtt客户端消息接受、分发、过滤、分发重试等

不同协议分发：

现已实现HTTP/HTTPS服务的分发，后续会迭代加入其他服务协议的分发机制。

## 迭代：
后续考虑接入GPRC、数据库、其他中间件服务（例如kafka等）

消息格式可配置等



