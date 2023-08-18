package stack

//Структура стекового байта
//В дальнейшем она потребуется для реализации перевода строки в польскую запись
//И использовании в расчете

//TODO: Переделать структуру хранения стэка
// - Это должна быть структура с внутренним стэком

type StringStack struct {
	stack []string
}

func CreateNewStringStack() *StringStack {
	bs := StringStack{
		stack: make([]string, 0),
	}
	return &bs
}

func (ss *StringStack) Empty() bool {
	return len(ss.stack) == 0
}

func (ss *StringStack) Push(s string) {
	// TODO: Нужно добавить проверку на пустоту стэка
	ss.stack = append(ss.stack, s)
}

func (ss *StringStack) Pop() string {
	// TODO: Нужно добавить проверку на пустоту стэка
	s := ss.stack[len(ss.stack)-1]
	ss.stack = ss.stack[:len(ss.stack)-1]
	return s
}

func (ss *StringStack) Peek() string {
	// TODO: Нужно добавить проверку на пустоту стэка
	return ss.stack[len(ss.stack)-1]
}
