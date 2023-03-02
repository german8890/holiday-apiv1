package usecases_test

import (
	"context"
	"holidays-seeker/internal/core/domain/holiday"
	"holidays-seeker/internal/core/domain/holiday/mocks"
	"holidays-seeker/internal/core/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFinderHolidaysByExtra_Execute_ShouldReturnsAHolidayData(t *testing.T) {
	t.Log("Should returns a holiday from his extra")
	// Setup
	controller := gomock.NewController(t)

	extra := "test extra"
	holidayResult := holiday.Holiday{
		Date:        "test date ",
		Title:       "test title",
		Type:        "test type",
		Inalienable: true,
		Extra:       "test extra ",
	}

	repository := mocks.NewMockRepository(controller)
	repository.EXPECT().FindOne(gomock.Any(), extra).Return(holidayResult, nil).Times(1)

	find := usecases.NewFinderHolidaysByExtra(repository)

	// Execute
	result, err := find.Execute(context.TODO(), extra)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, holidayResult, result)
}

func TestFinderHolidayByExtraName_Execute_ShouldReturnsAnErrorForInvalidName(t *testing.T) {
	t.Log("Should returns an error for invalid extra name")
	// Setup
	controller := gomock.NewController(t)

	extraName := ""
	repository := mocks.NewMockRepository(controller)

	find := usecases.NewFinderHolidaysByExtra(repository)

	// Execute
	result, err := find.Execute(context.TODO(), extraName)

	// Verify
	require.Error(t, err, "Invalid name")
	assert.Equal(t, holiday.Holiday{}, result)
}
