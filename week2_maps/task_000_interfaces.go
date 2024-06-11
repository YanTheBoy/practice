package week2_maps

type Formatter interface {
	Format() string
}

type PlainText struct {
	Text string
}

type BoldText struct {
	Text string
}

type CodeText struct {
	Text string
}

type ItalicText struct {
	Text string
}

func (p PlainText) Format() string {
	return p.Text
}

func (b BoldText) Format() string {
	return "**" + b.Text + "**"
}

func (c CodeText) Format() string {
	return "`" + c.Text + "`"
}

func (i ItalicText) Format() string {
	return "_" + i.Text + "_"
}

/*Message Formatter
Реализуйте интерфейс Formatter с методом Format, который возвращает
отформатированную строку.
Определите структуры, удовлетворяющие интерфейсу Formatter: обычный
текст(как есть), жирным шрифтом(** **), код(` `), курсив(_ _)
Опционально: иметь возможность задавать цепочку модификаторов
chainFormatter.AddFormatter(plainText)
chainFormatter.AddFormatter(bold)
chainFormatter.AddFormatter(code)
*/
