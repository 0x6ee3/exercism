package strand

const testVersion = 3

// ToRNA ..
func ToRNA(dna string) string {
	var rna string
	for _, c := range dna {
		switch c {
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
