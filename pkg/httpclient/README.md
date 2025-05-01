# httpclient

- recomend use [go-resty/resty](https://github.com/go-resty/resty) Simple HTTP, REST, and SSE client library for Go

1. 重试机制：
   - 支持配置最大重试次数
   - 默认重试3次
   - 可自定义哪些HTTP状态码需要重试（默认408, 429, 500, 502, 503, 504）

2. 退避睡眠：
   - 指数退避算法
   - 添加随机抖动避免惊群效应
   - 可配置最小和最大等待时间

3. 日志记录：
   - 内置默认日志实现
   - 支持自定义日志记录器
   - 记录请求尝试、失败和成功信息

4. 便捷方法：
   - 提供Get、Post等便捷方法
   - 支持context.Context

5. 可配置性：
   - 通过Option模式灵活配置
   - 可替换底层http.Client
