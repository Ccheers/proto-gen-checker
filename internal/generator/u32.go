package generator

import (
	"context"
	"fmt"

	interPlugin "github.com/ccheers/proto-gen-checker/internal/plugin"
	"github.com/ccheers/proto-gen-checker/proto/checker"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type U32 struct {
}

func NewU32() *U32 {
	return &U32{}
}

func (x *U32) Generator(ctx context.Context, file *protogen.GeneratedFile, field *protogen.Field, rule *checker.CheckRule) error {
	u32Rule, ok := rule.Rule.(*checker.CheckRule_U32)
	if !ok {
		return fmt.Errorf("rule type is not u32")
	}
	u32 := u32Rule.U32

	if u32.Lt != nil {
		file.P(fmt.Sprintf("if !(x.%s < %d) {", field.GoName, *u32.Lt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于 %d\")", field.GoName, *u32.Lt))
		file.P("}")
	}
	if u32.Lte != nil {
		file.P(fmt.Sprintf("if !(x.%s <= %d) {", field.GoName, *u32.Lte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于等于 %d\")", field.GoName, *u32.Lte))
		file.P("}")
	}
	if u32.Gt != nil {
		file.P(fmt.Sprintf("if !(x.%s > %d) {", field.GoName, *u32.Gt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于 %d\")", field.GoName, *u32.Gt))
		file.P("}")
	}
	if u32.Gte != nil {
		file.P(fmt.Sprintf("if !(x.%s >= %d) {", field.GoName, *u32.Gte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于等于 %d\")", field.GoName, *u32.Gte))
		file.P("}")
	}
	file.P("")
	return nil
}

func (x *U32) Kind() protoreflect.Kind {
	return protoreflect.Uint32Kind
}
