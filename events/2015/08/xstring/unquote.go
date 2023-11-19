package xstring

import (
	"fmt"
	"strconv"
	"strings"
)

type state uint8

const (
	_ state = iota
	beginState
	endState
	normalState
	escapeState
	hexState
)

func (s state) String() string {
	switch s {
	case beginState:
		return "Begin"
	case normalState:
		return "Normal"
	case escapeState:
		return "Escape"
	case hexState:
		return "Hex"
	case endState:
		return "End"
	default:
		return "Unknown"
	}
}

func parseHex(b []rune) (rune, error) {
	if len(b) != 2 {
		return 0, fmt.Errorf("invalid hex length (%d)", len(b))
	}

	v, err := strconv.ParseUint(string(b), 16, 8)
	if err != nil {
		return 0, err
	}

	return rune(v), nil
}

func isHex(c rune) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f':
		return true
	default:
		return false
	}
}

func Unquote(s string) (string, error) {
	sb := strings.Builder{}

	st := beginState

	var hex []rune

	for _, c := range s {
		switch c {
		case '"':
			switch st {
			case beginState:
				st = normalState
			case normalState:
				st = endState
			case escapeState:
				sb.WriteRune(c)
				st = normalState
			case hexState:
				r, err := parseHex(hex)
				if err != nil {
					return "", fmt.Errorf("unexpected char %q at %d: %w", c, len(sb.String()), err)
				}
				sb.WriteRune(r)
				st = endState
			default:
				return "", fmt.Errorf("unexpected char %q at %d", c, len(sb.String()))
			}
		case '\\':
			switch st {
			case normalState:
				st = escapeState
			case escapeState:
				sb.WriteByte('\\')
				st = normalState
			case hexState:
				r, err := parseHex(hex)
				if err != nil {
					return "", fmt.Errorf("unexpected char %q at %d: %w", c, len(sb.String()), err)
				}
				sb.WriteRune(r)
				st = escapeState
			default:
				return "", fmt.Errorf("unexpected char %q at %d", c, len(sb.String()))
				//return "", fmt.Errorf("invalid transition (%s -> %s)", state, Escape)
			}
		case 'x':
			switch st {
			case normalState:
				sb.WriteRune('x')
			case hexState:
				r, err := parseHex(hex)
				if err != nil {
					return "", fmt.Errorf("unexpected char %q at %d: %w", c, len(sb.String()), err)
				}
				sb.WriteRune(r)
				sb.WriteRune('x')
				st = normalState
			case escapeState:
				st = hexState
				hex = hex[:0]
			default:
				return "", fmt.Errorf("unexpected char %q at %d", c, len(sb.String()))
			}
		default:
			switch st {
			case normalState:
				sb.WriteRune(c)
			case escapeState:
				return "", fmt.Errorf("unexpected char %q at %d", c, len(sb.String()))
			case hexState:
				if len(hex) == 2 {
					r, err := parseHex(hex)
					if err != nil {
						return "", fmt.Errorf("unexpected char %q at %d: %w", c, len(sb.String()), err)
					}
					sb.WriteRune(r)
					sb.WriteRune(c)
					st = normalState
					continue
				}

				if isHex(c) {
					hex = append(hex, c)
					continue
				}

				return "", fmt.Errorf("unexpected char %q at %d", c, len(sb.String()))

			default:
				return "", fmt.Errorf("unexpected char %q at %d", c, len(sb.String()))
			}
		}
	}

	switch st {
	case beginState:
		return "", fmt.Errorf("invalid input: empty string")
	case endState:
		return sb.String(), nil
	default:
		return "", fmt.Errorf("unexpected end of line (%s)", st)
	}
}
