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

func TestFinderAllHolidays_Execute_ShouldReturnsAHolidayList(t *testing.T) {
	t.Log("Should returns a holiday list")
	// Setup
	controller := gomock.NewController(t)

	repository := mocks.NewMockRepository(controller)

	extraName := "test extra name"
	holidaysList := []holiday.Holiday{{
		Date:        "test date",
		Title:       "test tittle",
		Inalienable: true,
		Extra:       "test Extra",
	}, {
		Date:        "test Date name 2",
		Title:       "test Tittle name 2",
		Inalienable: true,
		Extra:       "test Extra 2",
	}}
	repository.EXPECT().Find(gomock.Any(), extraName).Return(holidaysList, nil).Times(1)

	find := usecases.NewFinderDateHolidays(repository)

	// Execute
	result, err := find.Execute(context.TODO(), extraName)

	// Verify
	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, holidaysList, result)
}
