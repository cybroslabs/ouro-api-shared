package common

import (
	"fmt"
	"strings"

	structpb "google.golang.org/protobuf/types/known/structpb"
	"k8s.io/utils/ptr"
)

// NewFormattedMessage creates a new FormattedMessage with the given format and optional format arguments. The format string must be compatible with fmt.Sprintf.
func NewFormattedMessage(format string, a ...any) *FormattedMessage {
	aa := make([]string, 0, len(a))
	state := false
	var res strings.Builder
	var f strings.Builder
	for _, c := range format {
		if state {
			f.WriteRune(c)
			switch c { // terminators
			case '%': // special case
				res.WriteString("%") // fixed string even with invalid prefix like: %0%, still prints % only
				state = false
			case 'v', 'T', 't', 'b', 'c', 'd', 'o', 'O', 'q', 'x', 'X', 'U', 'e', 'E', 'f', 'F', 'g', 'G', 's', 'p':
				res.WriteString("%s")
				state = false
				if len(a) <= len(aa) {
					aa = append(aa, "invalid")
				} else {
					aa = append(aa, fmt.Sprintf(f.String(), a[len(aa)]))
				}
			case 'w': // special case for %w, without damn wrapping, avoid mistakes
				res.WriteString("%s")
				state = false
				if len(a) <= len(aa) {
					aa = append(aa, "invalid")
				} else {
					aa = append(aa, fmt.Sprintf("%v", a[len(aa)]))
				}
			}
		} else {
			if c == '%' {
				f.Reset()
				f.WriteRune(c)
				state = true
			} else {
				res.WriteRune(c)
			}
		}
	}

	return FormattedMessage_builder{
		Message: ptr.To(res.String()),
		Args:    aa,
	}.Build()
}

// SetParams sets the parameters for the FormattedMessage. It expects an even number of arguments, where each pair consists of a key and a value.
func (fm *FormattedMessage) AddParams(a ...any) error {
	a_cnt := len(a)
	if a_cnt%2 != 0 {
		return fmt.Errorf("SetParams requires an even number of arguments, got %d", len(a))
	}
	p := fm.GetParams().AsMap()
	for i := 0; i < a_cnt; i += 2 {
		key, ok := a[i].(string)
		if !ok {
			return fmt.Errorf("SetParams expects string keys, got %T at index %d", a[i], i)
		}
		p[key] = a[i+1]
	}
	if s, err := structpb.NewStruct(p); err != nil {
		return err
	} else {
		fm.SetParams(s)
	}
	return nil
}
