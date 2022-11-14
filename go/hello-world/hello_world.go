package greeting
import "fmt"

// HelloWorld greets the world.
func HelloWorld() string {
	return "Hello, World!" 
}

func main() {
    fmt.Println(HelloWorld())
}
