package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"multichecker/errcheck"
)

func main() {
	multichecker.Main(
		errcheck.ErrCheckAnalyzer,
		errcheck.GoErrCheckAnalyzer,
		errcheck.DeferErrCheckAnalyzer,
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		shift.Analyzer,
	)
}
