# ElasticSearch 
Elasticsearch 是一个分布式、RESTful 风格的搜索和数据分析引擎，能够解决不断涌现出的各种用例。

- 全文搜索引擎：
- 快速的存储，搜索和分析海量数据：

## 安装和配置
### Elasticsearch
- Download and unzip Elasticsearch(https://www.elastic.co/cn/downloads/elasticsearch)
- Run `bin/elasticsearch` (or `bin\elasticsearch.bat`on Windows)
- Run curl `http://localhost:9200/` or Invoke-RestMethod `http://localhost:9200` with PowerShell
### Docker
1. 拉取镜像
```
docker pull docker.elastic.co/elasticsearch/elasticsearch:6.3.2
```
2. 运行容器
```
docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:6.3.2
```
3. 配置跨域
- 进入容器
由于要进行配置，因此需要进入容器当中修改相应的配置信息。
```
docker exec -it es /bin/bash
```
- 进行配置
```
# 显示文件
ls
结果如下：
LICENSE.txt  README.textile  config  lib   modules
NOTICE.txt   bin             data    logs  plugins

# 进入配置文件夹
cd config


# 修改配置文件
vi elasticsearch.yml

# 加入跨域配置
http.cors.enabled: true
http.cors.allow-origin: "*"
```

- 重启容器
由于修改了配置，因此需要重启ElasticSearch容器。
```
docker restart es
```
4. 测试
浏览器输入：`localhost:9200`
## 安装ElasticSearche的客户端的包
* https://www.elastic.co/guide/en/elasticsearch/client/index.html
- 使用Go包管理工具管理依赖如dep
```
dep ensure -add github.com/olivere/elastic
```
- 导入包
```
import "github.com/olivere/elastic"
```
## 参考资料
* https://www.elastic.co/guide/en/elasticsearch/reference/current/getting-started.html