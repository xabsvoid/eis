package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
	"github.com/xabsvoid/eis/internal/pkg/mock"
)

func TestService_GetParticipantsByEventID(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	type testCase struct {
		name             string
		eventID          int64
		setupMock        func(m *mock.MockRepository)
		wantParticipants []entity.Participant
		wantPersons      []entity.Person
		wantErr          error
	}

	tests := []testCase{
		{
			name:    "successfully retrieve participants and related persons",
			eventID: 23,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetParticipantsByEventID(ctx, int64(23)).
					Return([]entity.Participant{{ID: 1, EventID: 23, PersonID: 11}, {ID: 2, EventID: 23, PersonID: 12}, {ID: 3, EventID: 23, PersonID: 12}}, nil).
					Once()
				m.EXPECT().
					GetPersonsByIDs(ctx, []int64{11, 12, 12}).
					Return([]entity.Person{{ID: 11, FirstName: "John"}, {ID: 12, FirstName: "Kate"}}, nil).
					Once()
			},
			wantParticipants: []entity.Participant{{ID: 1, EventID: 23, PersonID: 11}, {ID: 2, EventID: 23, PersonID: 12}, {ID: 3, EventID: 23, PersonID: 12}},
			wantPersons:      []entity.Person{{ID: 11, FirstName: "John"}, {ID: 12, FirstName: "Kate"}},
		},
		{
			name:    "repository error fetching participants",
			eventID: 2,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetParticipantsByEventID(ctx, int64(2)).
					Return(nil, errors.New("repository error")).
					Once()
			},
			wantErr: errors.New("repository.GetParticipantsByEventID: repository error"),
		},
		{
			name:    "repository error fetching persons",
			eventID: 543,
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					GetParticipantsByEventID(ctx, int64(543)).
					Return([]entity.Participant{{ID: 1, EventID: 543, PersonID: 11}, {ID: 2, EventID: 543, PersonID: 12}}, nil).
					Once()
				m.EXPECT().
					GetPersonsByIDs(ctx, []int64{11, 12}).
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

			participants, persons, err := NewService(repositoryMock).GetParticipantsByEventID(ctx, test.eventID)

			if test.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, test.wantParticipants, participants)
			assert.Equal(t, test.wantPersons, persons)
		})
	}
}
