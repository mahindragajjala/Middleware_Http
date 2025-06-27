# Middleware_Http

| Section | Topic                 | Summary                                                | Sample Code Snippet                               |
| ------- | --------------------- | ------------------------------------------------------ | ------------------------------------------------- |
| 1.1     | What is Middleware?   | A function that intercepts HTTP requests/responses.    | `func Middleware(next http.Handler) http.Handler` |
| 1.2     | Why Middleware?       | For logging, auth, rate-limit, recovery, etc.          | -                                                 |
| 1.3     | Role in Web Framework | Injects logic before/after request handlers.           | -                                                 |
| 1.4     | net/http Arch         | Middleware chains handlers using `http.Handler`.       | `next.ServeHTTP(w, r)`                            |
| 1.5     | Gin Arch              | Uses `gin.HandlerFunc` for middleware chaining.        | `func(c *gin.Context) { c.Next() }`               |
| 1.6     | Execution Flow        | Executes in order of addition; `Next()` moves to next. | -                                                 |

---

| Section | Middleware Type | Purpose                 | Code Sample                                          |
| ------- | --------------- | ----------------------- | ---------------------------------------------------- |
| 2.1     | Custom net/http | Base template.          | `http.Handle("/", Middleware(http.HandlerFunc(fn)))` |
| 2.2     | Custom Gin      | Base template.          | `router.Use(MyGinMiddleware())`                      |
| 2.3     | Logging         | Log before/after.       | `log.Println(r.URL.Path)`                            |
| 2.4     | Header Modify   | Add header to response. | `w.Header().Set("X-Custom", "true")`                 |
| 2.5     | Request Time    | Track latency.          | `start := time.Now()`                                |
| 2.6     | Gzip            | Compress response.      | `gzip.NewWriter(w)`                                  |

---

| Section | Auth Middleware | Use                      | Code Idea              |
| ------- | --------------- | ------------------------ | ---------------------- |
| 3.1     | Auth Intro      | Secure endpoints.        | `checkToken(r)`        |
| 3.2     | JWT             | Token validation.        | `jwt.Parse(token)`     |
| 3.3     | Basic Auth      | Username/password check. | `r.BasicAuth()`        |
| 3.4     | OAuth2          | Token validation.        | Use 3rd-party libs     |
| 3.5     | Role Auth       | RBAC control.            | `user.Role == "admin"` |
| 3.6     | IP Whitelist    | Allow trusted IPs.       | `if ip == allowed`     |

---

| Section | Request/Response   | Goal                         | Snippet                          |
| ------- | ------------------ | ---------------------------- | -------------------------------- |
| 4.1     | Read Req Headers   | Inspect auth tokens.         | `r.Header.Get("Auth")`           |
| 4.2     | Modify Res Headers | Add security headers.        | `w.Header().Set(...)`            |
| 4.3     | Context Injection  | Add data to req.             | `context.WithValue(r.Context())` |
| 4.4     | Custom Fields      | Inject structs into context. | `c.Set("key", value)` (Gin)      |
| 4.5     | Read Body          | Access request body.         | `ioutil.ReadAll(r.Body)`         |

---

| Section | Error Handling | Purpose               | Sample                             |
| ------- | -------------- | --------------------- | ---------------------------------- |
| 5.1     | Central Error  | One place for errors. | `c.AbortWithStatusJSON(...)`       |
| 5.2     | Panic Recovery | Avoid crash.          | `defer recover()`                  |
| 5.3     | Global Except. | Handle all errors.    | `router.Use(ErrorHandler)`         |
| 5.4     | Custom Errors  | User-friendly errors. | `c.JSON(400, gin.H{"error": ...})` |

---

| Section | Security       | Type                  | Code                                            |
| ------- | -------------- | --------------------- | ----------------------------------------------- |
| 6.1     | Rate Limit     | Limit requests.       | Token bucket, leaky bucket                      |
| 6.2     | IP Block       | Blocklist IPs.        | `if ip in blocklist`                            |
| 6.3     | CORS           | Cross-origin control. | `w.Header().Set("Access-Control-Allow-Origin")` |
| 6.4     | CSRF           | Token verification.   | Match CSRF tokens                               |
| 6.5     | Secure Headers | Add security headers. | XSS, CSP headers                                |

