package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/eis/internal/pkg/mock"
)

func TestService_UpdateParticipantConfirmationByParticipantID(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	tests := []struct {
		name          string
		participantID int64
		exist         bool
		setupMock     func(m *mock.MockRepository)
		wantErr       error
	}{
		{
			name:          "successfully delete confirmation",
			participantID: 123,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					DeleteParticipantConfirmationByParticipantID(ctx, int64(123)).
					Return(nil).
					Once()
			},
		},
		{
			name:          "error while deleting confirmation",
			participantID: 123,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					DeleteParticipantConfirmationByParticipantID(ctx, int64(123)).
					Return(errors.New("delete error")).
					Once()
			},
			wantErr: errors.New("repository.DeleteParticipantConfirmationByParticipantID: delete error"),
		},
		{
			name:          "successfully create confirmation",
			participantID: 456,
			exist:         true,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreateParticipantConfirmation(ctx, int64(456)).
					Return(nil).
					Once()
			},
		},
		{
			name:          "error while creating confirmation",
			participantID: 456,
			exist:         true,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreateParticipantConfirmation(ctx, int64(456)).
					Return(errors.New("create error")).
					Once()
			},
			wantErr: errors.New("repository.CreateParticipantConfirmation: create error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			err := NewService(repositoryMock).UpdateParticipantConfirmationByParticipantID(ctx, test.participantID, test.exist)

			if test.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
