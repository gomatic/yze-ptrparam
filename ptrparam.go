// Package ptrparam provides a go/analysis analyzer enforcing the gomatic Go
// immutability standard: function parameters are passed by value, never by
// pointer, unless the pointed-to type is a standard-library type where a pointer
// is the idiomatic calling convention.
package ptrparam

import (
	"go/ast"
	"go/types"

	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// allowedPointerParams are the standard-library types conventionally passed by
// pointer.
var allowedPointerParams = map[string]bool{
	"log/slog.Logger":        true,
	"testing.T":              true,
	"testing.B":              true,
	"testing.F":              true,
	"testing.M":              true,
	"sync.WaitGroup":         true,
	"os.File":                true,
	"net/http.Request":       true,
	"net/http.Response":      true,
	"bytes.Buffer":           true,
	"text/template.Template": true,
	"html/template.Template": true,
	"crypto/tls.Config":      true,
	"database/sql.DB":        true,
	"database/sql.Tx":        true,
	"database/sql.Stmt":      true,
}

// Analyzer reports pointer parameters that are not idiomatic standard-library types.
var Analyzer = &analysis.Analyzer{
	Name:     "ptrparam",
	Doc:      "reports pointer parameters unless the pointed-to type is a standard-library type where a pointer is idiomatic",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

// Registration declares this analyzer to the yze framework.
var Registration = goyze.Registration{
	Name:       "ptrparam",
	Group:      "go",
	Categories: []goyze.Category{"immutability"},
	URL:        "https://docs.gomatic.dev/yze/go/ptrparam",
	Analyzer:   Analyzer,
}

// run reports each disallowed pointer parameter.
func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	insp.Preorder([]ast.Node{(*ast.FuncType)(nil)}, func(n ast.Node) {
		for _, field := range n.(*ast.FuncType).Params.List {
			check(pass, field)
		}
	})
	return nil, nil
}

// check reports a parameter whose type is a non-idiomatic pointer.
func check(pass *analysis.Pass, field *ast.Field) {
	star, ok := field.Type.(*ast.StarExpr)
	if !ok || allowedPointer(pass, star.X) {
		return
	}
	pass.Reportf(star.Pos(), "pointer parameter; pass by value unless it is a standard-library type where a pointer is idiomatic")
}

// allowedPointer reports whether the pointed-to type expression names an
// allow-listed standard-library type.
func allowedPointer(pass *analysis.Pass, x ast.Expr) bool {
	named, ok := pass.TypesInfo.TypeOf(x).(*types.Named)
	if !ok || named.Obj().Pkg() == nil {
		return false
	}
	return allowedPointerParams[named.Obj().Pkg().Path()+"."+named.Obj().Name()]
}
