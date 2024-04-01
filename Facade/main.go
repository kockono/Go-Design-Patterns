package facade

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height,
		make([]rune, width*height)}
}

func NewViewport(buffer *Buffer) *ViewPort {
	return &ViewPort{buffer, 0}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

func (v ViewPort) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

type Console struct {
	buffers   []*Buffer
	viewports []*ViewPort
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{
		[]*Buffer{b},
		[]*ViewPort{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

type ViewPort struct {
	buffer *Buffer
	offset int
}

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
	_ = u
}
