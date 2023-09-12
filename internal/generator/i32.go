package generator

import (
	"context"
	"fmt"

	interPlugin "github.com/ccheers/proto-gen-checker/internal/plugin"
	"github.com/ccheers/proto-gen-checker/proto/checker"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type I32 struct {
}

func NewI32() *I32 {
	return &I32{}
}

func (x *I32) Generator(ctx context.Context, file *protogen.GeneratedFile, field *protogen.Field, rule *checker.CheckRule) error {
	i32Rule, ok := rule.Rule.(*checker.CheckRule_I32)
	if !ok {
		return fmt.Errorf("rule type is not i32")
	}
	i32 := i32Rule.I32

	if i32.Lt != nil {
		file.P(fmt.Sprintf("if !(x.%s < %d) {", field.GoName, *i32.Lt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于 %d\")", field.GoName, *i32.Lt))
		file.P("}")
	}
	if i32.Lte != nil {
		file.P(fmt.Sprintf("if !(x.%s <= %d) {", field.GoName, *i32.Lte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于等于 %d\")", field.GoName, *i32.Lte))
		file.P("}")
	}
	if i32.Gt != nil {
		file.P(fmt.Sprintf("if !(x.%s > %d) {", field.GoName, *i32.Gt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于 %d\")", field.GoName, *i32.Gt))
		file.P("}")
	}
	if i32.Gte != nil {
		file.P(fmt.Sprintf("if !(x.%s >= %d) {", field.GoName, *i32.Gte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于等于 %d\")", field.GoName, *i32.Gte))
		file.P("}")
	}
	file.P("")
	return nil
}

func (x *I32) Kind() protoreflect.Kind {
	return protoreflect.Int32Kind
}
