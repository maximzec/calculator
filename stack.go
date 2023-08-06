package main

//Структура стекового байта
//В дальнейшем она потребуется для реализации перевода строки в польскую запись
//И использовании в расчете

type ByteStack []byte

func (bs ByteStack) empty() bool {
	return true
}

func (bs ByteStack) push(b byte) bool {
	return true
}

func (bs ByteStack) pop() byte {
	return 0
}
