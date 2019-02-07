package ark

import "math/rand"

//Lion animal
type Lion struct {
	Name string
}

func (l Lion) String() string {
	return l.Name
}

//Move animal
func (l Lion) Move() string {
	switch rand.Intn(3) {
	case 0:
		return "running"
	case 1:
		return "hunting"
	default:
		return "resting"
	}
}

//Eat something
func (l Lion) Eat() string {
	switch rand.Intn(6) {
	case 0:
		return "deer"
	case 1:
		return "antilope"
	case 2:
		return "rabbit"
	case 3:
		return "human"
	case 4:
		return "bufallo"
	default:
		return "nothing"
	}
}
