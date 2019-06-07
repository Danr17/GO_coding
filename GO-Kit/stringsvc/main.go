
import "context"

//StringService provides operations on strings
type StringService interface{
	Uppercase(string) (string, error)
	Count(string) int
}