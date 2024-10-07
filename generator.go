package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stoewer/go-strcase"
)

//go:embed templates/*.tmpl
var tmplFS embed.FS

var tmpl = template.Must(template.New("contract").
	Funcs(template.FuncMap{
		"PrivateVarName": func(name string) string {
			return strcase.LowerCamelCase(name)
		},
		"PublicVarName": func(name string) string {
			return strcase.UpperCamelCase(name)
		},
		"String": func(s fmt.Stringer) string {
			return s.String()
		},
	}).ParseFS(tmplFS, "templates/*.tmpl"))

func GenerateCodeForJSON(pkg, name, address string, jsonABI string, out io.Writer) error {
	contractABI, err := abi.JSON(strings.NewReader(string(jsonABI)))
	if err != nil {
		return fmt.Errorf("failed to decode abi from json: %w", err)
	}

	methods := make(map[string][]abi.Method)
	for methodName, method := range contractABI.Methods {
		stored := methods[methodName]
		methods[methodName] = append(stored, method)
	}

	contract := Contract{
		Version: version,
		Package: pkg,
		Name:    name,
		Address: address,
		JSONABI: jsonABI,
		Types:   make(map[string][]reflect.StructField),
	}
	for name, method := range methods {
		if len(method) != 1 {
			panic("not implemented: multiple variants for the same method name")
		}
		err = collectMethod(&contract, method[0])
		if err != nil {
			return fmt.Errorf("failed to collect method %s: %w", name, err)
		}
	}
	sort.Slice(contract.Method, func(i, j int) bool {
		return contract.Method[i].Name < contract.Method[j].Name
	})
	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "contract.tmpl", contract)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	contents, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return fmt.Errorf("failed to format content: %w", err)
	}
	_, err = fmt.Fprintln(out, string(contents))
	return err
}

func collectMethod(contract *Contract, method abi.Method) error {
	if !isFunction(method) {
		return nil
	}
	typ, inputs := asArguments(method.Inputs)
	for i := range typ {
		contract.Types[typ[i].Name] = typ[i].Fields
	}
	typ, outputs := asArguments(method.Outputs)
	for i := range typ {
		contract.Types[typ[i].Name] = typ[i].Fields
	}

	contract.Method = append(contract.Method, Method{
		Contract: contract.Name,
		Name:     method.Name,
		Inputs:   inputs,
		Outputs:  outputs,
	})
	return nil
}

type Contract struct {
	Version string
	Package string
	Name    string
	Address string
	JSONABI string
	Method  []Method
	Types   map[string][]reflect.StructField
}

type Type struct {
	Name   string
	Fields []reflect.StructField
}

type Method struct {
	Contract string
	Name     string
	Inputs   Arguments
	Outputs  Arguments
}

type Argument struct {
	Name       string
	Type       string
	CustomType bool
	IsSlice    bool
}

type Arguments []Argument

func isOfCustomType(typ reflect.Type) bool {
	typ = asTypeInstance(typ)
	if typ.Kind() != reflect.Struct {
		return false
	}
	return typ.Name() == ""
}

func typPrefixes(typ reflect.Type) string {
	if typ.Kind() == reflect.Ptr {
		return "*" + typPrefixes(typ.Elem())
	}
	if typ.Kind() == reflect.Slice {
		return "[]" + typPrefixes(typ.Elem())
	}
	if typ.Kind() == reflect.Array {
		return "[" + strconv.Itoa(typ.Len()) + "]" + typPrefixes(typ.Elem())
	}
	return ""
}

func asTypeInstance(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Ptr {
		return asTypeInstance(typ.Elem())
	}
	if typ.Kind() == reflect.Slice {
		return asTypeInstance(typ.Elem())
	}
	if typ.Kind() == reflect.Array {
		return asTypeInstance(typ.Elem())
	}
	return typ
}

func asStructFields(typ reflect.Type) (res []reflect.StructField) {
	typ = asTypeInstance(typ)
	res = make([]reflect.StructField, typ.NumField())
	for i := range res {
		res[i] = typ.Field(i)
	}
	return res
}

func asArguments(abiArgs abi.Arguments) ([]Type, Arguments) {
	var args Arguments
	var types []Type
	for _, arg := range abiArgs {
		argType := arg.Type.GetType().String()
		customType := false
		isSlice := false

		if typ := arg.Type.GetType(); isOfCustomType(typ) {
			typName := strcase.UpperCamelCase(arg.Name)
			if typName == "" {
				typName = arg.Type.TupleRawName
				if typName == "" {
					typName = arg.Type.Elem.TupleRawName
					isSlice = true
				}
			}
			argType = typPrefixes(typ) + typName

			ctype := Type{
				Name:   typName,
				Fields: asStructFields(typ),
			}
			types = append(types, ctype)
			customType = true
		}
		args = append(args, Argument{
			Name:       arg.Name,
			Type:       argType,
			CustomType: customType,
			IsSlice:    isSlice,
		})
	}
	return types, args
}

func isFunction(method abi.Method) bool {
	return method.Type == abi.Function
}
