package dummy

import (
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/mocks"
	"testing"
)

func TestNewService(t *testing.T) {
	NewService(
		new(mocks.Logger),
		new(mocks.Repository),
		new(mocks.IdCreator),
	)
	t.Log("Service created successfully")
}

func TestExecute(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		logger := new(mocks.Logger)
		logger.Mock.On("Debug", "Creating new dummy").Return("mocado")
		repository := new(mocks.Repository)
		repository.On("Create", &dummy.Dummy{
			DummyId:   "123ABC",
			DummyName: "samuca",
		}).Return(nil)
		creator := new(mocks.IdCreator)
		creator.On("Create").Return("123ABC")
		service := NewService(logger, repository, creator)
		valid := new(mocks.Validator)
		valid.On("ValidateStruct", &Data{
			DummyId:   "",
			DummyName: "samuca",
		}).Return(nil)
		req := NewRequest(
			&Data{
				DummyId:   "",
				DummyName: "samuca",
			},
			valid,
		)
		service.Execute(req)
		t.Log("Service executed successfully")
	})

	//t.Run("Error", func(t *testing.T) {
	//	logger := new(mocks.Logger)
	//	logger.Mock.On("Debug", "Creating new dummy").Return("mocado")
	//	repository := new(mocks.Repository)
	//	repository.On("Create", &dummy.Dummy{
	//		DummyId:   "123ABC",
	//		DummyName: "samuca",
	//	}).Return(nil)
	//	creator := new(mocks.IdCreator)
	//	creator.On("Create").Return("123ABC")
	//	service := NewService(logger, repository, creator)
	//	valid := new(mocks.Validator)
	//	valid.On("ValidateStruct", &Data{
	//		DummyId:   "",
	//		DummyName: "samuca",
	//	}).Return()
	//	req := NewRequest(
	//		&Data{
	//			DummyId:   "",
	//			DummyName: "samuca",
	//		},
	//		valid,
	//	)
	//	service.Execute(req)
	//	t.Log("Service executed successfully")
	//})
}
