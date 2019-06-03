package main

import (
	"context"
	"mathOp_service/service"
	"testing"
)

func TestAddition(t *testing.T) {
	s := server{}
	tests := []struct {
		X      int64
		Y      int64
		Answer int64
	}{
		{
			X:      1,
			Y:      2,
			Answer: 3,
		},
		{
			X:      1000,
			Y:      1000,
			Answer: 2000,
		},
		{
			X:      -10,
			Y:      10,
			Answer: 0,
		},
	}

	for _, test := range tests {
		req := &service.AddRequest{X: int64(test.X), Y: int64(test.Y)}
		res, _ := s.Add(context.Background(), req)
		if res.Answer != test.Answer {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", test.X, test.Y, res.Answer, test.Answer)
		}
	}
}

func TestSqrt(t *testing.T) {
	s := server{}
	tests := []struct {
		X      float64
		Answer float64
	}{
		{
			X:      16,
			Answer: 4,
		},
		{
			X:      64,
			Answer: 8,
		},
	}

	for _, test := range tests {
		req := &service.SqrtRequest{X: float64(test.X)}
		res, _ := s.Sqrt(context.Background(), req)
		if res.Answer != test.Answer {
			t.Errorf("Sqrt of (%f) was incorrect, got: %f, want: %f.", test.X, res.Answer, test.Answer)
		}
	}
}
