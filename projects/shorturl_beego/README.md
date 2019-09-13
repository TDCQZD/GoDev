# shorturl_beego
使用beego开发API应用. 他包含了两个API接口:

- /v1/shorten
- /v1/expand
## 新建项目
```
bee api shorturl_beego
```

## 测试
```
# shortening url example
http://localhost:8080/v1/shorten/?longurl=http://www.baidu.com

{
  "UrlShort": "5laZG",
  "UrlLong": "http://www.baidu.com"
}

# expanding url example
http://localhost:8080/v1/expand/?shorturl=5laZG

{
  "UrlShort": "5laZG",
  "UrlLong": "http://www.baidu.com"
}
```