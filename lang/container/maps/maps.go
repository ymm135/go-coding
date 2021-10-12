package main

import "fmt"

func main() {
	//map[course:golang name:ccmouse quality:notbad site:imooc]
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	// map[]
	m2 := make(map[string]int) // m2 == empty map 
    // map[]
	var m3 map[string]int // m3 == nil

	fmt.Println("m, m2, m3:")
	fmt.Println(m, m2, m3)

	//map遍历
	fmt.Println("Traversing map m")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(`m["course"] =`, courseName)
	// 可以有两个返回值
	if causeName, ok := m["cause"]; ok { // key不存在, 返回空值
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
}
