package common

import "github.com/google/uuid"

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
