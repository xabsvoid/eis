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

func TestService_GetEventByID(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	timeNow := time.Now()

	tests := []struct {
		name      string
		id        int64
		setupMock func(m *mock.MockRepository)
		wantEvent entity.Event
		wantErr   error
	}{
		{
			name: "success",
			id:   1,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetEventByID(ctx, int64(1)).
					Return(entity.Event{ID: 3, Title: "title", Date: timeNow}, nil).
					Once()
			},
			wantEvent: entity.Event{ID: 3, Title: "title", Date: timeNow},
		},
		{
			name: "repository error",
			id:   2,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetEventByID(ctx, int64(2)).
					Return(entity.Event{}, errors.New("repository error")).
					Once()
			},
			wantErr: errors.New("repository.GetEventByID: repository error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			event, err := NewService(repositoryMock).GetEventByID(ctx, test.id)

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
