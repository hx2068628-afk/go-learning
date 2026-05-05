# Gin Redirect 和 HandleContext 示例说明

这个项目用 Gin 演示两个容易混淆的概念：

1. `c.Redirect(...)`：让浏览器重新请求另一个地址。
2. `c.Request.URL.Path = "..."; r.HandleContext(c)`：在服务器内部重新匹配一次路由。

当前代码：

```go
r.GET("/", func(c *gin.Context) {
    c.Redirect(http.StatusFound, "/a")
})

r.GET("/a", func(c *gin.Context) {
    c.JSON(http.StatusOK, "a")
    c.Request.URL.Path = "/b"
    r.HandleContext(c)
    c.JSON(http.StatusOK, "c")
})

r.GET("/b", func(c *gin.Context) {
    c.JSON(http.StatusOK, "bd")
})
```

## 访问 `/`

访问：

```text
http://localhost:8080/
```

会执行：

```go
c.Redirect(http.StatusFound, "/a")
```

这是真正的 HTTP 重定向。服务端会返回类似这样的响应：

```text
302 Found
Location: /a
```

浏览器收到以后，会重新请求：

```text
http://localhost:8080/a
```

所以地址栏会从 `/` 变成 `/a`。

## 访问 `/a`

访问：

```text
http://localhost:8080/a
```

会进入 `/a` 对应的 handler：

```go
r.GET("/a", func(c *gin.Context) {
    c.JSON(http.StatusOK, "a")
    c.Request.URL.Path = "/b"
    r.HandleContext(c)
    c.JSON(http.StatusOK, "c")
})
```

执行顺序是：

```text
1. 进入 /a
2. 输出 "a"
3. 把当前请求对象里的 Path 改成 /b
4. 调用 r.HandleContext(c)
5. Gin 根据新的 Path 匹配到 /b
6. 执行 /b 的 handler，输出 "bd"
7. /b 执行完，回到 /a
8. 继续执行最后一行，输出 "c"
```

所以浏览器最终看到：

```text
"a""bd""c"
```

## 为什么 Path 变成 `/b` 后还能输出 `"c"`

因为：

```go
c.Request.URL.Path = "/b"
```

只改了请求对象里的路径字段。

它不会让当前 `/a` 函数结束，也不会让浏览器真的跳转到 `/b`。

`r.HandleContext(c)` 的作用是：用当前这个 `gin.Context` 再交给 Gin 路由器匹配一次。因为此时 `Path` 已经是 `/b`，所以会执行 `/b` 的 handler。

但是 `/b` 执行完以后，程序会回到这行代码的下一行：

```go
c.JSON(http.StatusOK, "c")
```

所以 `"c"` 仍然会输出。

## `"c"` 是在什么路径下输出的

严格来说，`"c"` 不是“在某个路径下输出”的。

HTTP 里一次请求只有一个响应体。浏览器这次请求的是 `/a`，服务端就把所有内容都写进这一次 `/a` 请求的响应体里。

`gin.Context` 里面可以简单理解成有两部分：

```text
c.Request  请求信息，比如 URL.Path
c.Writer   响应写入器，用来把内容写回浏览器
```

执行：

```go
c.Request.URL.Path = "/b"
```

改变的是：

```text
c.Request.URL.Path
```

没有改变的是：

```text
c.Writer
```

而：

```go
c.JSON(http.StatusOK, "c")
```

使用的是 `c.Writer` 写响应体，不是根据 `c.Request.URL.Path` 找一个新的输出位置。

因此，即使当时 `c.Request.URL.Path` 已经是 `/b`，`"c"` 还是写进同一个响应体里，也就是浏览器访问 `/a` 得到的响应里。

## `Redirect` 和 `HandleContext` 的区别

| 写法 | 本质 | 浏览器地址栏会变吗 | 会不会产生新的浏览器请求 | 当前函数会自动结束吗 |
| --- | --- | --- | --- | --- |
| `c.Redirect(http.StatusFound, "/a")` | 告诉浏览器去请求 `/a` | 会 | 会 | 不会 |
| `c.Request.URL.Path = "/b"; r.HandleContext(c)` | 服务端内部重新匹配路由 | 不会 | 不会 | 不会 |

## 注意事项

`c.Redirect(...)` 不会自动 `return`。

如果重定向后不想继续执行后面的代码，应该写：

```go
r.GET("/", func(c *gin.Context) {
    c.Redirect(http.StatusFound, "/a")
    return
})
```

同理，`r.HandleContext(c)` 也不会自动结束当前函数。它只是执行一次新的路由匹配，匹配到的 handler 执行完以后，还会回到原来的位置继续往下执行。

## 总结

- `Redirect` 是浏览器层面的跳转。
- `HandleContext` 是服务器内部的路由分发。
- `c.Request.URL.Path` 改的是请求信息，不是响应输出位置。
- `c.JSON(...)` 写的是当前 HTTP 响应体，所以 `"a"`、`"bd"`、`"c"` 会连续出现在同一次响应里。
