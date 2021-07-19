package svc

import (
	"context"

	"github.com/xdu31/exercise-service/pkg/server/pb"
	"github.com/xdu31/exercise-service/pkg/server/service"
	"github.com/xdu31/exercise-service/pkg/util"
)

func newExerciseService(ctx context.Context) (service.ExerciseService, error) {
	ai, err := util.GetAccountID(ctx)
	if err != nil {
		return nil, err
	}
	return service.NewExerciseService(ai), nil
}

/*
 * @Description : GetExercisesForMuscle function implements retrieval of exercises
 */
func (s *exerciseServer) GetExercisesForMuscle(ctx context.Context, req *pb.GetExercisesForMuscleReq) (*pb.GetExercisesForMuscleResp, error) {
	srv, err := newExerciseService(ctx)
	if err != nil {
		return nil, util.ErrStatus(err)
	}

	exerciseList, err := srv.GetExercisesForMuscle(req.MuscleName)
	if err != nil {
		return nil, util.ErrStatus(err)
	}

	return &pb.GetExercisesForMuscleResp{Results: exerciseList}, util.SetOK(ctx, "Found %d Exercise(s)", len(exerciseList))
}

/*
 * @Description : GetMusclesForExercise function implements retrieval of muscles
 */
func (s *exerciseServer) GetMusclesForExercise(ctx context.Context, req *pb.GetMusclesForExerciseReq) (*pb.GetMusclesForExerciseResp, error) {
	srv, err := newExerciseService(ctx)
	if err != nil {
		return nil, util.ErrStatus(err)
	}

	muscleList, err := srv.GetMusclesForExercise(req.ExerciseName)
	if err != nil {
		return nil, util.ErrStatus(err)
	}

	return &pb.GetMusclesForExerciseResp{Results: muscleList}, util.SetOK(ctx, "Found %d Muscle(s)", len(muscleList))
}
