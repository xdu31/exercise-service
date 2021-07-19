package db

import (
	"github.com/jinzhu/gorm"
	"github.com/xdu31/exercise-service/pkg/util"
)

const (
	getMusclesSql   = "SELECT m.name FROM training_data AS t LEFT JOIN exercises AS e ON t.exercise_id = e.id LEFT JOIN muscle_groups AS m ON t.muscle_group_id = m.id WHERE e.name = ?"
	getExercisesSql = "SELECT e.name FROM training_data AS t LEFT JOIN exercises AS e ON t.exercise_id = e.id LEFT JOIN muscle_groups AS m ON t.muscle_group_id = m.id WHERE m.name = ?"
)

type ExerciseData interface {
	GetExercisesForMuscle(muscle string) ([]string, error)
	GetMusclesForExercise(exercise string) ([]string, error)
}

type exerciseData struct {
	db *gorm.DB
}

func NewExerciseData(db *gorm.DB) ExerciseData {
	return &exerciseData{db: db}
}

/*
 * @Description : GetExercisesForMuscle function returns list of Exercise objects.
 */
func (d *exerciseData) GetExercisesForMuscle(muscle string) ([]string, error) {
	var (
		exercises []string
	)

	exerciseNameRows, err := d.db.Raw(getExercisesSql, muscle).Rows()
	if err != nil {
		return nil, err
	}
	defer exerciseNameRows.Close()

	for exerciseNameRows.Next() {
		var exerciseName string
		exerciseNameRows.Scan(&exerciseName)
		exercises = append(exercises, exerciseName)
	}

	if len(exercises) == 0 {
		return nil, util.NotFoundErr("Muscle group %s not found", muscle)
	}

	return exercises, nil
}

/*
 * @Description : Create function creates a new Ip object.
 */
func (d *exerciseData) GetMusclesForExercise(exercise string) ([]string, error) {
	var (
		muscleGroups []string
	)

	muscleNameRows, err := d.db.Raw(getMusclesSql, exercise).Rows()
	if err != nil {
		return nil, err
	}
	defer muscleNameRows.Close()

	for muscleNameRows.Next() {
		var muscleName string
		muscleNameRows.Scan(&muscleName)
		muscleGroups = append(muscleGroups, muscleName)
	}

	if len(muscleGroups) == 0 {
		return nil, util.NotFoundErr("Exercise %s not found", exercise)
	}

	return muscleGroups, nil
}
