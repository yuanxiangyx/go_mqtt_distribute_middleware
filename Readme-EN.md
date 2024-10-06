[Chinese](about:blank)|[English](about:blank)

# Sketch：
With the gradual popularization of IoT technology, the complexity of networks is increasing. Faced with massive amounts of device information data, such as intelligent driving, industrial equipment, etc., HTTP/HTTPS protocol is difficult to meet response needs in such a large-scale system. MQTT, with its advantages of lightweight, high efficiency, reliable security, and bidirectional communication, has become the mainstream protocol in the IoT industry.

How to achieve millions of concurrent messages? With the continuous iteration of IoT architecture technology, the commonly used solutions currently include:

![画板](https://cdn.nlark.com/yuque/0/2024/jpeg/29382993/1728182548389-16faf499-cee7-40e9-9aaf-47a636fe46e5.jpeg)

        This project is a middleware application for EMQX high availability cluster, aiming to achieve large-scale clustering through simple configuration and container cooperation, and apply it in actual production.

# Project Function：
This project can achieve multi Mqtt client message reception, distribution, filtering, distribution retry, etc. through configuration files. HTTP/HTTPS service distribution has been implemented, and distribution mechanisms for other service protocols will be iteratively added in the future.

# Structural diagram：
![画板](https://cdn.nlark.com/yuque/0/2024/jpeg/29382993/1728117289099-8151b2a6-e174-4e6f-9544-4e82bc7b2a3e.jpeg)

# Usage Method：
## Configuration File：
| Name | Type | Mandatory  | Explain |
| --- | --- | --- | --- |
| mqtt_brokers | list[MqBrokerConfig] | true | MQTT configuration |
| log_config | struct | true | Log output configuration |


### MqBrokerConfig:
| Name | Type | Mandatory  | Explain |
| --- | --- | --- | --- |
| client_id | str | false | Unique identifier of client |
| username | str | false | Connect username |
| password | str | false | Connection password |
| alive | int | true | Connection activity detection time |
| broker_ip | str | true | Connect IP |
| broker_port | int | true | Connection port |
| sub_deal_config | list[topicConfig] | true | Topic forwarding processing configuration |


topicConfig：

| Name | Type | Mandatory  | Explain |
| --- | --- | --- | --- |
| app_name | str | true | Forwarding service name |
| app_id | str | true | Unique ID of the service |
| enabled | bool | true | Is it enabled |
| callbackMethod | str | true | Callback method:<br/>Currently only supports HTTP/HTTPS |
| callbaccallbackAddress | list[str] | true | Callback address, supports multiple addresses |
| subTopic | struct | true | Subtopic configuration for subscription |
| >>>topic | str | true | The subscribed topic supports wildcard characters |
| >>>qos | int | byte | true | QoS level |
| excludeTopics | list[str] | false | Topics that need to be excluded |
| retry | int | true | Number of message processing exception retries |


### log_config:  
| Name | Type | Mandatory  | Explain |
| --- | --- | --- | --- |
| level | str | true | Logging level |
| filename | str | true | Log file name |
| maxsize | int | true | Log size, in MB |
| max_age | int | true | The number of days to retain logs, in days |
| max_backups | int | true | Maximum number of logs retained |


## Example：
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

# Message Format ：
The MQTT message format is divided into header and body, and the message format is as follows:

```json
{
  "header":"2024-10-06 10:57:30.220911642 +0800 CST m=+108.080484125",
  "body":"2024-10-06 10:57:30.220927581 +0800 CST m=+108.080500051"
}
```

# Test：
Need to be used in conjunction with configuration files

## Py service testing
The following is a simple distribution service test using Py's FastAPI

Dependency installation:

```python
pip install fastapi
pip intsall uvicorn
```

Simple forwarding server:

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

function:

```python
uvicorn main:app --reload
```

# Function and iteration：
## support：
Multiple Mqtt client message reception, distribution, filtering, distribution retry, etc. have been implemented.

Distribution of different protocols:

The distribution of HTTP/HTTPS services has been implemented, and distribution mechanisms for other service protocols will be iteratively added in the future.

## Iteration：
Consider integrating GPRC, databases, and other middleware services (such as Kafka, etc) in the future

Message format can be configured, etc



