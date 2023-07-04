package enum

type Gender string

const (
	MaleGender   Gender = "MALE"
	FemaleGender Gender = "FEMALE"
	OthersGender Gender = "OTHERS"
)

func (g Gender) ToString() string {
	return string(g)
}
