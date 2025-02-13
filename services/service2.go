package services

import(
	"context"

	// local packages
	model "products/models"
)

type AsampleRepository interface{
	CreateASampleByAppId(ctx context.Context,appID string, model *model.Asamplemodel) error 
	GetASampleByAppId(ctx context.Context, appID string, id string) (*model.Asamplemodel,error)
	DeleteASampleByAppId(ctx context.Context, appID string, id string) error
	UpdateASampleByAppId(ctx context.Context, appID string, model *model.Asamplemodel) error
}

type asample struct{
	repo AsampleRepository
}

func NewAsampleService(repo AsampleRepository) *asample{
	return &asample{repo: repo}
}

func (s *asample) CreateASampleByAppId(ctx context.Context,appID string, model *model.Asamplemodel) error  {
	return s.repo.CreateASampleByAppId(ctx,appID,model)
}

func (s *asample) GetASampleByAppId(ctx context.Context,appID string, id string) (*model.Asamplemodel,error){
	return s.repo.GetASampleByAppId(ctx,appID,id)
}

func (s *asample) UpdateASampleByAppId(ctx context.Context,appID string, model *model.Asamplemodel) error{
	return s.repo.UpdateASampleByAppId(ctx,appID,model)
}

func (s *asample) DeleteASampleByAppId(ctx context.Context,appID string, id string) error{
	return s.repo.DeleteASampleByAppId(ctx,appID,id)
}


	
