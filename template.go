package nagos

type printer interface {
	open() string
	print() string
	close() string
}

type AbstractDisplay struct {
}

func (a *AbstractDisplay) Display(printer printer) string {
	result := printer.open()
	for i := 0; i < 5; i++ {
		result += printer.print()
	}
	result += printer.close()
	return result
}

type CharDisplay struct {
	*AbstractDisplay
	Char rune
}

func (c *CharDisplay) open() string {
	return "<<"
}

func (c *CharDisplay) print() string {
	return string(c.Char)
}

func (c *CharDisplay) close() string {
	return ">>"
}


