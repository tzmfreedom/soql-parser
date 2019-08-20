package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/soql-parser/parser"
)

func Parse(src string) {
	input := antlr.NewInputStream(src)
	lexer := parser.NewSOQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSOQLParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//p.RemoveErrorListeners()
	//p.AddErrorListener(&ErrorListener{
	//	parser: p,
	//})
	p.BuildParseTrees = true
	tree := p.Soql_query()
	t := tree.Accept(&Builder{
		Source: src,
	})
	return t.(Node)
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	parser antlr.Parser
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	stream := l.parser.GetTokenStream()
	fmt.Println(l.parser.GetRuleNames()[l.parser.GetParserRuleContext().GetRuleIndex()])
	fmt.Println(l.parser.GetRuleInvocationStack(l.parser.GetParserRuleContext()))
	fmt.Println(l.parser.GetExpectedTokens().StringVerbose(
		l.parser.GetLiteralNames(),
		l.parser.GetSymbolicNames(),
		false,
	))
	for i := 0; i < 4; i++ {
		token := stream.Get(stream.Index() - 1 - i)
		// last token is always "fake" EOF token
		tokenType := recognizer.GetSymbolicNames()[token.GetTokenType()]

		fmt.Println(tokenType)
		fmt.Println(token.GetText())
	}
}
