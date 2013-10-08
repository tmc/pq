package main

func init() {
	addTestCases(pgerror_removalTests, pgerror_removal)
}

var pgerror_removalTests = []testCase{
	{
		Name: "pgerror_removal.0",
		In: `package main

import "github.com/lib/pq"

func f() {
	conn, _ := sql.Open("postgres", "")
	_, e := conn.Exec("invalid")
	if pge, ok := e.(pq.PGError); ok {
		fmt.Println("pg error:", pge)
	} else {
		fmt.Println("non-pg error:", e)
	}
}
`,
		Out: `package main

import "github.com/lib/pq"

func f() {
	conn, _ := sql.Open("postgres", "")
	_, e := conn.Exec("invalid")
	if pge, ok := e.(*pq.Error); ok {
		fmt.Println("pg error:", pge)
	} else {
		fmt.Println("non-pg error:", e)
	}
}
`,
	},
}
