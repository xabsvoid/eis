package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/eis/internal/pkg/mock"
)

func TestService_DeleteEventByID(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	tests := []struct {
		name      string
		id        int64
		setupMock func(m *mock.MockRepository)
		wantErr   error
	}{
		{
			name: "success",
			id:   1,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					DeleteEventByID(ctx, int64(1)).
					Return(nil).
					Once()
			},
		},
		{
			name: "repository error",
			id:   2,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					DeleteEventByID(ctx, int64(2)).
					Return(errors.New("repository error")).
					Once()
			},
			wantErr: errors.New("repository.DeleteEventByID: repository error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			err := NewService(repositoryMock).DeleteEventByID(ctx, test.id)

			if test.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
