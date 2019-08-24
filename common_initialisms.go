package commoninitialisms

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"sync"

	"golang.org/x/tools/go/packages"

	// import lint package
	_ "golang.org/x/lint"
)

var (
	// ErrNotFound NotFound Error
	ErrNotFound = errors.New("not found")
)

const (
	lintPacakageName         = "golang.org/x/lint"
	lintFileName             = "lint.go"
	commonInitialismsVarName = "commonInitialisms"
)

var (
	once      sync.Once
	cache     map[string]bool
	cachedErr error
)

// Must is panicable methoad
func Must(data map[string]bool, err error) map[string]bool {
	if err != nil {
		panic(err)
	}
	return data
}

// GetCommonInitialisms is get line.commonInitialisms value
func GetCommonInitialisms() (map[string]bool, error) {
	once.Do(func() {
		lintFilePath, err := searchLintFile()
		if err != nil {
			cachedErr = err
			return
		}
		fs := token.NewFileSet()
		parsedFile, err := parser.ParseFile(fs, lintFilePath, nil, 0)
		if err != nil {
			cachedErr = err
			return
		}
		cache, err = getCommonInitialisms(parsedFile)
		if err != nil {
			cachedErr = err
			return
		}
	})
	return cache, cachedErr
}

func searchLintFile() (string, error) {
	pkgs, err := packages.Load(nil, lintPacakageName)
	if err != nil {
		return "", ErrNotFound
	}
	if len(pkgs) == 0 {
		return "", ErrNotFound
	}
	for _, gofile := range pkgs[0].GoFiles {
		if strings.HasSuffix(gofile, fmt.Sprintf("/%s", lintFileName)) {
			return gofile, nil
		}
	}
	return "", ErrNotFound
}

func getCommonInitialisms(astFile *ast.File) (map[string]bool, error) {
	for _, decl := range astFile.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.VAR {
			continue
		}
		for _, spec := range genDecl.Specs {
			varSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			names := varSpec.Names
			if len(names) != 1 || names[0].Name != commonInitialismsVarName {
				continue
			}
			for _, value := range varSpec.Values {
				compositeLit, ok := value.(*ast.CompositeLit)
				if !ok {
					continue
				}
				data := make(map[string]bool)
				for _, itemExpr := range compositeLit.Elts {
					keyValueExpr, ok := itemExpr.(*ast.KeyValueExpr)
					if !ok {
						continue
					}
					basicLit, ok := keyValueExpr.Key.(*ast.BasicLit)
					if !ok {
						continue
					}
					key := strings.Trim(basicLit.Value, "\"")
					data[key] = true
				}
				return data, nil
			}
		}
	}
	return nil, ErrNotFound
}
