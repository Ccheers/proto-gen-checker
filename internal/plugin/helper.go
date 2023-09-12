package plugin

import (
	"fmt"
	"strconv"

	"github.com/ccheers/proto-gen-checker/proto/checker"
	"google.golang.org/protobuf/compiler/protogen"
)

func GenErrorMsg(file *protogen.GeneratedFile, rule *checker.CheckRule, msg string) {
	if rule.ErrorMsg != "" {
		file.P("return ", FmtPkg.Ident("Errorf"), fmt.Sprintf("(%s)", strconv.Quote(rule.ErrorMsg)))
		return
	}
	file.P("return ", FmtPkg.Ident("Errorf"), fmt.Sprintf("(%s)", strconv.Quote(msg)))
}
