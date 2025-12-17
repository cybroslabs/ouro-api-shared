package common

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock object for testing
type mockObject struct {
	metadata *MetadataFields
}

func (m *mockObject) GetMetadata() *MetadataFields {
	return m.metadata
}

func (m *mockObject) SetMetadata(md *MetadataFields) {
	m.metadata = md
}

func TestCreateObjectMetadata_Success(t *testing.T) {
	obj := &mockObject{}
	userId := uuid.New()

	id, err := CreateObjectMetadata(obj, userId)
	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

	// Verify metadata was set
	assert.NotNil(t, obj.GetMetadata())
	assert.Equal(t, id.String(), obj.GetMetadata().GetId())
	assert.Equal(t, userId.String(), obj.GetMetadata().GetUserId())
	assert.Equal(t, int32(1), obj.GetMetadata().GetGeneration())
}

func TestCreateObjectMetadata_WithExistingMetadata(t *testing.T) {
	existingMd := MetadataFields_builder{}.Build()
	obj := &mockObject{metadata: existingMd}
	userId := uuid.New()

	id, err := CreateObjectMetadata(obj, userId)
	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

	// Verify same metadata object was used
	assert.Equal(t, existingMd, obj.GetMetadata())
	assert.Equal(t, id.String(), obj.GetMetadata().GetId())
	assert.Equal(t, userId.String(), obj.GetMetadata().GetUserId())
}

func TestCreateObjectMetadataV7_Success(t *testing.T) {
	obj := &mockObject{}
	userId := uuid.New()

	id, err := CreateObjectMetadataV7(obj, userId)
	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

	// UUIDv7 should have version 7
	assert.Equal(t, uuid.Version(7), id.Version())

	// Verify metadata was set
	assert.NotNil(t, obj.GetMetadata())
	assert.Equal(t, id.String(), obj.GetMetadata().GetId())
	assert.Equal(t, userId.String(), obj.GetMetadata().GetUserId())
	assert.Equal(t, int32(1), obj.GetMetadata().GetGeneration())
}

func TestCreateObjectMetadataV7_WithExistingMetadata(t *testing.T) {
	existingMd := MetadataFields_builder{}.Build()
	obj := &mockObject{metadata: existingMd}
	userId := uuid.New()

	id, err := CreateObjectMetadataV7(obj, userId)
	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

	// Verify same metadata object was used
	assert.Equal(t, existingMd, obj.GetMetadata())
	assert.Equal(t, id.String(), obj.GetMetadata().GetId())
}

func TestUpdateObjectMetadataUserId_Success(t *testing.T) {
	md := MetadataFields_builder{}.Build()
	obj := &mockObject{metadata: md}

	oldUserId := uuid.New()
	md.SetUserId(oldUserId.String())

	newUserId := uuid.New()
	err := UpdateObjectMetadataUserId(obj, newUserId)
	require.NoError(t, err)

	assert.Equal(t, newUserId.String(), obj.GetMetadata().GetUserId())
}

func TestUpdateObjectMetadataUserId_NoMetadata(t *testing.T) {
	obj := &mockObject{}
	userId := uuid.New()

	err := UpdateObjectMetadataUserId(obj, userId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "metadata is not set")
}

func TestMergeObjectMetadataFields_Success(t *testing.T) {
	md := MetadataFields_builder{}.Build()
	id := uuid.New()
	md.SetId(id.String())
	md.SetGeneration(1)

	obj := &mockObject{metadata: md}

	// Add some fields
	fieldsData := map[string]*FieldValue{
		"field1": FieldValue_builder{StringValue: stringPtr("value1")}.Build(),
		"field2": FieldValue_builder{IntegerValue: int64Ptr(42)}.Build(),
	}

	err := MergeObjectMetadataFields(obj, fieldsData)
	require.NoError(t, err)

	// Verify fields were set
	assert.Len(t, obj.GetMetadata().GetFields(), 2)
	assert.Equal(t, "value1", obj.GetMetadata().GetFields()["field1"].GetStringValue())
	assert.Equal(t, int64(42), obj.GetMetadata().GetFields()["field2"].GetIntegerValue())

	// Verify generation was incremented
	assert.Equal(t, int32(2), obj.GetMetadata().GetGeneration())
}

