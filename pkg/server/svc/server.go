package svc

import (
	pb "github.com/xdu31/exercise-service/pkg/server/pb"
)

func NewExerciseServer() pb.ExerciseServiceServer {
	return &exerciseServer{}
}

type exerciseServer struct{}
