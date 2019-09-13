# HTTP 调用添加自定义处理逻辑

## 编译测试

```
curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user

{"level":"INFO","timestamp":"2018-10-19 18:45:27.058","file":"middleware/logging.go:80","msg":"26.0015ms     | 127.0.0.1    | GET /v1/user | {code: 0, message: OK}"}
{"level":"INFO","timestamp":"2018-10-19 18:45:35.138","file":"middleware/logging.go:80","msg":"24.0014ms     | 127.0.0.1    | GET /v1/user | {code: 0, message: OK}"}

```