func TestMergeObjectMetadataFields_MergeWithExisting(t *testing.T) {
	md := MetadataFields_builder{}.Build()
	id := uuid.New()
	md.SetId(id.String())
	md.SetGeneration(1)

	// Set initial fields
	initialFields := map[string]*FieldValue{
		"field1": FieldValue_builder{StringValue: stringPtr("initial")}.Build(),
	}
	md.SetFields(initialFields)

	obj := &mockObject{metadata: md}

	// Merge new fields
	newFields := map[string]*FieldValue{
		"field2": FieldValue_builder{IntegerValue: int64Ptr(100)}.Build(),
	}

	err := MergeObjectMetadataFields(obj, newFields)
	require.NoError(t, err)

	// Verify both fields exist
	assert.Len(t, obj.GetMetadata().GetFields(), 2)
	assert.Equal(t, "initial", obj.GetMetadata().GetFields()["field1"].GetStringValue())
	assert.Equal(t, int64(100), obj.GetMetadata().GetFields()["field2"].GetIntegerValue())

	// Verify generation was incremented
	assert.Equal(t, int32(2), obj.GetMetadata().GetGeneration())
}

func TestMergeObjectMetadataFields_OverwriteExisting(t *testing.T) {
	md := MetadataFields_builder{}.Build()
	md.SetGeneration(1)

	// Set initial fields
	initialFields := map[string]*FieldValue{
		"field1": FieldValue_builder{StringValue: stringPtr("initial")}.Build(),
	}
	md.SetFields(initialFields)

	obj := &mockObject{metadata: md}

	// Merge with same key
	newFields := map[string]*FieldValue{
		"field1": FieldValue_builder{StringValue: stringPtr("updated")}.Build(),
	}

	err := MergeObjectMetadataFields(obj, newFields)
	require.NoError(t, err)

	// Verify field was overwritten
	assert.Len(t, obj.GetMetadata().GetFields(), 1)
	assert.Equal(t, "updated", obj.GetMetadata().GetFields()["field1"].GetStringValue())
}

func TestMergeObjectMetadataFields_NoMetadata(t *testing.T) {
	obj := &mockObject{}
	fieldsData := map[string]*FieldValue{
		"field1": FieldValue_builder{StringValue: stringPtr("value1")}.Build(),
	}

	err := MergeObjectMetadataFields(obj, fieldsData)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "metadata fields are not set")
}

func TestMetadataFields_Create_Success(t *testing.T) {
	md := &MetadataFields{}

	id, err := md.Create()
	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

	// Verify UUID version 4
	assert.Equal(t, uuid.Version(4), id.Version())

	// Verify metadata was set
	assert.Equal(t, id.String(), md.GetId())
	assert.Equal(t, int32(1), md.GetGeneration())
}

func TestMetadataFields_Create_NilMetadata(t *testing.T) {
	var md *MetadataFields
	id, err := md.Create()
	assert.NoError(t, err)
	assert.Equal(t, uuid.Nil, id)
}

func TestMetadataFields_Create_WithManagedFields(t *testing.T) {
	md := &MetadataFields{}
	// Set managed fields (should cause error)
	md.SetManagedFields(map[string]*FieldValue{
		"managed": FieldValue_builder{StringValue: stringPtr("value")}.Build(),
	})

	id, err := md.Create()
	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, id)
	assert.ErrorIs(t, err, ErrOuroManagedFieldsMustBeEmpty)
}

func TestMetadataFields_CreateV7_Success(t *testing.T) {
	md := &MetadataFields{}

	id, err := md.CreateV7()
	require.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

	// Verify UUID version 7
	assert.Equal(t, uuid.Version(7), id.Version())

	// Verify metadata was set
	assert.Equal(t, id.String(), md.GetId())
	assert.Equal(t, int32(1), md.GetGeneration())
}

func TestMetadataFields_CreateV7_NilMetadata(t *testing.T) {
	var md *MetadataFields
	id, err := md.CreateV7()
	assert.NoError(t, err)
	assert.Equal(t, uuid.Nil, id)
}

func TestMetadataFields_CreateV7_WithManagedFields(t *testing.T) {
	md := &MetadataFields{}
	// Set managed fields (should cause error)
	md.SetManagedFields(map[string]*FieldValue{
		"managed": FieldValue_builder{StringValue: stringPtr("value")}.Build(),
	})

	id, err := md.CreateV7()
	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, id)
	assert.ErrorIs(t, err, ErrOuroManagedFieldsMustBeEmpty)
}

