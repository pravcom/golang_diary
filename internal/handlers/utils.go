package handlers

type SaveErr string

const (
	ErrLoginIsEmpty     = SaveErr("fill: Login")
	ErrDateIsEmpty      = SaveErr("fill: Date")
	ErrProjectIsEmpty   = SaveErr("fill: Project")
	ErrTaskIsEmpty      = SaveErr("fill: Task")
	ErrTimeHoursIsEmpty = SaveErr("fill: Time")
	ErrBadMethod        = SaveErr("Method %s is no available")
)

func (e SaveErr) Error() string {
	return string(e)
}
