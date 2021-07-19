package service

import (
	"github.com/jinzhu/gorm"
	data "github.com/xdu31/exercise-service/pkg/server/db"
	"github.com/xdu31/exercise-service/pkg/storage"
)

type ExerciseService interface {
	GetExercisesForMuscle(muscle string) ([]string, error)
	GetMusclesForExercise(exercise string) ([]string, error)
}

type exerciseService struct {
	ai uint32
}

func NewExerciseService(ai uint32) ExerciseService {
	return &exerciseService{ai: ai}
}

/*
 * @Description : GetExercisesForMuscle function returns all Exercise objects for a muscle group
 * account given in AuthInfo.
 */
func (s *exerciseService) GetExercisesForMuscle(muscle string) ([]string, error) {
	var exercises []string

	if err := storage.WithTx(func(tx *gorm.DB) (err error) {
		exercises, err = data.NewExerciseData(tx).GetExercisesForMuscle(muscle)
		return err
	}); err != nil {
		return nil, err
	}

	return exercises, nil
}

/*
 * @Description : GetMusclesForExercise function returns all Muscle object for an exercise
 * for an account given in AuthInfo.
 */
func (s *exerciseService) GetMusclesForExercise(exercise string) ([]string, error) {
	var muscles []string

	if err := storage.WithTx(func(tx *gorm.DB) (err error) {
		muscles, err = data.NewExerciseData(tx).GetMusclesForExercise(exercise)
		return err
	}); err != nil {
		return nil, err
	}

	return muscles, nil
}
