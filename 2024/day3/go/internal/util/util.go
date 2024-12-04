package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

const (
	_ = iota
	FUNC
	OPERAND
	OPAREN
	CPAREN
	COMMA
	UNK
)

type Token struct {
	Value     string
	TokenType int
}

type Mul struct {
	Operand1 int
	Operand2 int
}

func RunPart1(testMode bool) {
	r, file, err := LoadData(testMode)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	tokens, err := Tokenize(r)

	if err != nil {
		panic(err)
	}

	exprs, err := Parse(tokens)

	if err != nil {
		panic(err)
	}

    s := Eval(exprs)

    fmt.Printf("The answer is: %d\n", s)
}

func Eval(exprs []Mul) int {
    s := 0

    for _, expr := range exprs {
        s += (expr.Operand1 * expr.Operand2)
    }

    return s
}

func LoadData(testMode bool) (*bufio.Reader, *os.File, error) {
	var path string

	if testMode {
		path = "../day3.test"
	} else {
		path = "../day3.input"
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, nil, err
	}

	r := bufio.NewReader(file)

	return r, file, nil
}

func Parse(tokens []Token) ([]Mul, error) {

	o := []Mul{}

	var operand1 int
	var operand2 int

	inExpr := false
	nextExpected := FUNC
	opCount := 0

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.TokenType != nextExpected && inExpr {
			inExpr = false
			nextExpected = FUNC
			opCount = 0
		}

		if token.TokenType != nextExpected {
			continue
		}

		exprComplete := false

		switch nextExpected {
		case FUNC:
			nextExpected = OPAREN
			inExpr = true
		case OPAREN:
			nextExpected = OPERAND
		case OPERAND:
			if opCount == 2 {
				inExpr = false
				nextExpected = FUNC
				opCount = 0
			} else {
				opCount++
				v, err := strconv.ParseInt(token.Value, 10, 64)

				if err != nil {
					return nil, err
				}

				if opCount == 1 {
					operand1 = int(v)
					nextExpected = COMMA
				} else {
					operand2 = int(v)
					nextExpected = CPAREN
				}
			}
		case COMMA:
			nextExpected = OPERAND
		case CPAREN:
			exprComplete = true
		default:
			nextExpected = FUNC
			opCount = 0
			inExpr = false
		}

		if exprComplete {
			o = append(o, Mul{
				Operand1: operand1,
				Operand2: operand2,
			})

			nextExpected = FUNC
			opCount = 0
			inExpr = false
		}
	}

	return o, nil
}

func Tokenize(r *bufio.Reader) ([]Token, error) {
	o := []Token{}

	currentToken := ""
	var currentTokenType int
	var lastRune = 'x'
	for {
		var c rune
		if lastRune != 'x' {
			c = lastRune
			lastRune = 'x'
		} else {
			y, _, err := r.ReadRune()

			if err == io.EOF {
				return o, nil
			} else if err != nil {
				return nil, err
			}

			c = y
		}

		tokenValid := true
		tokenComplete := false

		if currentToken == "" {
			if c == 'm' {
				currentToken = "m"
				currentTokenType = FUNC
			} else if unicode.IsDigit(c) {
				currentToken = string(c)
				currentTokenType = OPERAND
			} else if c == '(' {
				currentToken = "("
				currentTokenType = OPAREN
				tokenComplete = true
			} else if c == ')' {
				currentToken = ")"
				currentTokenType = CPAREN
				tokenComplete = true
			} else if c == ',' {
				currentToken = ","
				currentTokenType = COMMA
				tokenComplete = true
			} else {
				currentToken += string(c)
				currentTokenType = UNK
			}
		} else {
			switch currentTokenType {
			case FUNC:
				if currentToken == "m" && c == 'u' {
					currentToken = "mu"
				} else if currentToken == "mu" && c == 'l' {
					currentToken = "mul"
					tokenComplete = true
				} else {
					tokenValid = false
				}
			case OPERAND:
				if len(currentToken) == 3 {
					tokenValid = false
				} else if unicode.IsDigit(c) {
					currentToken += string(c)

					if len(currentToken) == 3 {
						tokenComplete = true
					}
				} else {
					lastRune = c
					tokenComplete = true
				}
			default:
				tokenValid = false
			}
		}

		if !tokenValid {
			if c == 'm' {
				tokenComplete = true
				lastRune = c
			}

			currentTokenType = UNK
		}

		if tokenComplete {
			o = append(o, Token{
				Value:     currentToken,
				TokenType: currentTokenType,
			})

			tokenComplete = false
			currentToken = ""
		}
	}
}
