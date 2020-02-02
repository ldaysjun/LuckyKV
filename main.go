package main

import "fmt"

func main() {
	//r := redis.NewClient(nil)
	//r.Get()
	//r.Set()
	//t := cmd.NewClient()
	//t.Get("test")

	test := make(map[string]string)
	test["ok"] = "ok1"
	delete(test, "ok")

	if _, ok := test["ok"]; ok {
		fmt.Println(test["ok"])
	} else {
		fmt.Println("不存在")
	}

	fmt.Println(len(test))

	fmt.Println(test)

}

type name struct {
	s string
}
