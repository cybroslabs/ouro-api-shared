package helpers

import (
	"context"
	"errors"
	"maps"
	"slices"
	"sync"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
	"github.com/cybroslabs/ouro-api-shared/gen/go/helpers/database/postgres"
	"go.uber.org/zap"
	"k8s.io/utils/ptr"
)

type PathToDbPathFunc func(objectType common.ObjectType, fieldDescriptorMap map[string]string) postgres.PathToDbPathFunc

type FieldDescriptorManager interface {
	PathToDbPath(objectType common.ObjectType) postgres.PathToDbPathFunc
}

type FieldDescriptorManagerOpts struct {
	Logger       *zap.SugaredLogger
	Connectors   Connectors
	PathToDbPath PathToDbPathFunc

	RelatedObjectTypes []common.ObjectType
	SystemDescriptors  []*common.FieldDescriptorInternal
}

type fieldDescriptorManager struct {
	logger            *zap.SugaredLogger
	connectors        Connectors
	pathToDbPath      PathToDbPathFunc
	systemDescriptors []*common.FieldDescriptorInternal

	relatedObjectTypes         []common.ObjectType
	fieldDescriptorPathMapLock sync.RWMutex
	fieldDescriptorPathMap     map[common.ObjectType]map[string]string
	knownDescriptors           []*common.FieldDescriptorInternal
}

