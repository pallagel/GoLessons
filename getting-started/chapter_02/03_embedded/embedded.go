package main

import "fmt"

type sender struct {
	rateLimit int
	user      //an embedded structs
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("Sender name: ", s.name)
	fmt.Println("Sender number: ", s.number)
	fmt.Println("sender rate limit: ", s.rateLimit)
	fmt.Println("************************************")
}

func main() {
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Dan Shaw",
			number: 18004443334,
		},
	})

	test(sender{
		rateLimit: 200,
		user: user{
			name:   "Peter Pan",
			number: 18883334435,
		},
	})

	test(sender{
		rateLimit: 200000,
		user: user{
			name:   "Donald Duck",
			number: 734567898,
		},
	})
}
