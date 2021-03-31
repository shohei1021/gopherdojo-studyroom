package convert

type Convert struct {
	btype     string
	atype     string
	directory string
}

func newConvertHandler(btype string, atype string, directory string) (*Convert, error) {
	return &Convert{
		btype:     btype,
		atype:     atype,
		directory: directory,
	}, nil
}
