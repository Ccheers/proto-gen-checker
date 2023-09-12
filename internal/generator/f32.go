package generator

import (
	"context"
	"fmt"

	interPlugin "github.com/ccheers/proto-gen-checker/internal/plugin"
	"github.com/ccheers/proto-gen-checker/proto/checker"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type F32 struct {
}

func NewF32() *F32 {
	return &F32{}
}

func (x *F32) Generator(ctx context.Context, file *protogen.GeneratedFile, field *protogen.Field, rule *checker.CheckRule) error {
	f32Rule, ok := rule.Rule.(*checker.CheckRule_Float)
	if !ok {
		return fmt.Errorf("rule type is not f32")
	}
	f32 := f32Rule.Float

	if f32.Lt != nil {
		file.P(fmt.Sprintf("if !(x.%s < %f) {", field.GoName, *f32.Lt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于 %f\")", field.GoName, *f32.Lt))
		file.P("}")
	}
	if f32.Lte != nil {
		file.P(fmt.Sprintf("if !(x.%s <= %f) {", field.GoName, *f32.Lte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于等于 %f\")", field.GoName, *f32.Lte))
		file.P("}")
	}
	if f32.Gt != nil {
		file.P(fmt.Sprintf("if !(x.%s > %f) {", field.GoName, *f32.Gt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于 %f\")", field.GoName, *f32.Gt))
		file.P("}")
	}
	if f32.Gte != nil {
		file.P(fmt.Sprintf("if !(x.%s >= %f) {", field.GoName, *f32.Gte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于等于 %f\")", field.GoName, *f32.Gte))
		file.P("}")
	}
	file.P("")
	return nil
}

func (x *F32) Kind() protoreflect.Kind {
	return protoreflect.FloatKind
}
