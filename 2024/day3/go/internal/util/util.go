package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	_ = iota
	MUL
	DO
	DONT
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

func RunPart2(testMode bool) {
	r, file, err := LoadData(testMode, false)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	tokens, err := Tokenize(r)

	if err != nil {
		panic(err)
	}

	exprs, err := Parse(tokens, true)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The answer is: %d\n", Eval(exprs))

}

func RunPart1(testMode bool) {
	r, file, err := LoadData(testMode, true)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	tokens, err := Tokenize(r)

	if err != nil {
		panic(err)
	}

	exprs, err := Parse(tokens, false)

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

func LoadData(testMode bool, part1 bool) (*bufio.Scanner, *os.File, error) {
	var path string

	if testMode {
		if part1 {
			path = "../day3.1.test"
		} else {
			path = "../day3.2.test"
		}
	} else {
		path = "../day3.input"
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, nil, err
	}

	r := bufio.NewScanner(file)

	return r, file, nil
}

func expect(tokenType int) map[int]struct{} {
	return map[int]struct{}{tokenType: {}}
}

func Parse(tokens []Token, honorIgnore bool) ([]Mul, error) {
	o := []Mul{}

	var operand1 int
	var operand2 int

	inExpr := false

	start := map[int]struct{}{
		DO:   {},
		DONT: {},
		MUL:  {},
	}

	nextExpected := start
	opCount := 0
	currTyp := UNK
	ignoring := false

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		_, ok := nextExpected[token.TokenType]

		if !ok && inExpr {
			inExpr = false
			nextExpected = start
			opCount = 0
		}

		if !ok {
			continue
		}

		exprComplete := false

		switch token.TokenType {
		case DO:
			currTyp = DO
			nextExpected = expect(OPAREN)
			inExpr = true
		case DONT:
			currTyp = DONT
			nextExpected = expect(OPAREN)
			inExpr = true
		case MUL:
			currTyp = MUL
			nextExpected = expect(OPAREN)
			inExpr = true
		case OPAREN:
			if currTyp == DO || currTyp == DONT {
				nextExpected = expect(CPAREN)
			} else {
				nextExpected = expect(OPERAND)
			}
		case OPERAND:
			if opCount == 2 {
				inExpr = false
				nextExpected = start
				opCount = 0
			} else {
				opCount++
				v, err := strconv.ParseInt(token.Value, 10, 64)

				if err != nil {
					return nil, err
				}

				if opCount == 1 {
					operand1 = int(v)
					nextExpected = expect(COMMA)
				} else {
					operand2 = int(v)
					nextExpected = expect(CPAREN)
				}
			}
		case COMMA:
			nextExpected = expect(OPERAND)
		case CPAREN:
			exprComplete = true
			if honorIgnore {
				if currTyp == DO {
					ignoring = false
				} else if currTyp == DONT {
					ignoring = true
				}

			}
		default:
			nextExpected = start
			opCount = 0
			inExpr = false
		}

		if exprComplete {
			if !ignoring && currTyp == MUL {
				o = append(o, Mul{
					Operand1: operand1,
					Operand2: operand2,
				})

			}

			nextExpected = start
			opCount = 0
			inExpr = false
		}
	}

	return o, nil
}

func Tokenize(scanner *bufio.Scanner) ([]Token, error) {
	q := &Queue{}
	for scanner.Scan() {
		runes := []rune(scanner.Text())

		for _, r := range runes {
			q.Enqueue(r)
		}
	}

	tokens := []Token{}

	currTok := &StrBuild{}
	currTyp := UNK
	var expect rune

	for {
		r, err := q.Dequeue()

		if err != nil {
			break
		}

		done := false
		invalid := false

		if currTok.IsEmpty() {
			currTok.Append(r)

			if unicode.IsDigit(r) {
				currTyp = OPERAND
			} else {
				switch r {
				case 'm':
					currTyp = MUL
					expect = 'u'
				case 'd':
					currTyp = DO
					expect = 'o'
				case '(':
					currTyp = OPAREN
					done = true
				case ')':
					currTyp = CPAREN
					done = true
				case ',':
					currTyp = COMMA
					done = true
				default:
					currTyp = UNK
				}

			}
		} else {
			switch currTyp {
			case OPERAND:
				if !unicode.IsDigit(r) || len(currTok.String()) == 3 {
					done = true
					q.Push(r)
				} else {
					currTok.Append(r)
				}
			case MUL:
				if r != expect {
					invalid = true
				} else {
					currTok.Append(r)
					if expect == 'u' {
						expect = 'l'
					} else {
						done = true
					}
				}
			case DO:
				if r != expect {
					invalid = true
				} else {
					currTok.Append(r)
					expect = 'n'
					currTyp = DONT
				}
			case DONT:
				if r != expect {
					if currTok.String() == "do" {
						currTyp = DO
						done = true
						q.Push(r)
					} else {
						invalid = true
					}
				} else {
					currTok.Append(r)

					if expect == 'n' {
						expect = '\''
					} else if expect == '\'' {
						expect = 't'
					} else {
						done = true
					}
				}
			default:
				invalid = true
			}
		}

		if invalid {
			starts := map[rune]struct{}{'m': {}, 'd': {}, '(': {}, ')': {}, ',': {}}

			_, ok := starts[r]

			if ok || unicode.IsDigit(r) {
				q.Push(r)
				done = true
			} else {
				currTok.Append(r)
			}

			currTyp = UNK
		}

		if done {
			tokens = append(tokens, Token{
				Value:     currTok.Flush(),
				TokenType: currTyp,
			})
		}

	}

	return tokens, nil
}
