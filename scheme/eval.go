package scheme

// This function takes a program by string and interprets it in the top-level environment.
func Eval(program string) (Type, error) {
	return NewExpression(program).Eval()
}