---

| Section | 3rd Party Middleware | Tool                 | Example                      |
| ------- | -------------------- | -------------------- | ---------------------------- |
| 7.1     | Popular Libs         | Reusable packages.   | gorilla/handlers             |
| 7.2     | Gin Ecosystem        | Auth, cors, etc.     | gin-contrib/cors             |
| 7.3     | Gorilla Handlers     | Logging, compress.   | handlers.LoggingHandler(...) |
| 7.4     | Negroni              | Middleware chaining. | `negroni.New(...)`           |
| 7.5     | Prometheus           | Track metrics.       | `promhttp.Handler()`         |

---

| Section | Composition  | Concept              | Code Idea                         |
| ------- | ------------ | -------------------- | --------------------------------- |
| 8.1     | Chaining     | Combine multiple.    | `middleware1(middleware2(final))` |
| 8.2     | Order        | First in, first run. | Add order-aware logs              |
| 8.3     | Templates    | Reusable logic.      | Return middleware function        |
| 8.4     | Configurable | Accept config/data.  | `MyMiddleware(cfg)`               |

---

| Section | Patterns     | Idea                       | Code                              |
| ------- | ------------ | -------------------------- | --------------------------------- |
| 9.1     | Decorator    | Wrap logic.                | return func(next)                 |
| 9.2     | Functional   | Functions return handlers. | `func Auth() gin.HandlerFunc`     |
| 9.3     | Pipeline     | Process chain of logic.    | middleware stack                  |
| 9.4     | Conditional  | Route-specific.            | `if path == "/admin"`             |
| 9.5     | Route Groups | Group middleware.          | `admin := router.Group("/admin")` |

---

| Section | Testing       | Area                   | Tool/Code                |
| ------- | ------------- | ---------------------- | ------------------------ |
| 10.1    | Unit Testing  | Test logic only.       | std `testing` pkg        |
| 10.2    | Mock Req/Res  | Custom inputs.         | `httptest.NewRecorder()` |
| 10.3    | httptest      | Simulate server.       | `httptest.NewServer()`   |
| 10.4    | Verify Output | Headers, context, etc. | `assert.Equal(...)`      |

---

| Section | Prod Use          | Feature                   | Tip                     |
| ------- | ----------------- | ------------------------- | ----------------------- |
| 11.1    | Microservices     | Middleware everywhere.    | Log, trace, auth        |
| 11.2    | Observability     | Log/tracing.              | Zap, OpenTelemetry      |
| 11.3    | Distributed Trace | Track across services.    | Inject trace ID         |
| 11.4    | Load Shedding     | Drop requests under load. | Use token bucket        |
| 11.5    | TLS/mTLS          | Secure transport.         | `http.Server.TLSConfig` |

---

| Section | Real World  | Example             | Key Idea                    |
| ------- | ----------- | ------------------- | --------------------------- |
| 12.1    | Gin + JWT   | Auth middleware.    | Secure endpoints            |
| 12.2    | CRUD + Log  | Track + Auth.       | Combine middleware          |
| 12.3    | RBAC        | Role-based.         | Read role from JWT          |
| 12.4    | Docker App  | Container-ready.    | Use Dockerfile + entrypoint |
| 12.5    | Custom Rate | Custom limit logic. | Map\[IP]count + goroutine   |

---

| Section | Best Practices     | Idea               | Summary              |
| ------- | ------------------ | ------------------ | -------------------- |
| 13.1    | Lightweight        | Fast logic.        | Avoid slow DB calls  |
| 13.2    | Avoid Side Effects | No global changes. | Pure functions       |
| 13.3    | Performance Test   | Benchmark chain.   | `go test -bench`     |
| 13.4    | Version Docs       | Clear API + code.  | Markdown docs        |
| 13.5    | Benchmarking       | Test speed.        | Benchmark each layer |

---

