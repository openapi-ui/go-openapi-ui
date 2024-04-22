# go-openapi-ui

`go-openapi-ui` is an embedded OpenAPI documentation ui for Go using [OpenAPI-UI](https://github.com/rookie-luochao/openapi-ui) and Go's [1.17+'s embed package](https://golang.org/pkg/embed/), with middleware implementations for: `net/http`, `gin`, `fiber`, and `echo`.

The template is based on the OpenAPI-UI [bundle.js](https://github.com/rookie-luochao/openapi-ui/blob/master/lib/openapi-ui.umd.js) with the script already placed in the html instead of depending on a CDN.

This package does not generate openapi spec file. Check [this example](_examples/gen) for using code generation with swag.

## Usage

```go
import "github.com/rookie-luochao/go-openapi-ui"

...

doc := doc.Doc{
    Title:       "Example API",
    Description: "Example API Description",
    SpecFile:    "./openapi.json", // "./openapi.yaml"
    SpecPath:    "/openapi.json",  // "/openapi.yaml"
    DocsPath:    "/docs",
}
```

- `net/http`

```go
import (
	"net/http"
	"github.com/rookie-luochao/go-openapi-ui"
)

...

http.ListenAndServe(address, doc.Handler())
```

- `gin`

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/rookie-luochao/go-openapi-ui"
	ginopenapiui "github.com/rookie-luochao/go-openapi-ui/gin"
)

...

r := gin.New()
r.Use(ginopenapiui.New(doc))
```

- `echo`

```go
import (
	"github.com/labstack/echo/v4"
	"github.com/rookie-luochao/go-openapi-ui"
	echoopenapiui "github.com/rookie-luochao/go-openapi-ui/echo"
)

...

r := echo.New()
r.Use(echoopenapiui.New(doc))
```

- `fiber`

```go
import (
	"github.com/gofiber/fiber/v2"
	"github.com/rookie-luochao/go-openapi-ui"
	fiberopenapiui "github.com/rookie-luochao/go-openapi-ui/fiber"
)

...

r := fiber.New()
r.Use(fiberopenapiui.New(doc))
```

See [examples](/_examples)

## 致谢

- [go-redoc](https://github.com/mvrilo/go-redoc)

## LICENSE

[MIT](LICENSE)
