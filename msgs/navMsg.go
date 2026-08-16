package msgs 

type NavMsg struct {
	// Destination int
	// Source int
	Delta int 
}

func Forward() int {
	return NavMsg{
		Delta: 1, 
	}
}

func Back() int {
	return NavMsg{
		Delta: -1, 
	}
}

// func (m NavMsg) Content() int {
	// return NavMsg.Content 
// }








































