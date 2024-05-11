# go-openapi-ui

`go-openapi-ui` is an embedded OpenAPI documentation ui for Go using [OpenAPI-UI](https://github.com/rookie-luochao/openapi-ui) and Go's [1.17+'s embed package](https://golang.org/pkg/embed/), with middleware implementations for: `net/http`, `gin`, `fiber`, and `echo`.

The template is based on the OpenAPI-UI [bundle.js](https://github.com/rookie-luochao/openapi-ui/blob/master/lib/openapi-ui.umd.js) with the script already placed in the html instead of depending on a CDN.

This package does not generate openapi spec file. Check [this example](https://github.com/swaggo/swag/tree/master/example) for using code generation with swag.

## Usage

```go
import "github.com/openapi-ui/go-openapi-ui/pkg/doc"

...

doc := doc.Doc{
    Title:       "Example API",
    Description: "Example API Description",
    SpecFile:    "./openapi.json", // "./openapi.yaml"
    SpecPath:    "/openapi.json",  // "/openapi.yaml"
    DocsPath:    "/docs",
    Theme:       "light", // default is light, support light or dark
}
```

- `net/http`

```go
import (
	"net/http"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
)

...

http.ListenAndServe(address, doc.Handler())
```

- `gin`

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
	ginopenapiui "github.com/openapi-ui/go-openapi-ui/pkg/middleware/gin"
)

...

r := gin.New()
r.Use(ginopenapiui.New(doc))
```

- `echo`

```go
import (
	"github.com/labstack/echo/v4"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
	echoopenapiui "github.com/openapi-ui/go-openapi-ui/pkg/middleware/echo"
)

...

r := echo.New()
r.Use(echoopenapiui.New(doc))
```

- `fiber`

```go
import (
	"github.com/gofiber/fiber/v2"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
	fiberopenapiui "github.com/openapi-ui/go-openapi-ui/pkg/middleware/fiber"
)

...

r := fiber.New()
r.Use(fiberopenapiui.New(doc))
```

See [examples](/_examples)

## Thanks

- [go-redoc](https://github.com/mvrilo/go-redoc)

## LICENSE

[MIT](LICENSE)
