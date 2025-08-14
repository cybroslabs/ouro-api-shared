package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
)

var (
	_emptyProtoDescriptorData = &protoDescriptorData{}
)

type protoDescriptorDataNested struct {
	FieldDescriptor protoreflect.FieldDescriptor
	NestedMessage   protoreflect.MessageDescriptor
}

type protoDescriptorData struct {
	Fields         map[string]protoreflect.FieldDescriptor
	NestedMessages map[string]protoDescriptorDataNested
}

type protoValidator struct {
	logger                 *zap.SugaredLogger
	fieldDescriptorManager FieldDescriptorManager

	cacheLock sync.RWMutex
	cache     map[protoreflect.MessageDescriptor]*protoDescriptorData
}

func newProtoValidator(logger *zap.SugaredLogger, fieldDescriptorManager FieldDescriptorManager) *protoValidator {
	return &protoValidator{
		logger:                 logger,
		fieldDescriptorManager: fieldDescriptorManager,
		cache:                  make(map[protoreflect.MessageDescriptor]*protoDescriptorData),
	}
}

func (pv *protoValidator) fetchMesssageDescriptor(md protoreflect.MessageDescriptor) *protoDescriptorData {
	pv.cacheLock.RLock()
	if cached, ok := pv.cache[md]; ok {
		pv.cacheLock.RUnlock()
		return cached
	}
	pv.cacheLock.RUnlock()

	pv.cacheLock.Lock()
	defer pv.cacheLock.Unlock()

	nested_fields := pv.fetchMesssageDescriptorInternal(md)
	pv.cache[md] = nested_fields
	return nested_fields
}

func (pv *protoValidator) fetchMesssageDescriptorInternal(md protoreflect.MessageDescriptor) *protoDescriptorData {
	if result, done := pv.cache[md]; done {
		return result
	}

	fields := md.Fields()
	if fields == nil {
		pv.cache[md] = _emptyProtoDescriptorData
		return nil
	}

	result := &protoDescriptorData{
		Fields:         make(map[string]protoreflect.FieldDescriptor),
		NestedMessages: make(map[string]protoDescriptorDataNested),
	}
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		fd_name := fd.JSONName()
		if nested_message := fd.Message(); nested_message != nil {
			result.NestedMessages[fd_name] = protoDescriptorDataNested{
				FieldDescriptor: fd,
				NestedMessage:   nested_message,
			}
			_ = pv.fetchMesssageDescriptorInternal(nested_message)
		} else {
			result.Fields[fd_name] = fd
		}
	}

	pv.cache[md] = result
	return result
}

func (pv *protoValidator) validateInternal(m protoreflect.Message, fieldDescriptors map[string]*common.FieldDescriptor, fdPrefix string) error {
	if m == nil {
		return nil
	}

	md := m.Descriptor()
	if md == nil {
		return nil
	}

	pfd_map := pv.fetchMesssageDescriptor(md)

	for k, fd := range pfd_map.Fields {
		if fdPrefix != "" {
			k = fdPrefix + "." + k
		}
		pv.logger.Infof("Validating field: %s", k)
		if fdv, ok := fieldDescriptors[k]; !ok {
			continue
		} else {
			if !m.Has(fd) {
				if fdv.GetRequired() {
					return errors.New("required field " + k + " is missing")
				}
				continue
			}
			validation := fdv.GetValidation()
			if validation == nil {
				continue
			}

			v := m.Get(fd)
			switch fd.Kind() {
			case protoreflect.StringKind:
				if validation.HasMinLength() {
					min_length := int(validation.GetMinLength())
					if len(v.String()) < int(min_length) {
						return fmt.Errorf("field %s must have at least %d characters", k, min_length)
					}
				}
				if validation.HasMaxLength() {
					max_length := int(validation.GetMaxLength())
					if len(v.String()) > max_length {
						return fmt.Errorf("field %s must have at most %d characters", k, max_length)
					}
				}
				if validation.HasRe() {
					re, err := regexp.Compile(validation.GetRe())
					if err != nil {
						return errors.New("field " + k + " has invalid regex: " + validation.GetRe())
					}
					if !re.MatchString(v.String()) {
						return errors.New("field " + k + " does not match regex: " + re.String())
					}
				}
			case protoreflect.Int32Kind, protoreflect.Int64Kind:
				if validation.HasMinInteger() {
					min_value := validation.GetMinInteger()
					if v.Int() < min_value {
						return fmt.Errorf("field %s must be greater than or equal to %d", k, min_value)
					}
				}
				if validation.HasMaxInteger() {
					max_value := validation.GetMaxInteger()
					if v.Int() > max_value {
						return fmt.Errorf("field %s must be less than or equal to %d", k, max_value)
					}
				}
			case protoreflect.Uint32Kind, protoreflect.Uint64Kind:
				if validation.HasMinInteger() {
					min_value := validation.GetMinInteger()
					if int64(v.Uint()) < min_value {
						return fmt.Errorf("field %s must be greater than or equal to %d", k, min_value)
					}
				}
				if validation.HasMaxInteger() {
					max_value := validation.GetMaxInteger()
					if int64(v.Uint()) > max_value {
						return fmt.Errorf("field %s must be less than or equal to %d", k, max_value)
					}
				}
			}
		}
	}

	for k, nested_md := range pfd_map.NestedMessages {
		if !m.Has(nested_md.FieldDescriptor) {
			// TODO: The nested message is not set but there might be required fields in it.
			continue
		}
		nested_m := m.Get(nested_md.FieldDescriptor)
		nested_m_typed := nested_m.Interface().(protoreflect.Message)
		var nested_fd_prefix string
		if fdPrefix != "" {
			nested_fd_prefix = fdPrefix + "." + k
		} else {
			nested_fd_prefix = k
		}
		if err := pv.validateInternal(nested_m_typed, fieldDescriptors, nested_fd_prefix); err != nil {
			return err
		}
	}

	return nil
}

func (pv *protoValidator) ValidateMessage(objectType common.ObjectType, message protoreflect.Message, fieldDescriptorsMap map[string]*common.FieldDescriptor) error {
	return pv.validateInternal(message, fieldDescriptorsMap, "")
}
