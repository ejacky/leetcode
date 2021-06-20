package _07_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {

	var prereqsMap = make(map[int][]int)
	for _, prerequisite := range prerequisites {
		prereqsMap[prerequisite[0]] = append(prereqsMap[prerequisite[0]], prerequisite[1])
	}

	var circle bool
	seen := make(map[int]bool)
	flag := make(map[int]bool)
	var checkCircle func(items []int)

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
			}
		}

	}

	for course, _ := range prereqsMap {
		checkCircle([]int{course})
	}

	return !circle

}
