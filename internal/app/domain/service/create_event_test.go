package service

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
	"github.com/xabsvoid/eis/internal/pkg/mock"
)

func TestService_CreateEvent(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	timeNow := time.Now().Add(time.Hour * 24)

	tests := []struct {
		name      string
		event     entity.Event
		setupMock func(m *mock.MockRepository)
		wantEvent entity.Event
		wantErr   error
	}{
		{
			name:  "success",
			event: entity.Event{Title: "Test Event", Date: timeNow},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreateEvent(ctx, entity.Event{Title: "Test Event", Date: timeNow}).
					Return(entity.Event{ID: 1, Title: "Test Event", Date: timeNow}, nil).
					Once()
			},
			wantEvent: entity.Event{ID: 1, Title: "Test Event", Date: timeNow},
			wantErr:   nil,
		},
		{
			name:  "repository error",
			event: entity.Event{Title: "Failing Event", Date: timeNow},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreateEvent(ctx, entity.Event{Title: "Failing Event", Date: timeNow}).
					Return(entity.Event{}, errors.New("failed to create event")).
					Once()
			},
			wantEvent: entity.Event{},
			wantErr:   errors.New("repository.CreateEvent: failed to create event"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			event, err := NewService(repositoryMock).CreateEvent(ctx, test.event)

			if test.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, test.wantEvent, event)
		})
	}
}
