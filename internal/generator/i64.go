package generator

import (
	"context"
	"fmt"

	interPlugin "github.com/ccheers/proto-gen-checker/internal/plugin"
	"github.com/ccheers/proto-gen-checker/proto/checker"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type I64 struct {
}

func NewI64() *I64 {
	return &I64{}
}

func (x *I64) Generator(ctx context.Context, file *protogen.GeneratedFile, field *protogen.Field, rule *checker.CheckRule) error {
	i64Rule, ok := rule.Rule.(*checker.CheckRule_I64)
	if !ok {
		return fmt.Errorf("rule type is not i64")
	}
	i64 := i64Rule.I64

	if i64.Lt != nil {
		file.P(fmt.Sprintf("if !(x.%s < %d) {", field.GoName, *i64.Lt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于 %d\")", field.GoName, *i64.Lt))
		file.P("}")
	}
	if i64.Lte != nil {
		file.P(fmt.Sprintf("if !(x.%s <= %d) {", field.GoName, *i64.Lte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于等于 %d\")", field.GoName, *i64.Lte))
		file.P("}")
	}
	if i64.Gt != nil {
		file.P(fmt.Sprintf("if !(x.%s > %d) {", field.GoName, *i64.Gt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于 %d\")", field.GoName, *i64.Gt))
		file.P("}")
	}
	if i64.Gte != nil {
		file.P(fmt.Sprintf("if !(x.%s >= %d) {", field.GoName, *i64.Gte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于等于 %d\")", field.GoName, *i64.Gte))
		file.P("}")
	}
	file.P("")
	return nil
}

func (x *I64) Kind() protoreflect.Kind {
	return protoreflect.Int64Kind
}
