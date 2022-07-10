package utils

import "math/big"

type textChunk struct {
	stringVal string
}

func bigIntToString(n *big.Int) string {
	big256 := big.NewInt(256)
	big0 := big.NewInt(0)
	temp, _ := new(big.Int).SetString(n.Text(10), 10)

	stringVal := ""

	if temp.Cmp(big0) == 0 {
		stringVal = "0"
	} else {
		for temp.Cmp(big0) > 0 {
			quot := big.NewInt(0)
			rem := big.NewInt(0)
			quot, rem = temp.DivMod(temp, big256, rem)

			charNum := rem.Int64()
			stringVal = stringVal + string(rune(charNum))
			temp = quot
		}
	}

	return stringVal
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewString(n string) textChunk {
	t := textChunk{n}
	return t
}

func (t textChunk) Text() string {
	return t.stringVal
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewBigInt(n *big.Int) textChunk {
	t := textChunk{bigIntToString(n)}

	return t
}

func (t textChunk) BigIntValue() *big.Int {
	big256 := big.NewInt(256)
	result := big.NewInt(0)

	runes := []rune(t.stringVal)

	for i := len(runes) - 1; i >= 0; i-- {
		result = result.Mul(result, big256)
		result = result.Add(result, big.NewInt(int64(runes[i])))
	}

	return result
}

func BlockSize(n big.Int) int {
	big1 := big.NewInt(1)
	big2 := big.NewInt(2)
	temp, _ := new(big.Int).SetString(n.Text(10), 10)

	blockSize := 0

	temp = temp.Sub(temp, big1)
	for temp.Cmp(big1) > 0 {
		temp = temp.Div(temp, big2)
		blockSize++
	}

	return blockSize / 8
}
