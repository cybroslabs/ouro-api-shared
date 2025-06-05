package postgres

import (
	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
	"github.com/google/uuid"
)

// DbSelector is a structure used to select database records based on various criteria.
type DbSelector struct {
	Id       []uuid.UUID
	FilterBy *common.ListSelector

	err error
}

// PersistentWhere is a structure used to represent a persistent query condition.
type PersistentWhere struct {
	Query string
	Arg   any
}

func WithListSelector(selector *common.ListSelector) *DbSelector {
	return &DbSelector{FilterBy: selector}
}

func WithRawId(id string) *DbSelector {
	u, _ := uuid.Parse(id)
	return &DbSelector{Id: []uuid.UUID{u}}
}

func WithId(id uuid.UUID) *DbSelector {
	return &DbSelector{Id: []uuid.UUID{id}}
}

func (s *DbSelector) Err() error {
	if s != nil {
		return s.err
	}
	return nil
}

func (s *DbSelector) GetFilterBy() []*common.ListSelectorFilterBy {
	if s != nil {
		return s.FilterBy.GetFilterBy()
	}
	return nil
}

func (s *DbSelector) GetSortBy() []*common.ListSelectorSortBy {
	if s != nil {
		return s.FilterBy.GetSortBy()
	}
	return nil
}

func (s *DbSelector) GetOffset() uint32 {
	if s != nil {
		return s.FilterBy.GetOffset()
	}
	return 0
}

func (s *DbSelector) GetPageSize() uint32 {
	if s != nil {
		return s.FilterBy.GetPageSize()
	}
	return 0
}
