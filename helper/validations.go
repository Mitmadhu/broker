package helper

import (
	"fmt"
)

func CheckEmpty(data string, errs *[]error, name string){
	if(len(data) == 0){
		*errs = append(*errs, fmt.Errorf("%s can not be empty", name))
	}		
}