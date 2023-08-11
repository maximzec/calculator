package stack

//Структура стекового байта
//В дальнейшем она потребуется для реализации перевода строки в польскую запись
//И использовании в расчете

type ByteStack []byte

func createNewByteStack() ByteStack {
	bs := make(ByteStack, 0)
	return bs
}

func (bs ByteStack) empty() bool {
	return len(bs) == 0
}

func (bs ByteStack) push(b byte) bool {
	bs = append(bs, b)
	return true
}

func (bs ByteStack) pop() byte {
	b := bs[len(bs)-1]
	bs = bs[:len(bs)-1]
	return b
}

func (bs ByteStack) peek() byte {
	return bs[len(bs)-1]
}
