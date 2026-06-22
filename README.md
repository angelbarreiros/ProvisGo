# ProvisGo

Go client for selected PROVIS integration API endpoints using HMAC-256 request signing.

## Install

```bash
go get github.com/angelbarreiros/ProvisGo
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/angelbarreiros/ProvisGo/proviscore"
	"github.com/angelbarreiros/ProvisGo/provisentities"
)

func main() {
	cfg := proviscore.NewConfig(
		"apibase-integraciones.provis.es",
		"your-application-key",
		"your-secret-key",
		false,
	)

	provider := proviscore.Init(cfg)
	defer provider.Close()

	showAll := true
	courses, err := provider.Cursillos("your-installation-id", &provisentities.CursillosParams{
		ShowAllCourses: &showAll,
	})
	if err != nil {
		fmt.Printf("PROVIS error %d: %s\n", err.Code, err.Message)
		return
	}

	fmt.Printf("courses: %d\n", len(courses.Cursillos))
}
```

## Configuration

`proviscore.NewConfig(host, applicationKey, secretKey, debug)` creates the API configuration.

- `host`: API host without scheme, for example `apibase-integraciones.provis.es`.
- `applicationKey`: PROVIS application key.
- `secretKey`: PROVIS secret key used to sign requests.
- `debug`: when `true`, request and response details are written to stdout.

Keep `applicationKey`, `secretKey`, installation IDs, and exported customer data out of source control.

## Supported Endpoints

`proviscore.Init` returns a `*proviscore.ProvisProvider`, which can be used as a named type in your own APIs.

The provider currently exposes:

- `Cursillos`
- `Cuotas`
- `Workers`
- `Personaldata`
- `PersonImage`
- `Families`
- `Installations`
- `Groups`
- `AccessByDate`

Each method returns the typed response and a `*provisentities.ErrorResponse`. Check the error before using the response.



## Debug Output

Debug mode prints the generated curl command, response status, response headers, and response body:

```text
Body: {"example":"response"}
```

This can expose credentials, authorization headers, personal data, and API response contents. Use debug mode only in trusted local environments.

## Release Checklist

Before tagging a new version:

```bash
go test ./...
go vet ./...
git status --short
git tag -a v0.1.4 -m "v0.1.4"
git push origin main
git push origin v0.1.4
```

Make the GitHub repository public only after confirming no private credentials, exports, or non-redistributable documents are tracked.
