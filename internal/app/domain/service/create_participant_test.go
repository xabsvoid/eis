package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabsvoid/eis/internal/app/domain/model/entity"
	"github.com/xabsvoid/eis/internal/pkg/mock"
)

func TestService_CreateParticipant(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	type testCase struct {
		name            string
		eventID         int64
		person          entity.Person
		setupMock       func(m *mock.MockRepository)
		wantParticipant entity.Participant
		wantPerson      entity.Person
		wantErr         error
	}

	tests := []testCase{
		{
			name:    "success",
			eventID: 2,
			person:  entity.Person{FirstName: "John"},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreatePerson(ctx, entity.Person{FirstName: "John"}).
					Return(entity.Person{ID: 1, FirstName: "John"}, nil).
					Once()
				m.EXPECT().
					CreateParticipant(ctx, entity.Participant{EventID: 2, PersonID: 1}).
					Return(entity.Participant{ID: 3, EventID: 2, PersonID: 1}, nil).
					Once()
			},
			wantParticipant: entity.Participant{ID: 3, EventID: 2, PersonID: 1},
			wantPerson:      entity.Person{ID: 1, FirstName: "John"},
		},
		{
			name:    "error creating person",
			eventID: 2,
			person:  entity.Person{FirstName: "John"},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreatePerson(ctx, entity.Person{FirstName: "John"}).
					Return(entity.Person{}, errors.New("failed to create person")).
					Once()
			},
			wantErr: errors.New("repository.CreatePerson: failed to create person"),
		},
		{
			name:    "error creating participant",
			eventID: 2,
			person:  entity.Person{FirstName: "John"},
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					CreatePerson(ctx, entity.Person{FirstName: "John"}).
					Return(entity.Person{ID: 1, FirstName: "John"}, nil).
					Once()
				m.EXPECT().
					CreateParticipant(ctx, entity.Participant{EventID: 2, PersonID: 1}).
					Return(entity.Participant{}, errors.New("failed to create participant")).
					Once()
			},
			wantErr: errors.New("repository.CreateParticipant: failed to create participant"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			repositoryMock := mock.NewMockRepository(t)
			defer repositoryMock.AssertExpectations(t)

			test.setupMock(repositoryMock)

			participant, person, err := NewService(repositoryMock).CreateParticipant(ctx, test.eventID, test.person)

			if test.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantErr.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, test.wantParticipant, participant)
			assert.Equal(t, test.wantPerson, person)
		})
	}
}
