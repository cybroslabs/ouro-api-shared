package helpers

import (
	"context"
	"errors"
	"reflect"
	"slices"
	"sync"
	"time"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
	"github.com/cybroslabs/ouro-api-shared/gen/go/system"
)

type ConfigurationServiceOpts struct {
	Connectors    Connectors
	CacheTime     time.Duration
	ComponentName string
	Descriptors   []*common.FieldDescriptor
}

type ConfigurationService interface {
	UpdateDescriptors(descriptors []*common.FieldDescriptor)
	FlushCache()
	GetOption(optionName string) (any, error)
	GetOptionWithDefault(optionName string, defaultValue any) (any, error)
}

type configurationService struct {
	ConfigurationService

	connectors Connectors

	componentName string
	descriptors   []*common.FieldDescriptor

	cacheTime time.Duration

	cacheLock sync.RWMutex
	cacheLast time.Time
	cache     map[string]any
}

func NewConfigurationService(opts *ConfigurationServiceOpts) (ConfigurationService, error) {
	if opts == nil {
		return nil, errors.New("the ConfigurationServiceOpts cannot be nil")
	}
	if opts.Connectors == nil {
		return nil, errors.New("the Connectors cannot be nil")
	}
	if len(opts.ComponentName) == 0 {
		return nil, errors.New("the ComponentName cannot be empty")
	}

	var descriptors []*common.FieldDescriptor
	if opts.Descriptors != nil {
		// Create a local copy of the descriptor so that the items can be cleaned up after the first call
		descriptors = slices.Clone(opts.Descriptors)
	}

	return &configurationService{
		connectors:    opts.Connectors,
		componentName: opts.ComponentName,
		descriptors:   descriptors,
		cacheTime:     opts.CacheTime,
		cache:         make(map[string]any),
		cacheLast:     time.Time{},
	}, nil
}

func (cs *configurationService) getConfiguration() (*system.ComponentConfig, error) {
	cli, close, err := cs.connectors.OpenOuroOperatorServiceClient()
	if err != nil {
		return nil, err
	}
	defer close()

	config, err := cli.SynchronizeComponentConfig(context.Background(), system.ComponentConfigDescriptor_builder{
		Name:  &cs.componentName,
		Items: cs.descriptors,
	}.Build())
	if err != nil {
		return nil, err
	}
	cs.descriptors = nil // Clear descriptor after the first successful call to avoid repeated synchronization

	return config, nil
}

func (cs *configurationService) loadConfiguration() error {
	cs.cacheLock.Lock()
	defer cs.cacheLock.Unlock()

	config, err := cs.getConfiguration()
	if err != nil {
		return err
	}

	cs.cache = make(map[string]any)
	for k, v := range config.GetItems().GetAttributes() {
		cs.cache[k] = v.GetAnyValue()
	}

	cs.cacheLast = time.Now()

	return nil
}

func (cs *configurationService) FlushCache() {
	cs.cacheLock.Lock()
	defer cs.cacheLock.Unlock()

	cs.cacheLast = time.Time{} // Reset cache timestamp to force reload on next GetOption call
}

func (cs *configurationService) UpdateDescriptors(descriptors []*common.FieldDescriptor) {
	if descriptors == nil {
		return
	}

	cs.cacheLock.Lock()
	defer cs.cacheLock.Unlock()

	// Re-set the descriptor with the new items
	cs.descriptors = slices.Clone(descriptors)
	cs.cacheLast = time.Time{} // Reset cache timestamp to force reload on next GetOption call
}

func (cs *configurationService) GetOption(optionName string) (any, error) {
	cs.cacheLock.RLock()
	t_delta := time.Since(cs.cacheLast)
	cs.cacheLock.RUnlock()

	if t_delta > cs.cacheTime {
		if err := cs.loadConfiguration(); err != nil {
			return nil, err
		}
	}

	cs.cacheLock.RLock()
	defer cs.cacheLock.RUnlock()

	if value, ok := cs.cache[optionName]; ok {
		return value, nil
	}

	// Not found
	return nil, errors.ErrUnsupported
}

func (cs *configurationService) GetOptionWithDefault(optionName string, defaultValue any) (any, error) {
	value, err := cs.GetOption(optionName)
	if err != nil {
		if errors.Is(err, errors.ErrUnsupported) {
			return defaultValue, nil
		}
		return nil, err
	}
	return value, nil
}

func GetOptionWithDefault[T any](cfg ConfigurationService, name string, defaultValue T) T {
	val, err := cfg.GetOption(name)
	if err != nil {
		return defaultValue
	}

	// Direct type support
	if v, ok := val.(T); ok {
		return v
	}

	// Fallback: try reflection-based conversion
	valValue := reflect.ValueOf(val)
	defValue := reflect.ValueOf(defaultValue)

	if valValue.Type().ConvertibleTo(defValue.Type()) {
		converted := valValue.Convert(defValue.Type())
		if v, ok := converted.Interface().(T); ok {
			return v
		}
	}

	return defaultValue
}
