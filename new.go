package genid

import "github.com/general252/gen-id/generator"

// NewGeneratorData
func NewGeneratorData(isFullAge *bool) (ret *generator.GeneratorData) {
	var (
		data = new(generator.GeneratorData)
	)
	data.GeneratorBankID()
	data.GeneratorAddress()
	data.GeneratorEmail()
	data.GeneratorIDCart(isFullAge)
	data.GeneratorName()
	data.GeneratorPhone()
	ret = data
	return
}
