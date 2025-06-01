package helpers

import (
	"context"
	"errors"
	"slices"
	"sync"
	"time"

	"github.com/cybroslabs/hes-2-apis/gen/go/system"
	"k8s.io/utils/ptr"
)

type ConfigurationServiceOpts struct {
	Connectors Connectors
	CacheTime  time.Duration
	Descriptor *system.ComponentConfigDescriptor
}

type ConfigurationService interface {
	GetOption(optionName string) (any, error)
	GetOptionWithDefault(optionName string, defaultValue any) (any, error)
}

type configurationService struct {
	ConfigurationService

	connectors Connectors
	descriptor *system.ComponentConfigDescriptor

	cacheTime time.Duration

	cacheLock sync.Mutex
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
	if opts.Descriptor == nil {
		return nil, errors.New("the Descriptor cannot be nil")
	}

	return &configurationService{
		connectors: opts.Connectors,
		// Create a local copy of the descriptor so that the items can be cleaned up after the first call
		descriptor: system.ComponentConfigDescriptor_builder{
			Name:  ptr.To(opts.Descriptor.GetName()),
			Items: slices.Clone(opts.Descriptor.GetItems()),
		}.Build(),
		cacheTime: opts.CacheTime,
		cache:     make(map[string]any),
		cacheLast: time.Time{},
	}, nil
}

func (cs *configurationService) getConfiguration() (*system.ComponentConfig, error) {
	cli, close, err := cs.connectors.OpenOuroOperatorServiceClient()
	if err != nil {
		return nil, err
	}
	defer close()

	config, err := cli.SynchronizeComponentConfig(context.Background(), cs.descriptor)
	if err != nil {
		return nil, err
	}
	cs.descriptor = nil // Clear descriptor after the first successful call to avoid repeated synchronization

	return config, nil
}

func (cs *configurationService) GetOption(optionName string) (any, error) {
	cs.cacheLock.Lock()
	defer cs.cacheLock.Unlock()

	if cs.cache == nil || time.Since(cs.cacheLast) > cs.cacheTime {
		config, err := cs.getConfiguration()
		if err != nil {
			return nil, err
		}

		cs.cache = make(map[string]any)
		for k, v := range config.GetItems().GetAttributes() {
			cs.cache[k] = v.GetIntegerValue()
		}
		cs.cacheLast = time.Now()
	}

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
	if v, ok := val.(T); ok {
		return v
	}
	return defaultValue
}