func TestMetadataFields_IncGeneration(t *testing.T) {
	md := &MetadataFields{}
	md.SetGeneration(1)

	md.IncGeneration()
	assert.Equal(t, int32(2), md.GetGeneration())

	md.IncGeneration()
	assert.Equal(t, int32(3), md.GetGeneration())
}

func TestMetadataFields_IncGeneration_FromZero(t *testing.T) {
	md := &MetadataFields{}
	// Default generation is 0
	assert.Equal(t, int32(0), md.GetGeneration())

	md.IncGeneration()
	assert.Equal(t, int32(1), md.GetGeneration())
}

func TestMetadataFields_Id_Success(t *testing.T) {
	md := &MetadataFields{}
	expectedId := uuid.New()
	md.SetId(expectedId.String())

	id := md.Id()
	assert.Equal(t, expectedId, id)
}

func TestMetadataFields_Id_NilMetadata(t *testing.T) {
	var md *MetadataFields
	id := md.Id()
	assert.Equal(t, uuid.Nil, id)
}

func TestMetadataFields_Id_InvalidUUID(t *testing.T) {
	md := &MetadataFields{}
	md.SetId("invalid-uuid")

	id := md.Id()
	assert.Equal(t, uuid.Nil, id)
}

func TestMetadataFields_Id_EmptyString(t *testing.T) {
	md := &MetadataFields{}
	md.SetId("")

	id := md.Id()
	assert.Equal(t, uuid.Nil, id)
}

func TestCreateObjectMetadata_UUIDUniqueness(t *testing.T) {
	userId := uuid.New()
	ids := make(map[uuid.UUID]bool)

	// Create multiple objects and verify UUIDs are unique
	for range 100 {
		obj := &mockObject{}
		id, err := CreateObjectMetadata(obj, userId)
		require.NoError(t, err)

		// Verify UUID is unique
		assert.False(t, ids[id], "UUID collision detected")
		ids[id] = true
	}
}

func TestCreateObjectMetadataV7_UUIDUniqueness(t *testing.T) {
	userId := uuid.New()
	ids := make(map[uuid.UUID]bool)

	// Create multiple objects and verify UUIDs are unique
	for range 100 {
		obj := &mockObject{}
		id, err := CreateObjectMetadataV7(obj, userId)
		require.NoError(t, err)

		// Verify UUID is unique and version 7
		assert.False(t, ids[id], "UUID collision detected")
		assert.Equal(t, uuid.Version(7), id.Version())
		ids[id] = true
	}
}

func TestMergeObjectMetadataFields_EmptyMerge(t *testing.T) {
	md := MetadataFields_builder{}.Build()
	md.SetGeneration(1)
	obj := &mockObject{metadata: md}

	// Merge empty fields
	err := MergeObjectMetadataFields(obj, map[string]*FieldValue{})
	require.NoError(t, err)

	// Generation should still be incremented
	assert.Equal(t, int32(2), obj.GetMetadata().GetGeneration())
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}

// Benchmark tests
func BenchmarkCreateObjectMetadata(b *testing.B) {
	userId := uuid.New()

	for b.Loop() {
		obj := &mockObject{}
		_, _ = CreateObjectMetadata(obj, userId)
	}
}

func BenchmarkCreateObjectMetadataV7(b *testing.B) {
	userId := uuid.New()

	for b.Loop() {
		obj := &mockObject{}
		_, _ = CreateObjectMetadataV7(obj, userId)
	}
}

func BenchmarkMetadataFields_Id(b *testing.B) {
	md := &MetadataFields{}
	id := uuid.New()
	md.SetId(id.String())

	for b.Loop() {
		_ = md.Id()
	}
}

func BenchmarkMergeObjectMetadataFields(b *testing.B) {
	md := MetadataFields_builder{}.Build()
	md.SetGeneration(1)
	obj := &mockObject{metadata: md}

	fieldsData := map[string]*FieldValue{
		"field1": FieldValue_builder{StringValue: stringPtr("value1")}.Build(),
		"field2": FieldValue_builder{IntegerValue: int64Ptr(42)}.Build(),
	}

	for b.Loop() {
		_ = MergeObjectMetadataFields(obj, fieldsData)
	}
}
