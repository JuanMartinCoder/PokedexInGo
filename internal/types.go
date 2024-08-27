package internal

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
