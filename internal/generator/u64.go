package generator

import (
	"context"
	"fmt"

	interPlugin "github.com/ccheers/proto-gen-checker/internal/plugin"
	"github.com/ccheers/proto-gen-checker/proto/checker"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type U64 struct {
}

func NewU64() *U64 {
	return &U64{}
}

func (x *U64) Generator(ctx context.Context, file *protogen.GeneratedFile, field *protogen.Field, rule *checker.CheckRule) error {
	u64Rule, ok := rule.Rule.(*checker.CheckRule_U64)
	if !ok {
		return fmt.Errorf("rule type is not u64")
	}
	u64 := u64Rule.U64

	if u64.Lt != nil {
		file.P(fmt.Sprintf("if !(x.%s < %d) {", field.GoName, *u64.Lt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于 %d\")", field.GoName, *u64.Lt))
		file.P("}")
	}
	if u64.Lte != nil {
		file.P(fmt.Sprintf("if !(x.%s <= %d) {", field.GoName, *u64.Lte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须小于等于 %d\")", field.GoName, *u64.Lte))
		file.P("}")
	}
	if u64.Gt != nil {
		file.P(fmt.Sprintf("if !(x.%s > %d) {", field.GoName, *u64.Gt))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于 %d\")", field.GoName, *u64.Gt))
		file.P("}")
	}
	if u64.Gte != nil {
		file.P(fmt.Sprintf("if !(x.%s >= %d) {", field.GoName, *u64.Gte))
		interPlugin.GenErrorMsg(file, rule, fmt.Sprintf("(\"%s 必须大于等于 %d\")", field.GoName, *u64.Gte))
		file.P("}")
	}
	file.P("")
	return nil
}

func (x *U64) Kind() protoreflect.Kind {
	return protoreflect.Uint64Kind
}