func NewFieldDescriptorManager(opts *FieldDescriptorManagerOpts) (FieldDescriptorManager, error) {
	r := &fieldDescriptorManager{
		logger:                 opts.Logger,
		connectors:             opts.Connectors,
		pathToDbPath:           opts.PathToDbPath,
		relatedObjectTypes:     slices.Clone(opts.RelatedObjectTypes),
		fieldDescriptorPathMap: make(map[common.ObjectType]map[string]string, 0),
		systemDescriptors:      slices.Clone(opts.SystemDescriptors),
		knownDescriptors:       make([]*common.FieldDescriptorInternal, 0, len(opts.SystemDescriptors)),
	}
	err := r.Refresh(context.Background())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (fdm *fieldDescriptorManager) PathToDbPath(objectType common.ObjectType) postgres.PathToDbPathFunc {
	// This lock does not cover modification of the fieldDescriptorPathMap lists but we don't modify the lists so in the worst case
	// the call to the returned function will use the old list reference.
	fdm.fieldDescriptorPathMapLock.RLock()
	defer fdm.fieldDescriptorPathMapLock.RUnlock()

	fd_list, ok := fdm.fieldDescriptorPathMap[objectType]
	if !ok {
		// This must not happen! All object types MUST be registered in FieldDescriptors.
		// Do NOT remove it, it is a safety reminder.
		panic("FieldDescriptors does not contain object type: " + objectType.String())
	}

	return fdm.pathToDbPath(objectType, fd_list)
}

func (fdm *fieldDescriptorManager) Refresh(ctx context.Context) error {
	fdm.fieldDescriptorPathMapLock.Lock()
	defer fdm.fieldDescriptorPathMapLock.Unlock()

	return fdm.refresh(ctx)
}

func (fdm *fieldDescriptorManager) refresh(ctx context.Context) error {
	cli, closer, err := fdm.connectors.OpenApiServiceClient()
	if err != nil {
		return err
	}
	defer closer()

	data, err := cli.ListFieldDescriptors(ctx, nil)
	if err != nil {
		return err
	}

	clear(fdm.fieldDescriptorPathMap)
	tmp := make(map[string]*common.FieldDescriptorInternal, len(data.GetItems())+len(fdm.systemDescriptors))

	for _, objectType := range fdm.relatedObjectTypes {
		fdm.fieldDescriptorPathMap[objectType] = make(map[string]string, 0)
	}

	for _, fd_wrapper := range data.GetItems() {
		fd := fd_wrapper.GetFieldDescriptor()
		if !slices.Contains(fdm.relatedObjectTypes, fd.GetObjectType()) {
			continue
		}
		fdm.fieldDescriptorPathMap[fd.GetObjectType()][fd.GetPath()] = fd_wrapper.GetDbPath()
		tmp[fd.GetGid()] = fd_wrapper
	}

	// Always overwrite by built-in system descriptors; just in case somebody creates the same FD manually or do some hard modification somehow...
	for _, fd_wrapper := range fdm.systemDescriptors {
		fd := fd_wrapper.GetFieldDescriptor()
		fdm.fieldDescriptorPathMap[fd.GetObjectType()][fd.GetPath()] = fd_wrapper.GetDbPath()
		tmp[fd.GetGid()] = fd_wrapper
	}

	fdm.knownDescriptors = slices.Collect(maps.Values(tmp))

	return nil
}

func (fdm *fieldDescriptorManager) validateFieldDescriptor(fdi *common.FieldDescriptorInternal, create bool) error {
	fd := fdi.GetFieldDescriptor()
	if fd == nil {
		return errors.New("field descriptor cannot be nil")
	}
	if !slices.Contains(fdm.relatedObjectTypes, fd.GetObjectType()) {
		return errors.New("invalid object type in field descriptor")
	}
	if !fd.GetIsUserDefined() {
		return errors.New("field descriptor must be user-defined")
	}

	// Yep, it's slow byt this is a rare operation so let's optimize for memory usage here.
	var found *common.FieldDescriptorInternal
	gid := fd.GetGid()
	for _, fd_wrapper := range fdm.knownDescriptors {
		if fd_wrapper.GetFieldDescriptor().GetGid() == gid {
			found = fd_wrapper
		}
	}

	if create && found != nil {
		return errors.New("field descriptor for this object type and path already exists")
	} else if !create && found == nil {
		return errors.New("field descriptor for this object type and path does not exist")
	} else if !create && !found.GetFieldDescriptor().GetIsUserDefined() {
		return errors.New("field descriptor for this object type and path is not user-defined")
	}

	return nil
}

func (fdm *fieldDescriptorManager) CreateFieldDescriptor(ctx context.Context, fdi *common.FieldDescriptorInternal) error {
	fdm.fieldDescriptorPathMapLock.Lock()
	defer fdm.fieldDescriptorPathMapLock.Unlock()

	// Validate the field descriptor
	if err := fdm.validateFieldDescriptor(fdi, true); err != nil {
		return err
	}

	data := slices.Clone(fdm.knownDescriptors)
	data = append(data, fdi)

	cli, closer, err := fdm.connectors.OpenApiServiceClient()
	if err != nil {
		return err
	}
	defer closer()

	_, err = cli.UpdateFieldDescriptors(ctx, common.UpdateFieldDescriptorsRequest_builder{
		Items:          data,
		CleanupMissing: ptr.To(true),
	}.Build())
	if err != nil {
		return err
	}

	err = fdm.refresh(ctx)

	return err
}

func (fdm *fieldDescriptorManager) UpdateFieldDescriptor(ctx context.Context, fdi *common.FieldDescriptorInternal) error {
	fdm.fieldDescriptorPathMapLock.Lock()
	defer fdm.fieldDescriptorPathMapLock.Unlock()

	// Validate the field descriptor
	if err := fdm.validateFieldDescriptor(fdi, false); err != nil {
		return err
	}

	data := make([]*common.FieldDescriptorInternal, 0, len(fdm.knownDescriptors))
	for _, fd_wrapper := range fdm.knownDescriptors {
		if fd_wrapper.GetFieldDescriptor().GetGid() == fdi.GetFieldDescriptor().GetGid() {
			data = append(data, fdi) // Replace the existing descriptor with the updated one
		} else {
			data = append(data, fd_wrapper) // Keep the existing descriptor
		}
	}

	cli, closer, err := fdm.connectors.OpenApiServiceClient()
	if err != nil {
		return err
	}
	defer closer()

	_, err = cli.UpdateFieldDescriptors(ctx, common.UpdateFieldDescriptorsRequest_builder{
		Items:          data,
		CleanupMissing: ptr.To(true),
	}.Build())
	if err != nil {
		return err
	}

	err = fdm.refresh(ctx)

	return err
}

func (fdm *fieldDescriptorManager) DeleteFieldDescriptor(ctx context.Context, selector *common.FieldDescriptorSelector) error {
	fdm.fieldDescriptorPathMapLock.Lock()
	defer fdm.fieldDescriptorPathMapLock.Unlock()

	found := false
	data := make([]*common.FieldDescriptorInternal, 0, len(fdm.knownDescriptors))
	for _, fd_wrapper := range fdm.knownDescriptors {
		fd := fd_wrapper.GetFieldDescriptor()
		if fd.GetGid() == selector.GetGid() && fd.GetObjectType() == selector.GetObjectType() {
			if !fd.GetIsUserDefined() {
				return errors.New("cannot delete a non-user-defined field descriptor")
			}
			found = true
			continue // Skip this descriptor, effectively deleting it
		}
		data = append(data, fd_wrapper) // Keep the existing descriptor
	}

	if !found {
		return errors.New("field descriptor not found")
	}

	cli, closer, err := fdm.connectors.OpenApiServiceClient()
	if err != nil {
		return err
	}
	defer closer()

	_, err = cli.UpdateFieldDescriptors(ctx, common.UpdateFieldDescriptorsRequest_builder{
		Items:          data,
		CleanupMissing: ptr.To(true),
	}.Build())
	if err != nil {
		return err
	}

	err = fdm.refresh(ctx)

	return err
}
