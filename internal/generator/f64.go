package generator

import (
	"context"
	"fmt"

	interPlugin "github.com/ccheers/proto-gen-checker/internal/plugin"
	"github.com/ccheers/proto-gen-checker/proto/checker"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type F64 struct {
}

func NewF64() *F64 {
	return &F64{}
}

func (x *F64) Generator(ctx context.Context, file *protogen.GeneratedFile, field *protogen.Field, rule *checker.CheckRule) error {
	f64Rule, ok := rule.Rule.(*checker.CheckRule_Double)
	if !ok {
		return fmt.Errorf("rule type is not f64")
	}
	f64 := f64Rule.Double

	if f64.Lt != nil {
		file.P(fmt.Sprintf("if !(x.%s < %f) {", field.GoName, *f64.Lt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于 %f\")", field.GoName, *f64.Lt))
		file.P("}")
	}
	if f64.Lte != nil {
		file.P(fmt.Sprintf("if !(x.%s <= %f) {", field.GoName, *f64.Lte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于等于 %f\")", field.GoName, *f64.Lte))
		file.P("}")
	}
	if f64.Gt != nil {
		file.P(fmt.Sprintf("if !(x.%s > %f) {", field.GoName, *f64.Gt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于 %f\")", field.GoName, *f64.Gt))
		file.P("}")
	}
	if f64.Gte != nil {
		file.P(fmt.Sprintf("if !(x.%s >= %f) {", field.GoName, *f64.Gte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于等于 %f\")", field.GoName, *f64.Gte))
		file.P("}")
	}
	file.P("")
	return nil
}

func (x *F64) Kind() protoreflect.Kind {
	return protoreflect.DoubleKind
}
