// Multiply multiplies two NumgoArrays element-wise or multiplies a NumgoArray with a scalar value
func (a *NumgoArray) Multiply(b interface{}) *NumgoArray {
	res := NewNumgoArray(a.Shape)
	switch b := b.(type) {
	case *NumgoArray:
		if !equalShapes(a.Shape, b.Shape) {
			panic("arrays have different shapes")
		}
		for i := 0; i < len(a.Data); i++ {
			res.Data[i] = a.Data[i] * b.Data[i]
		}
	case float64:
		for i := 0; i < len(a.Data); i++ {
			res.Data[i] = a.Data[i] * b
		}
	default:
		panic("invalid type")
	}
	return res
}
