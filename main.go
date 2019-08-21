package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/soql-parser/parser"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	lineNo, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	charNo, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	fp, err := os.Open(filename)
	defer fp.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReaderSize(fp, 4096)
	body := ""
	for  {
		line, _, err := reader.ReadLine()
		if lineNo == 0 {
			body += string(line)[0:charNo-1]
			break
		}
		if err == io.EOF {
			break
		}
		body += string(line)
		lineNo--
	}
	if err != nil {
		panic(err)
	}
	expectedTokens, _ := Parse(body)
	res, err := json.Marshal(expectedTokens)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func Parse(src string) ([]string, antlr.ParserRuleContext) {
	input := antlr.NewInputStream(src)
	lexer := parser.NewSOQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSOQLParser(stream)
	p.RemoveErrorListeners()
	listener := &ErrorListener{
		parser: p,
	}
	p.AddErrorListener(listener)
	p.BuildParseTrees = true
	tree := p.Soql_query()
	return listener.ExpectedTokens, tree
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	parser antlr.Parser
	ExpectedTokens []string
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//stream := l.parser.GetTokenStream()
	//fmt.Println(l.parser.GetRuleNames()[l.parser.GetParserRuleContext().GetRuleIndex()])
	//fmt.Println(l.parser.GetRuleInvocationStack(l.parser.GetParserRuleContext()))
	verbose := l.parser.GetExpectedTokens().StringVerbose(
		l.parser.GetLiteralNames(),
		l.parser.GetSymbolicNames(),
		false,
	)

	if strings.Contains(verbose, ",") {
		tags := strings.Split(verbose[1:len(verbose)-1], ",")
		expectedTokens := make([]string, len(tags))
		for i, tag := range tags {
			expectedTokens[i] = strings.TrimSpace(tag)
		}
		l.ExpectedTokens = append(l.ExpectedTokens, expectedTokens...)
	} else {
		l.ExpectedTokens = append(l.ExpectedTokens, verbose)
	}
}
