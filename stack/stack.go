package stack

//Структура стекового байта
//В дальнейшем она потребуется для реализации перевода строки в польскую запись
//И использовании в расчете

//TODO: Переделать структуру хранения стэка
// - Это должна быть структура с внутренним стэком

type ByteStack struct {
	stack []byte
}

func CreateNewByteStack() *ByteStack {
	bs := ByteStack{
		stack: make([]byte, 0),
	}
	return &bs
}

func (bs *ByteStack) Empty() bool {
	return len(bs.stack) == 0
}

func (bs *ByteStack) Push(b byte) {
	bs.stack = append(bs.stack, b)
}

func (bs *ByteStack) Pop() byte {
	b := bs.stack[len(bs.stack)-1]
	bs.stack = bs.stack[:len(bs.stack)-1]
	return b
}

func (bs *ByteStack) Peek() byte {
	return bs.stack[len(bs.stack)-1]
}
