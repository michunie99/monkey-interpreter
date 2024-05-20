package parser

import (
	"fmt"
	"strings"
)

var indentLevel = 0

const identPlaceHolder string = "\t"

func printIndent() string {
	return strings.Repeat(identPlaceHolder, indentLevel-1)
}

func printTrace(tr string) {
	fmt.Printf("%s%s\n", printIndent(), tr)
}

func incIndent() { indentLevel = indentLevel + 1 }
func decIndent() { indentLevel = indentLevel - 1 }

func trace(msg string) string {
	incIndent()
	printTrace("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	printTrace("END " + msg)
	decIndent()
}
