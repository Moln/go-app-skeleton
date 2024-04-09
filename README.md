应用基础框架
===========

- di: `go.uber.org/fx`
- cli: `github.com/urfave/cli/v2`
- logger: `github.com/sirupsen/logrus`
- http: `github.com/gin-gonic/gin`

------------------------

目录结构：

```
commands/    # cli
 - config.go # `cli config` Show config.
 - serve.go  # `cli serve` Start HTTP Server
modules/  # Application modules
 - app/   # Default app module
   - controller/ # Controllers of app module 
   - route.go    # Route config
```