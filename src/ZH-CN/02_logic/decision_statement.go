package main;

import (
	"fmt"
);

func main() {
	var number int64;
	fmt.Println("请输入数字：");
	fmt.Scan(&number);

	//#region
	// if + (else if) + (else) 判断逻辑，其中 else if/else 可以省略
	if number % 3 == 0 {
		fmt.Println("这个数能被3整除");
	} else if number % 3 == 1 {
		fmt.Println("这个数被3除余1");
	} else {
		fmt.Println("这个数被3除余2");
	}
	//#endregion

	// switch case 判断
	switch {
		case number >=0 && number <=100: 
			fmt.Println("这个数在0～100之间");
			break;
		case number < 0:
			fmt.Println("这个数小于0");
			break;
		case number > 0:
			fmt.Println("这个数大于100");
			break;
	}
}
