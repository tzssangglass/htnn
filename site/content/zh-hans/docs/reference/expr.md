---
title: 表达式
---

HTNN 许多地方使用了 CEL（Common Expression Language） 作为运行时动态执行的表达式。关于 CEL 的语法，请参考[官方文档](https://github.com/google/cel-spec)。为了让 CEL 可以在网络代理的上下文中大显身手，HTNN 提供了一套 CEL 拓展库以供表达式访问相关的信息。本文档将描述该 CEL 拓展库。

## 请求

| 名称               | 类型   | 说明                                    |
|--------------------|--------|-----------------------------------------|
| request.path()     | string | 请求的 path，如 `/x?a=1`                |
| request.url_path() | string | 请求的 path，去掉 query string，如 `/x` |
| request.host()     | string | 请求的 host                             |
| request.scheme()   | string | 请求的 scheme，小写形式，如 `http`      |
| request.method()   | string | 请求的 method，大写形式，如 `GET`       |
| request.id()       | string | `x-request-id` 请求头中的 ID            |