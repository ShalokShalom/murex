package textmanip

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/debug"

	"github.com/lmorg/murex/lang/proc"
)

func init() {
	proc.GoFunctions["match"] = cmdMatch
	proc.GoFunctions["!match"] = cmdMatch
	proc.GoFunctions["regexp"] = cmdRegexp
	proc.GoFunctions["!regexp"] = cmdRegexp
}

func cmdMatch(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.StringAll() == "" {
		return errors.New("No parameters supplied")
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(func(b []byte) {
		matched := bytes.Contains(b, p.Parameters.ByteAll())
		if (matched && !p.IsNot) || (!matched && p.IsNot) {
			err = aw.Write(b)
			if err != nil {
				p.Stderr.Writeln([]byte(err.Error()))
				p.Stdin.ForceClose()
			}
		}
	})

	return aw.Close()
}

func cmdRegexp(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	if p.Parameters.StringAll() == "" {
		return errors.New("No parameters supplied")
	}

	var sRegex []string
	if p.Parameters.Len() == 1 {
		sRegex, err = splitRegexParams(p.Parameters.ByteAll())
		if err != nil {
			return err
		}

	} else {
		// No need to get clever with the regex parser because the parameters are already split by murex's parser
		sRegex = p.Parameters.StringArray()
	}

	if len(sRegex) < 2 {
		return fmt.Errorf("Invalid regexp (too few parameters) in: `%s`", p.Parameters.StringAll())
	}
	if len(sRegex) > 4 {
		return fmt.Errorf("Invalid regexp (too many parameters) in: `%s`", p.Parameters.StringAll())
	}

	var rx *regexp.Regexp
	if rx, err = regexp.Compile(sRegex[1]); err != nil {
		return
	}

	switch sRegex[0][0] {
	case 'm':
		return regexMatch(p, rx, dt)

	case 's':
		return regexSubstitute(p, rx, sRegex, dt)

	case 'f':
		return regexFind(p, rx, dt)

	default:
		return errors.New("Invalid regexp. Please use either match (m), substitute (s) or find (f)")
	}
}

func splitRegexParams(regex []byte) ([]string, error) {
	if len(regex) < 2 {
		return nil, fmt.Errorf("Invalid regexp (too few characters) in: `%s`", string(regex))
	}

	switch regex[1] {
	default:
		return splitRegexDefault(regex)

	case '{':
		return nil, fmt.Errorf("The `{` character is not yet supported for separating regex parameters in: `%s`. (feature in development)", string(regex))
		//return splitRegexBraces(regex)

	case '\\':
		return nil, fmt.Errorf("The `\\` character is not valid for separating regex parameters in: `%s`", string(regex))
	}
}

func splitRegexDefault(regex []byte) (s []string, _ error) {
	var (
		param   []byte
		escaped bool
		token   = regex[1]
	)

	for _, c := range regex {
		switch c {
		default:
			if escaped {
				param = append(param, '\\', c)
				escaped = false
				continue
			}
			param = append(param, c)

		case '\\':
			if escaped {
				param = append(param, '\\', c)
				escaped = false
				continue
			}
			escaped = true

		case token:
			if escaped {
				escaped = false
				param = append(param, c)
				continue
			}

			s = append(s, string(param))
			param = []byte{}
		}
	}
	s = append(s, string(param))

	return
}

var rxCurlyBraceSplit = regexp.MustCompile(`\{(.*?)\}`)

func splitRegexBraces(regex []byte) ([]string, error) {
	s := rxCurlyBraceSplit.FindAllString(string(regex), -1)
	s = append([]string{string(regex[0])}, s...)
	debug.Json("s", s)
	return s, nil
}

// -------- regex functons --------

func regexMatch(p *proc.Process, rx *regexp.Regexp, dt string) error {
	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(func(b []byte) {
		matched := rx.Match(b)
		if (matched && !p.IsNot) || (!matched && p.IsNot) {

			err = aw.Write(b)
			if err != nil {
				p.Stderr.Writeln([]byte(err.Error()))
				p.Stdin.ForceClose()
			}

		}
	})

	return aw.Close()
}

func regexSubstitute(p *proc.Process, rx *regexp.Regexp, sRegex []string, dt string) error {
	if len(sRegex) < 3 {
		return fmt.Errorf("Invalid regex (too few parameters - expecting s/find/substitute/) in: `%s`", p.Parameters.StringAll())
	}

	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	sub := []byte(sRegex[2])

	p.Stdin.ReadArray(func(b []byte) {
		err = aw.Write(rx.ReplaceAll(b, sub))
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			p.Stdin.ForceClose()
		}
	})

	return aw.Close()
}

func regexFind(p *proc.Process, rx *regexp.Regexp, dt string) error {
	aw, err := p.Stdout.WriteArray(dt)
	if err != nil {
		return err
	}

	p.Stdin.ReadArray(func(b []byte) {
		found := rx.FindStringSubmatch(string(b))
		if len(found) > 1 {

			for i := 1; i < len(found); i++ {
				err = aw.WriteString(found[i])
				if err != nil {
					p.Stderr.Writeln([]byte(err.Error()))
					p.Stdin.ForceClose()
				}

			}

		}
	})

	return aw.Close()
}
