package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
	"github.com/xabsvoid/eis/internal/pkg/mock"
)

func TestService_UpdateParticipantReplacementByParticipantID(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	tests := []struct {
		name          string
		participantID int64
		exist         bool
		person        entity.Person
		setupMock     func(m *mock.MockRepository)
		wantErr       error
	}{
		{
			name:          "non-existent replacement - successful deletion",
			participantID: 1,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					DeleteParticipantReplacementByParticipantID(ctx, int64(1)).
					Return(nil).
					Once()
			},
		},
		{
			name:          "non-existent replacement - deletion error",
			participantID: 2,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					DeleteParticipantReplacementByParticipantID(ctx, int64(2)).
					Return(errors.New("deletion error")).
					Once()
			},
			wantErr: errors.New("repository.DeleteParticipantReplacementByParticipantID: deletion error"),
		},
		{
			name:          "existent replacement - creation successful",
			participantID: 3,
			exist:         true,
			person:        entity.Person{FirstName: "John", LastName: "Doe"},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreatePerson(ctx, entity.Person{FirstName: "John", LastName: "Doe"}).
					Return(entity.Person{ID: 4, FirstName: "John", LastName: "Doe"}, nil).
					Once()
				m.EXPECT().
					CreateParticipantReplacement(ctx, int64(3), int64(4)).
					Return(nil).
					Once()
			},
		},
		{
			name:          "existent replacement - create person error",
			participantID: 4,
			exist:         true,
			person:        entity.Person{FirstName: "Jane", LastName: "Doe"},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreatePerson(ctx, entity.Person{FirstName: "Jane", LastName: "Doe"}).
					Return(entity.Person{}, errors.New("create person error")).
					Once()
			},
			wantErr: errors.New("repository.CreatePerson: create person error"),
		},
		{
			name:          "existent replacement - create participant replacement error",
			participantID: 5,
			exist:         true,
			person:        entity.Person{FirstName: "Alice", LastName: "Smith"},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreatePerson(ctx, entity.Person{FirstName: "Alice", LastName: "Smith"}).
					Return(entity.Person{ID: 4, FirstName: "Alice", LastName: "Smith"}, nil).
					Once()
				m.EXPECT().
					CreateParticipantReplacement(ctx, int64(5), int64(4)).
					Return(errors.New("create participant replacement error")).
					Once()
			},
			wantErr: errors.New("repository.CreateParticipantReplacement: create participant replacement error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			err := NewService(repositoryMock).UpdateParticipantReplacementByParticipantID(ctx, test.participantID, test.exist, test.person)

			if test.wantErr != nil {
				require.Error(t, err)
				assert.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
