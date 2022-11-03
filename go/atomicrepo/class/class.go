package class

import "context"

type Class struct{}

type Repository interface {
	GetClassByCourseID(ctx context.Context, courseID int) (*Class, error)
}
