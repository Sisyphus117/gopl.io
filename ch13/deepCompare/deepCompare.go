package deepcompare

import "reflect"

func Deepcompare(x, y interface{}) bool {
	w := reflect.ValueOf(x)
	v := reflect.ValueOf(y)
	fw, isNumW := toFloat64(w)
	fv, isNumV := toFloat64(v)
	if !isNumW || !isNumV {
		return false
	}
	return fw >= -1e9 && fw <= 1e9 && fv >= -1e9 && fv <= 1e9 && fw == fv
}

func toFloat64(v reflect.Value) (float64, bool) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Float32, reflect.Float64:
		return float64(v.Float()), true
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint8, reflect.Uint64:
		return float64(v.Uint()), true
	default:
		return 0, false
	}
}
