package lang

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc/parameters"
)

func genEmptyParamTokens() (pt [][]parameters.ParamToken) {
	pt = make([][]parameters.ParamToken, 1)
	pt[0] = make([]parameters.ParamToken, 1)
	return
}

func parseBlock(block []rune) (nodes Nodes, pErr ParserError) {
	defer debug.Json("Parser", nodes)
	var (
		// Current state
		commentLine              bool
		escaped                  bool
		quoteSingle, quoteDouble bool
		braceCount               int
		ignoreWhitespace         bool = true
		scanFuncName             bool = true
		//newLine                  bool

		// Parsed thus far
		node   Node                   = Node{NewChain: true, ParamTokens: genEmptyParamTokens()}
		pop    *string                = &node.Name
		pCount int                    // parameter count
		pToken *parameters.ParamToken = &node.ParamTokens[0][0]
		//pToken parameters.ParamToken
	)
	defer debug.Json("Last node", node)

	startParameters := func() {
		scanFuncName = false
		node.Parameters = make([]string, 1)
		pop = &node.Parameters[0]
		node.ParamTokens = genEmptyParamTokens()
		pCount = 0
		pToken = &node.ParamTokens[pCount][0]
	}

	appendNode := func() {
		if len(node.Parameters) > 1 && len(node.Parameters[len(node.Parameters)-1]) == 0 {
			node.Parameters = node.Parameters[:len(node.Parameters)-1]
			//TODO: crop if node.ParamToken[][].Type == 0 ???
		}

		if node.Name != "" {
			nodes = append(nodes, node)
		}

		ignoreWhitespace = true
	}

	for i, b := range block {
		if commentLine {
			if b == '\n' {
				commentLine = false
			}
			continue
		}

		if pToken.Type != 0 {
			switch {
			case b == '-' ||
				('a' < b && b < 'z') ||
				('A' < b && b < 'Z') ||
				('0' < b && b < '9'):
				pToken.Key += string(b)
				*pop += string(b)
				continue

			case b == '{' && (*pop)[pToken.StrLoc-1] != '$':
				pToken.Type = parameters.TokenTypeBlockString
				*pop += string(b)
				continue

			case b == '{' && (*pop)[pToken.StrLoc-1] != '@':
				pToken.Type = parameters.TokenTypeBlockArray
				*pop += string(b)
				continue

			default:
				if len(pToken.Key) > 0 {
					// TODO: need to get this bit working
					node.ParamTokens[pCount] = append(node.ParamTokens[pCount], parameters.ParamToken{})
					pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				} else {
					pToken.Type = 0
				}
			}

		}

		switch b {
		case '#':
			switch {
			case escaped, quoteSingle, quoteDouble, braceCount > 0:
				*pop += string(b)
			default:
				commentLine = true
			}

		case '\\':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case escaped:
				*pop += string(b)
				escaped = false
			default:
				escaped = true
			}

		case '\'':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle:
				quoteSingle = false
			case quoteDouble:
				*pop += string(b)
			default:
				quoteSingle = true
			}

		case '"':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle:
				*pop += string(b)
			case quoteDouble:
				quoteDouble = false
			default:
				quoteDouble = true
			}

		case ':':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case scanFuncName:
				startParameters()
			default:
				*pop += string(b)
			}

		case '{':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case scanFuncName:
				startParameters()
				*pop += string(b)
				braceCount++
			default:
				*pop += string(b)
				braceCount++
			}

		case '}':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case scanFuncName:
				pErr = raiseErr(ErrUnexpectedCloseBrace, i)
				return
			case braceCount == 0:
				pErr = raiseErr(ErrClosingBraceNoOpen, i)
				return
			default:
				*pop += string(b)
				braceCount--
			}

		case ' ', '\t', '\r':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case !scanFuncName:
				if len(*pop) > 0 {
					node.Parameters = append(node.Parameters, "")
					pop = &node.Parameters[len(node.Parameters)-1]
					pCount++
					node.ParamTokens = append(node.ParamTokens, make([]parameters.ParamToken, 1))
					pToken = &node.ParamTokens[pCount][len(node.ParamTokens[pCount])-1]
				}
			case scanFuncName && !ignoreWhitespace:
				startParameters()

			default:
				// do nothing
			}

		case '\n':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case !scanFuncName:
				appendNode()
				node = Node{NewChain: true}
				pop = &node.Name
				scanFuncName = true
				//newLine = true
			case scanFuncName && !ignoreWhitespace:
				startParameters()
			default:
				// do nothing
			}

		case '|':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeToken, i)
				return
			/*case newLine:
			newLine = false
			node.NewChain = false
			if len(nodes) > 0 {
				nodes.Last().PipeOut = true
			}*/
			default:
				node.PipeOut = true
				appendNode()
				node = Node{}
				pop = &node.Name
				scanFuncName = true
			}

		case '?':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case len(node.Name) == 0:
				pErr = raiseErr(ErrUnexpectedPipeToken, i)
				return
			/*case newLine:
			newLine = false
			node.NewChain = false
			if len(nodes) > 0 {
				nodes.Last().PipeErr = true
			}*/
			default:
				node.PipeErr = true
				appendNode()
				node = Node{}
				pop = &node.Name
				scanFuncName = true
			}

		case '>':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
			case len(*pop) > 0 && (*pop)[len(*pop)-1] == '-':
				/*if len(node.Name) == 0 {
					pErr = raiseErr(ErrUnexpectedPipeToken, i)
					return
				}*/
				*pop = (*pop)[:len(*pop)-1]
				node.PipeOut = true
				appendNode()
				node = Node{Method: true}
				pop = &node.Name
				scanFuncName = true

				/*if newLine {
					node.NewChain = false
					node.Method = true
					nodes.Last().PipeOut = true
					newLine = false
				}*/
			default:
				*pop += string(b)
			}

		case ';':
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			case quoteSingle, quoteDouble:
				*pop += string(b)
			case braceCount > 0:
				*pop += string(b)
				//case !scanFuncName:
			default:
				appendNode()
				node = Node{NewChain: true}
				pop = &node.Name
				scanFuncName = true
				//default:
				// do nothing
			}

		case '$':
			if !scanFuncName && braceCount == 0 && !quoteSingle && !escaped {
				pToken.Type = parameters.TokenTypeString
				pToken.StrLoc = len(*pop)
			}
			*pop += string(b)

		case '@':
			if !scanFuncName && braceCount == 0 && !quoteSingle && !escaped {
				pToken.Type = parameters.TokenTypeArray
				pToken.StrLoc = len(*pop)
			}
			*pop += string(b)

		case 's':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += " "
				escaped = false
			default:
				*pop += string(b)
			}

		case 't':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += "\t"
				escaped = false
			default:
				*pop += string(b)
			}

		case 'r':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += "\r"
				escaped = false
			default:
				*pop += string(b)
			}

		case 'n':
			switch {
			case braceCount > 0:
				*pop += string(b)
			case quoteSingle:
				*pop += string(b)
			case escaped:
				*pop += "\n"
				escaped = false
			default:
				*pop += string(b)
			}

		default:
			switch {
			case escaped:
				*pop += string(b)
				escaped = false
			default:
				ignoreWhitespace = false
				*pop += string(b)
				/*if b != '-' {
					newLine = false
				}*/
			}
		}
	}

	switch {
	case escaped:
		return nil, raiseErr(ErrUnterminatedEscape, 0)
	case quoteSingle:
		return nil, raiseErr(ErrUnterminatedQuotesSingle, 0)
	case quoteDouble:
		return nil, raiseErr(ErrUnterminatedQuotesDouble, 0)
	case braceCount > 0:
		return nil, raiseErr(ErrUnterminatedBrace, 0)
	}

	appendNode()

	return
}
