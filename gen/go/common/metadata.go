package common

import (
	"errors"
	"maps"

	"github.com/google/uuid"
)

type objectMetadataGet interface {
	GetMetadata() *MetadataFields
}

type objectMetadata interface {
	objectMetadataGet
	SetMetadata(*MetadataFields)
}

// CreateObjectMetadata creates a new UUIDv4 and generation 1 in the Metadata.Id for the given object.
func CreateObjectMetadata[T objectMetadata](obj T, userId uuid.UUID) (id uuid.UUID, err error) {
	md := obj.GetMetadata()
	if md == nil {
		md = MetadataFields_builder{}.Build()
		obj.SetMetadata(md)
	}
	md.SetUserId(userId.String())
	return md.Create()
}

// CreateObjectMetadataV7 creates a new UUIDv7 and generation 1 in the Metadata.Id for the given object.
func CreateObjectMetadataV7[T objectMetadata](obj T, userId uuid.UUID) (id uuid.UUID, err error) {
	md := obj.GetMetadata()
	if md == nil {
		md = MetadataFields_builder{}.Build()
		obj.SetMetadata(md)
	}
	md.SetUserId(userId.String())
	return md.CreateV7()
}

// UpdateObjectMetadataUserId updates the Metadata.UserId for the given object.
// If metadata is not set, an error is returned.
func UpdateObjectMetadataUserId[T objectMetadataGet](obj T, userId uuid.UUID) error {
	md := obj.GetMetadata()
	if md == nil {
		return errors.New("metadata is not set")
	}
	md.SetUserId(userId.String())
	return nil
}

// MergeMetadataFields merges the given fieldsData into the Metadata.Fields of the given object, incrementing the generation.
// If metadata is not set, an error is returned. The values are a shallow copy.
func MergeObjectMetadataFields[T objectMetadataGet](obj T, fieldsData map[string]*FieldValue) error {
	md := obj.GetMetadata()
	if md == nil {
		return errors.New("metadata fields are not set")
	}

	mdf := md.GetFields()
	if mdf == nil {
		md.SetFields(fieldsData)
	} else {
		maps.Copy(mdf, fieldsData)
	}
	md.IncGeneration()
	return nil
}

// Create creates a new Metadata object with a new UUIDv4 and generation 1.
func (x *MetadataFields) Create() (id uuid.UUID, err error) {
	if x == nil {
		return uuid.Nil, nil
	}
	if len(x.GetManagedFields()) > 0 {
		return uuid.Nil, ErrOuroManagedFieldsMustBeEmpty
	}
	id, err = uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}
	x.SetId(id.String())
	x.SetGeneration(1)
	return id, nil
}

// Create creates a new Metadata object with a new UUIDv7 and generation 1.
func (x *MetadataFields) CreateV7() (id uuid.UUID, err error) {
	if x == nil {
		return uuid.Nil, nil
	}
	if len(x.GetManagedFields()) > 0 {
		return uuid.Nil, ErrOuroManagedFieldsMustBeEmpty
	}
	id, err = uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}
	x.SetId(id.String())
	x.SetGeneration(1)
	return id, nil
}

// IncGeneration increments the generation of the object.
func (x *MetadataFields) IncGeneration() {
	x.SetGeneration(x.GetGeneration() + 1)
}

// Id returns the UUID of the object or nil if the object is nil or the ID is invalid.
func (x *MetadataFields) Id() uuid.UUID {
	if x == nil {
		return uuid.Nil
	}
	id, err := uuid.Parse(x.GetId())
	if err != nil {
		return uuid.Nil
	}
	return id
}
