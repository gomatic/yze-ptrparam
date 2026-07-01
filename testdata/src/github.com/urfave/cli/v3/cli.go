// Package cli is a minimal stand-in for github.com/urfave/cli/v3, the
// sanctioned CLI framework whose *Command callback parameters are allowed.
package cli

// Command mirrors the framework's command type.
type Command struct{ Name string }
