package meimap

// StringPickByKeys | pick the key you want from the intput map according to the checkeys slice
func StringPickByKeys(intput map[string]interface{}, checkeys []string) map[string]interface{} {
	if checkeys == nil || len(checkeys) == 0 {
		return intput
	}
	output := map[string]interface{}{}
	for _, key := range checkeys {
		if val, ok := intput[key]; ok {
			output[key] = val
		}
	}
	return output
}

// StringDelByKeys | del one or many key from map
func StringDelByKeys(intput map[string]interface{}, dels []string) map[string]interface{} {
	if dels == nil || len(dels) == 0 {
		return intput
	}
	for _, key := range dels {
		delete(intput, key)
	}
	return intput
}
