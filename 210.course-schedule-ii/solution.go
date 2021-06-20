package _10_course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses < 0 {
		return []int{}
	}

	if len(prerequisites) == 0 {
		var defaul []int
		for i := numCourses - 1; i >= 0; i-- {
			defaul = append(defaul, i)
		}
		return defaul
	}

	var prereqsMap = make(map[int][]int)
	for _, prerequisite := range prerequisites {
		prereqsMap[prerequisite[0]] = append(prereqsMap[prerequisite[0]], prerequisite[1])
	}

	var circle bool
	seen := make(map[int]bool)
	flag := make(map[int]bool)
	var checkCircle func(items []int)
	var order []int

	checkCircle = func(items []int) {
		for _, course := range items {
			if has, ok := flag[course]; ok && !has {
				circle = true
			}

			if !seen[course] {
				seen[course] = true

				flag[course] = false
				checkCircle(prereqsMap[course])
				flag[course] = true

				order = append(order, course)
			}
		}

	}

	var keys []int
	for key := range prereqsMap {
		keys = append(keys, key)
	}
	checkCircle(keys)

	if circle {
		order = []int{}
	} else {
		if numCourses != len(order) {
			var tmp []int
			for i := numCourses - 1; i >= 0; i-- {
				var has = false
				for _, c := range order {
					if c == i {
						has = true
						break
					}
				}
				if !has {
					tmp = append(tmp, i)
				}
				if len(tmp)+len(order) == numCourses {
					break
				}
			}
			tmp = append(tmp, order...)
			return tmp
		}
	}

	return order

}
