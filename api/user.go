package api

import (
	"golang_project/api/definition"
	"golang_project/entity"
	"golang_project/serverconst"

	"gopkg.in/guregu/null.v4"
)

type FilterParams struct {
	Limit      int64
	Offset     int64
	SortCol    string
	SortOrder  string
	SearchTerm string
	Active     bool
}

type UserStorage interface {
	GetPerson(personId int64) (*entity.Person, error)
	New(user *entity.Person) (int64, error)
}

type userApi struct {
	UserStorage UserStorage
}

func NewUserApi(userStorage UserStorage) *userApi {
	return &userApi{
		UserStorage: userStorage,
	}
}

func (u *userApi) GetPersonInfo(id int64) definition.GetPersonInfoResponse {
	user, err := u.UserStorage.GetPerson(id)

	if err != nil {
		return definition.GetPersonInfoResponse{
			Error: definition.ErrorMsg{
				Message: null.StringFrom(err.Error()),
			},
		}
	}

	formatUser := formatEntityToPersonDefinition(user)

	return definition.GetPersonInfoResponse{
		Person: *formatUser,
	}
}

func formatEntityToPersonDefinition(e *entity.Person) *definition.Person {
	formattedResult := definition.Person{
		Name:        e.Name,
		PhoneNumber: e.PhoneNumber,
		City:        e.City,
		State:       e.State,
		Street1:     e.Street1,
		Street2:     e.Street2,
		ZipCode:     e.ZipCode,
	}
	return &formattedResult
}

func (u *userApi) CreatePerson(req definition.CreatePersonRequest) definition.CreatePersonResponse {

	fomattedUserEntity := &entity.Person{
		Name:        req.Name,
		City:        req.City,
		State:       req.State,
		Street1:     req.Street1,
		Street2:     req.Street2,
		ZipCode:     req.ZipCode,
		PhoneNumber: req.PhoneNumber,
	}

	userId, err := u.UserStorage.New(fomattedUserEntity)
	if err != nil {
		return definition.CreatePersonResponse{
			Error: definition.ErrorMsg{
				Message: null.StringFrom(err.Error()),
			},
		}
	}

	return definition.CreatePersonResponse{
		Success: definition.SuccessMsg{
			Message: null.StringFrom(serverconst.UserCreated),
		},
		PersonID: userId,
	}
}
