package main

import "fmt"

func main() {
	/*
	switch 变量/表达式 {
		case val1:
			...
		case val2:
			...
		...
		default :
			...
	fallthrough:不跳出switch语句，只要满足switch语句中的一个，后面的无条件执行，
		比如是周四，会把周四到周日的都打印出来

	注意：
		case后面可以跟条件
			如：case > 100:

	}
	 */

	a := 4
	switch a {
	case 1:
		fmt.Println("周一")
		break
	case 2:
		fmt.Println("周二")
		break
	case 3:
		fmt.Println("周三")
		break
	case 4:
		fmt.Println("周四")
		break
	case 5:
		fmt.Println("周五")
		break
	case 6:
		fmt.Println("周六")
		break
	case 7:
		fmt.Println("周日")
		break
	default:
		fmt.Println("无效的日期")
		break







	}
}
