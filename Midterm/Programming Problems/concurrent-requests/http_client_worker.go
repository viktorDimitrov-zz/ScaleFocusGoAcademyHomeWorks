package httpClientWorker

type Worker struct {
	serverURL string
}

// Filter the numbers using the rule
func FilterNumbers(data []int) []int {
	panic("Please implement the FilterNumbers function")
}

func NewWorker(serverUrl string) *Worker {
	panic("Please implement the New function")
}

// Get the data from the server
func (w Worker) GetNumbers() []int {
	panic("Please implement the GetNumbers function")
}

// Get, Filter and Submit the "special" numbers
func (w Worker) ProcessAndSubmit() {
	panic("Please implement the ProcessAndSubmit function")
}
