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

func findOrder1(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

func findOrder2(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
