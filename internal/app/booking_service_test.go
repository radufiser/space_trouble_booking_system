package app_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"spacetrouble.com/booking/internal/app"
	"spacetrouble.com/booking/internal/domain"
	"spacetrouble.com/booking/internal/domain/mocks"
)

func TestCreateBooking_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	// Arrange
	mockBookingRepo := mocks.NewMockBookingRepository(ctrl)
	mockScheduleRepo := mocks.NewMockScheduleRepository(ctrl)
	mockDestinationRepo := mocks.NewMockDestinationRepository(ctrl)
	mockLaunchpadRepo := mocks.NewMockLaunchpadRepository(ctrl)
	mockLaunchClient := mocks.NewMockLaunchClient(ctrl)

	bookingService := app.NewBookingService(mockBookingRepo, mockScheduleRepo, mockDestinationRepo, mockLaunchpadRepo, mockLaunchClient)

	booking := &domain.Booking{
		ID:            uuid.New().String(),
		FirstName:     "John",
		LastName:      "Doe",
		Birthday:      time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		LaunchpadID:   "5e9e4502f509094188566f88",
		DestinationID: "29a8f36a-14eb-47a3-beeb-bf48b5d2fefe",
		LaunchDate:    time.Now().AddDate(0, 0, 1),
	}

	// Set up mock expectations
	mockDestinationRepo.EXPECT().GetByID(booking.DestinationID).Return(&domain.Destination{ID: booking.DestinationID, Name: "Mars"}, nil)
	mockLaunchpadRepo.EXPECT().GetByID(booking.LaunchpadID).Return(&domain.Launchpad{ID: booking.LaunchpadID, Name: "KSC LC 39A"}, nil)
	mockScheduleRepo.EXPECT().FetchSchedule(booking.LaunchpadID, int(booking.LaunchDate.Weekday()), booking.DestinationID).Return(&domain.WeeklySchedule{}, nil)
	mockLaunchClient.EXPECT().GetUpcomingLaunches().Return([]domain.Launch{}, nil)
	mockBookingRepo.EXPECT().Create(booking).Return(nil)

	// Act
	err := bookingService.CreateBooking(booking)

	// Assert
	assert.NoError(t, err)
}
