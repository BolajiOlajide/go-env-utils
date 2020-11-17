# Go-Env-Utils

A Go (golang) port of the node [env-utils](https://github.com/danethurber/env-utils) project.

Easily Get Environment Variables

## Install

```sh
go get github.com/BolajiOlajide/go-env-utils
```

## Methods

### GetEnvVar

gets an environment variable

```go
import "github.com/BolajiOlajide/go-env-utils"

PORT := env.GetEnvVar("PORT", nil)
```

You can force the value to be a boolean

```go
options := env.Options{
  Boolean: true
}
PORT := env.GetEnvVar("PORT", options)
```

And if your env variable is a comma separated string you can get back an array instead

```go
options := env.Options{
  CommaSeparated: true
}
PORT := env.GetEnvVar("PORT", options)
```

You can also set a fallback value for development mode

```go
options := env.Options{
  DevDefault: "1234"
}
PORT := env.GetEnvVar("PORT", options)
```

