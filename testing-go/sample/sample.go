package sample

type Sample struct {
	FirstName string
	LastName  string
}

func PublicFunc() Sample {
	return Sample{
		FirstName: "first",
		LastName:  "last",
	}
}

func privateFunc() Sample {
	return Sample{
		FirstName: "first",
		LastName:  "last",
	}

}
