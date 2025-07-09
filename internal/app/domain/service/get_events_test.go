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

func TestService_GetEvents(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	timeNow := time.Now()

	tests := []struct {
		name       string
		setupMock  func(m *mock.MockRepository)
		wantEvents []entity.Event
		wantErr    error
	}{
		{
			name: "success",
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetEvents(ctx).
					Return([]entity.Event{{ID: 1, Title: "title0", Date: timeNow}, {ID: 3, Title: "title", Date: timeNow}}, nil).
					Once()
			},
			wantEvents: []entity.Event{{ID: 1, Title: "title0", Date: timeNow}, {ID: 3, Title: "title", Date: timeNow}},
		},
		{
			name: "repository error",
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetEvents(ctx).
					Return(nil, errors.New("repository error")).
					Once()
			},
			wantErr: errors.New("repository.GetEvents: repository error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			events, err := NewService(repositoryMock).GetEvents(ctx)

			if test.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, test.wantEvents, events)
		})
	}
}
