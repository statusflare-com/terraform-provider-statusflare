package terraform

// converting array of unknown types into array
// of strings
//
// Terraform is operating with lot of unknown types/ interfaces.
// It's easy to cast the primitive types but terraform
// is using []interface{} for any TypeList. We need quite often
// convert this array of interfaces into array of strings.
func toStrArray(arr []interface{}) []string {
	var strArr []string
	for _, val := range arr {
		strArr = append(strArr, val.(string))
	}
	return strArr
}
