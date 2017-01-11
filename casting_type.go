package main
import "fmt"
import "strconv"
import "reflect"

func main() {
	//Int to string and concat
	age_int := 26
	age_to_string := strconv.Itoa(age_int)
	fmt.Println("My age is " + age_to_string)
	//String to int and concat
	age_string := "25"
	age_to_int,_ := strconv.Atoi(age_string)
	fmt.Println(age_to_int + 1)
	//ParseBool,ParseFloat,ParseInt,ParseUint,etc...
	bool_string := "true"
	bool_to_bool,_ := strconv.ParseBool(bool_string)
	fmt.Println(bool_to_bool)
	fmt.Println(reflect.TypeOf(bool_to_bool))
}