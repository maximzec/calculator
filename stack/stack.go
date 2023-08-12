package stack

import "fmt"

//Структура стекового байта
//В дальнейшем она потребуется для реализации перевода строки в польскую запись
//И использовании в расчете

//TODO: Переделать структуру хранения стэка
// - Это должна быть структура с внутренним стэком
type ByteStack []byte

func CreateNewByteStack() ByteStack {
	bs := make(ByteStack, 0)
	return bs
}

func (bs ByteStack) Empty() bool {
	return len(bs) == 0
}

func (bs ByteStack) Push(b byte) ByteStack {
	return append(bs, b)
}

func (bs ByteStack) Pop() (byte, ByteStack) {
	b := bs[len(bs)-1]
	bs = bs[:len(bs)-1]
	return b, bs
}

func (bs ByteStack) Peek() byte {
	return bs[len(bs)-1]
}

func (bs ByteStack) Print() {
	for _, b := range bs {
		fmt.Println(b)
	}
}
