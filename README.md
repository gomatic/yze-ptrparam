# yze-go-ptrparam

A [`yze`](https://github.com/gomatic/yze) analyzer (group `go`, category `immutability`) enforcing the gomatic Go immutability standard: function parameters are passed **by value**, never by pointer, unless the pointed-to type is a standard-library type where a pointer is the idiomatic calling convention (`*slog.Logger`, `*testing.T`, `*http.Request`, `*sql.DB`, …).

- **Rule:** `yze/go/ptrparam`
- **Binary:** `cmd/yze-go-ptrparam` runs it standalone (`text`/`-json`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework. The allow-list is baked in for v1; a configurable `-allow` flag is planned.
