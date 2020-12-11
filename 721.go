package leetcode

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	roots := map[string]string{}      // email -> email
	email2user := map[string]string{} // email -> user name

	var root func(aim string) string
	root = func(aim string) string {
		if roots[aim] == aim {
			return aim
		}

		return root(roots[aim])
	}

	for _, account := range accounts {

		user := account[0]
		emails := account[1:]
		for _, email := range emails {
			if roots[email] == "" {
				email2user[email] = user
				roots[email] = email
			}

			x := root(email)
			y := root(emails[0])
			if x == email {
				roots[email] = y
			} else if x != y {
				roots[x] = y
			}
		}

	}

	//for key, value := range roots {
	//
	//	fmt.Printf("%v -> %v\n", key, value)
	//}

	rv := [][]string{}

	for _, account := range accounts {
		user := account[0]
		element := []string{}

		if email2user[account[1]] == "" {
			continue
		}

		rootOfAccount := root(account[1])

		for email, _ := range email2user {
			if root(email) == rootOfAccount {
				element = append(element, email)
				delete(email2user, email)
			}
		}

		sort.Slice(element, func(i, j int) bool {
			return element[i] < element[j]
		})

		rv = append(rv, append([]string{user}, element...))
	}

	return rv
}
