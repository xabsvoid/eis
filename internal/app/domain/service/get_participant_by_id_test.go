package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
	"github.com/xabsvoid/eis/internal/pkg/mock"
)

func TestService_GetParticipantByID(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	type testCase struct {
		name            string
		id              int64
		setupMock       func(m *mock.MockRepository)
		wantParticipant entity.Participant
		wantPersons     []entity.Person
		wantErr         error
	}

	tests := []testCase{
		{
			name: "successfully retrieve participant and related persons",
			id:   1,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetParticipantByID(ctx, int64(1)).
					Return(entity.Participant{ID: 1, EventID: 23, PersonID: 11}, nil).
					Once()
				m.EXPECT().
					GetPersonsByIDs(ctx, []int64{11}).
					Return([]entity.Person{{ID: 11, FirstName: "John"}}, nil).
					Once()
			},
			wantParticipant: entity.Participant{ID: 1, EventID: 23, PersonID: 11},
			wantPersons:     []entity.Person{{ID: 11, FirstName: "John"}},
		},
		{
			name: "repository error fetching participant",
			id:   2,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetParticipantByID(ctx, int64(2)).
					Return(entity.Participant{}, errors.New("repository error")).
					Once()
			},
			wantErr: errors.New("repository.GetParticipantByID: repository error"),
		},
		{
			name: "repository error fetching persons",
			id:   3,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetParticipantByID(ctx, int64(3)).
					Return(entity.Participant{ID: 3, EventID: 23, PersonID: 11}, nil).
					Once()
				m.EXPECT().
					GetPersonsByIDs(ctx, []int64{11}).
					Return(nil, errors.New("repository persons error")).
					Once()
			},
			wantErr: errors.New("repository.GetPersonsByIDs: repository persons error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			participant, persons, err := NewService(repositoryMock).GetParticipantByID(ctx, test.id)

			if test.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, test.wantParticipant, participant)
			assert.Equal(t, test.wantPersons, persons)
		})
	}
}
