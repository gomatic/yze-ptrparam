package a

import "log/slog"

type Plain struct{ x int }

// takesLocal is flagged: pointer to a local type.
func takesLocal(p *Plain) { _ = p } // want `pointer parameter`

// takesInt is flagged: pointer to a basic type.
func takesInt(n *int) { _ = n } // want `pointer parameter`

// takesErr is flagged: pointer to error (named, no package).
func takesErr(e *error) { _ = e } // want `pointer parameter`

// takesLogger is allowed: a standard-library type where a pointer is idiomatic.
func takesLogger(l *slog.Logger) { _ = l }

// takesValue is fine: a value parameter.
func takesValue(p Plain) { _ = p }
