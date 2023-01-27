package solver

const (
	ErrTimeout     = SolvingError("timeout. Failed to solve the grid")
	ErrNoSolutions = SolvingError("could not find a solution to a grid")
)

type SolvingError string

func (e SolvingError) Error() string {
	return string(e)
}
