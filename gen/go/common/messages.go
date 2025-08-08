package common

import (
	"fmt"

	structpb "google.golang.org/protobuf/types/known/structpb"
	"k8s.io/utils/ptr"
)

// NewFormattedMessage creates a new FormattedMessage with the given format and optional format arguments. The format string must be compatible with fmt.Sprintf.
func NewFormattedMessage(format string, a ...any) *FormattedMessage {
	tmp := make([]string, len(a))
	for i, v := range a {
		switch vs := v.(type) {
		case string:
			tmp[i] = vs
		case fmt.Stringer:
			tmp[i] = vs.String()
		default:
			tmp[i] = fmt.Sprintf("%v", v) // Fallback: convert anything else to string
		}
	}
	return FormattedMessage_builder{
		Message: ptr.To(format),
		Args:    tmp,
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
