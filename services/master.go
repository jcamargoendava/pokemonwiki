package services

import (
	"context"

	masterModel "github.com/jcamargoendava/pokemonwiki/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MasterRepository interface {
	GetMaster(ctx context.Context, id string) (masterModel.Master, error)
	SaveMaster(ctx context.Context, mModel *masterModel.Master) (*mongo.InsertOneResult, error)
	UpdateMaster(ctx context.Context, id string, mModel *masterModel.Master) (*mongo.UpdateResult, error)
	DeleteMaster(ctx context.Context, id string) error
}

type Master struct {
	Repo MasterRepository
}

func NewMaster(repo MasterRepository) *Master {
	return &Master{
		Repo: repo,
	}
}

func (m *Master) GetMaster(ctx context.Context, id string) (masterModel.Master, error) {
	foundMaster, err := m.Repo.GetMaster(ctx, id)
	return foundMaster, err
}

func (m *Master) SaveMaster(ctx context.Context, master *masterModel.Master) (*mongo.InsertOneResult, error) {
	createdMaster, err := m.Repo.SaveMaster(ctx, master)
	return createdMaster, err
}

func (m *Master) UpdateMaster(ctx context.Context, id string, master *masterModel.Master) (*mongo.UpdateResult, error) {
	updatedMaster, err := m.Repo.UpdateMaster(ctx, id, master)
	return updatedMaster, err
}

func (m *Master) DeleteMaster(ctx context.Context, id string) error {
	err := m.Repo.DeleteMaster(ctx, id)
	return err
}
