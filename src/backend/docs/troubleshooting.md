# Gin

Gin panic when inserting/reading from table. Example error:
```go
 2024/12/30 15:10:30 http: panic serving 172.18.0.1:37624: reflect: reflect.Value.Set using unaddressable value
 goroutine 56 [running]:
 net/http.(*conn).serve.func1()
   /usr/local/go/src/net/http/server.go:1947 +0xbe
 panic({0xf949a0?, 0xc0001256b0?})
   /usr/local/go/src/runtime/panic.go:785 +0x132
 go.opentelemetry.io/otel/sdk/trace.(*recordingSpan).End.deferwrap1()
   /go/pkg/mod/go.opentelemetry.io/otel/sdk@v1.33.0/trace/span.go:467 +0x25
 go.opentelemetry.io/otel/sdk/trace.(*recordingSpan).End(0xc000170b40, {0x0, 0x0, 0xc000228060?})
   /go/pkg/mod/go.opentelemetry.io/otel/sdk@v1.33.0/trace/span.go:506 +0xb7b
 panic({0xf949a0?, 0xc0001256b0?})
   /usr/local/go/src/runtime/panic.go:785 +0x132
 reflect.flag.mustBeAssignableSlow(0xc00043c9c0?)
   /usr/local/go/src/reflect/value.go:257 +0x74
 reflect.flag.mustBeAssignable(...)
   /usr/local/go/src/reflect/value.go:244
 reflect.Value.Set({0x116f700?, 0xc00043c9c8?, 0x1049520?}, {0x116f700?, 0xc0004869d8?, 0x116f700?})
   /usr/local/go/src/reflect/value.go:2307 +0x65
 gorm.io/gorm/schema.(*Field).setupValuerAndSetter.func12({0x134e880?, 0x1b9a060?}, {0x1049520?, 0xc00043c9c0?, 0xc0002e8320?}, {0x116f700, 0xc0004869d8})
   /go/pkg/mod/gorm.io/gorm@v1.25.12/schema/field.go:840 +0x405
 gorm.io/gorm/callbacks.ConvertToCreateValues(0xc0004808c0)
   /go/pkg/mod/gorm.io/gorm@v1.25.12/callbacks/create.go:326 +0x17cb
 gorm.io/gorm/callbacks.RegisterDefaultCallbacks.Create.func3(0xc0002e5e90)
   /go/pkg/mod/gorm.io/gorm@v1.25.12/callbacks/create.go:66 +0x137
 gorm.io/gorm.(*processor).Execute(0xc000287950, 0xc000355410?)
   /go/pkg/mod/gorm.io/gorm@v1.25.12/callbacks.go:130 +0x3cb
 gorm.io/gorm.(*DB).Create(0x1049520?, {0x1049520, 0xc00043c9c0})
   /go/pkg/mod/gorm.io/gorm@v1.25.12/finisher_api.go:24 +0xa8
 github.com/Sourceware-Lab/realquick/api/timeblock.Post({0xc00013b848?, 0x1357040?}, 0xc00043c8f0)
   /app/api/timeblock/handler.go:27 +0x59
 github.com/danielgtaylor/huma/v2.Register[...].func1()
   /go/pkg/mod/github.com/danielgtaylor/huma/v2@v2.27.0/huma.go:1451 +0x733
 github.com/danielgtaylor/huma/v2/adapters/humagin.(*ginAdapter).Handle.func1(0xc000176800)
   /go/pkg/mod/github.com/danielgtaylor/huma/v2@v2.27.0/adapters/humagin/humagin.go:137 +0x77
 github.com/gin-gonic/gin.(*Context).Next(...)
   /go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185
 github.com/gin-contrib/logger.SetLogger.func1(0xc000176800)
   /go/pkg/mod/github.com/gin-contrib/logger@v1.2.3/logger.go:145 +0x611
 github.com/gin-gonic/gin.(*Context).Next(...)
   /go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185
 go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin.Middleware.func1(0xc000176800)
   /go/pkg/mod/go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin@v0.58.0/gintrace.go:90 +0x799
 github.com/gin-gonic/gin.(*Context).Next(...)
   /go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185
 github.com/gin-gonic/gin.(*Engine).handleHTTPRequest(0xc0000fe340, 0xc000176800)
   /go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633 +0x892
 github.com/gin-gonic/gin.(*Engine).ServeHTTP(0xc0000fe340, {0x134b8e0, 0xc0001b42a0}, 0xc000240280)
   /go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589 +0x1b2
 net/http.serverHandler.ServeHTTP({0x1346028?}, {0x134b8e0?, 0xc0001b42a0?}, 0x6?)
   /usr/local/go/src/net/http/server.go:3210 +0x8e
 net/http.(*conn).serve(0xc00015b4d0, {0x134ea38, 0xc00012a570})
   /usr/local/go/src/net/http/server.go:2092 +0x5d0
 created by net/http.(*Server).Serve in goroutine 50
   /usr/local/go/src/net/http/server.go:3360 +0x485
```
This is due to the struct being passed to the not being a pointer.

Incorrect:
```go
result := dbpg.DB.Create(input.Body)

result := dbpg.DB.First(resp.Body, input.ID)
```

Correct:
```go
result := dbpg.DB.Create(&input.Body)

result := dbpg.DB.First(&resp.Body, input.ID)
```
