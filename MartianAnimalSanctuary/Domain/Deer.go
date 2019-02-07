package ark

import "math/rand"

//Deer animal
type Deer struct {
	Name string
}

func (l Deer) String() string {
	return l.Name
}

//Move animal
func (l Deer) Move() string {
	switch rand.Intn(4) {
	case 0:
		return "running"
	case 1:
		return "walking"
	case 2:
		return "jumping"
	default:
		return "resting"
	}
}

//Eat something
func (l Deer) Eat() string {
	switch rand.Intn(6) {
	case 0:
		return "grass"
	case 1:
		return "apple"
	case 2:
		return "carrot"
	case 3:
		return "root"
	case 4:
		return "corn"
	default:
		return "nothing"
	}
}
