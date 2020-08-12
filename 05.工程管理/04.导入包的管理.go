package main

import (
	"fmt"
	"os"
)

func main() {

	/*
	通过import的方式导入，导入的包必须使用，否则会报错
		import (
			"fmt"
			"os"
		)

	给包起别名:
		import io "fmt"

	忽略包：前面加_
		import _ "fmt"

	 */
	 fmt.Println(os.Args)

}
