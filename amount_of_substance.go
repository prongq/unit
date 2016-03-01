package unit

// https://en.wikipedia.org/wiki/Amount_of_substance
// https://en.wikipedia.org/wiki/Mole_(unit)

// AmountOfSubstance represents a SI unit of amount of substance (in mole, mol)
type AmountOfSubstance float64

// ...
const (
	Mole AmountOfSubstance = 1e0 // SI
)

// Moles returns the amount of substance in mol
func (a AmountOfSubstance) Moles() float64 {
	return float64(a)
}
