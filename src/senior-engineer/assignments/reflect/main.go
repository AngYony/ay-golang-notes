package main

import (
	"fmt"
	"reflect"
	"strings"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	if t.Kind() == reflect.Struct {
		tableName := t.Name()
		// 字符串拼接
		var cnstr, cvstr strings.Builder
		for i := 0; i < t.NumField(); i++ {
			cn := t.Field(i).Name // 获取列名
			cv := v.Field(i)
			// 拼接列
			_, _ = fmt.Fprintf(&cnstr, "%s,", cn)
			switch cv.Kind() {
			case reflect.Int:
				_, _ = fmt.Fprintf(&cvstr, "%d,", cv)
			case reflect.String:
				_, _ = fmt.Fprintf(&cvstr, "'%s',", cv)
			default:
				_, _ = fmt.Fprintf(&cvstr, "'%s',", cv)
			}
			// fmt.Println(f.Name, f.Type, e)
		}
		fmt.Printf("insert into %s (%s) values (%s);",
			tableName,
			strings.TrimRight((cnstr).String(), ","),
			strings.TrimRight(cvstr.String(), ","))
		fmt.Println()
	} else {
		fmt.Println("unsupported type")
	}
}
