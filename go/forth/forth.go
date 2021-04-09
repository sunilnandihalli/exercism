package forth

import "strings"
import "strconv"
import "errors"

type Op interface {
	run(args []int) ([]int, error)
	numArgs() int
}

type IInt struct {
	val int
}

const (
	Add = iota
	Sub
	Mul
	Div
)

type IArithmetic struct{ op int }
type IDup struct{}
type ISwap struct{}
type IDrop struct{}
type IOver struct{}

func (x IInt) run(args []int) ([]int, error) {
	return []int{x.val}, nil
}
func (x IInt) numArgs() int {
	return 0
}

func (x IArithmetic) run(args []int) ([]int, error) {
	x1 := args[0]
	x2 := args[1]
	var ret int
	switch x.op {
	case Add:
		ret = x1 + x2
	case Sub:
		ret = x1 - x2
	case Mul:
		ret = x1 * x2
	case Div:
		if x2 == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		ret = x1 / x2
	}
	return []int{ret}, nil
}
func (x IArithmetic) numArgs() int {
	return 2
}

func (IDup) run(args []int) ([]int, error) {
	return []int{args[0], args[0]}, nil
}
func (IDup) numArgs() int {
	return 1
}

func (ISwap) run(args []int) ([]int, error) {
	return []int{args[1], args[0]}, nil
}
func (ISwap) numArgs() int {
	return 2
}

func (IDrop) run(_ []int) ([]int, error) {
	return []int{}, nil
}
func (IDrop) numArgs() int {
	return 1
}

func (IOver) run(args []int) ([]int, error) {
	return append(args, args[0]), nil
}
func (IOver) numArgs() int {
	return 2
}

func pop(stk []Op) (item Op, remainingStk []Op) {
	if len(stk) == 0 {
		remainingStk = stk
		item = nil
	} else {
		remainingStk = stk[:len(stk)-1]
		item = stk[len(stk)-1]
	}
	return
}

func simplify(ops []Op) ([]Op, error) {
	if len(ops) == 0 {
		return ops, nil
	}
	op, rops := pop(ops)
	rops, err := simplify(rops)
	if err != nil {
		return nil, err
	}
	n := op.numArgs()
	if len(rops) >= n {
		args := []int{}
		for _, x := range rops[len(rops)-n:] {
			i, ok := x.(IInt)
			if !ok {
				break
			}
			args = append(args, i.val)
		}
		if len(args) == n {
			rops = rops[:len(rops)-n]
			op_val, err := op.run(args)
			if err != nil {
				return nil, err
			}
			for _, x := range op_val {
				rops = append(rops, IInt{x})
			}
		} else {
			rops = append(rops, op)
		}
	} else {
		rops = append(rops, op)
	}
	return rops, nil
}

func eval(ops []Op) (ret []int, err error) {
	err = nil
	if len(ops) == 0 {
		ret = []int{}
		err = nil
		return
	}
	op, rops := pop(ops)
	rints, merr := eval(rops)
	if merr != nil {
		ret = nil
		err = merr
		return
	}
	if len(rints) < op.numArgs() {
		return nil, errors.New("insufficient arguments")
	}
	op_val, op_err := op.run(rints[len(rints)-op.numArgs():])
	if op_err != nil {
		ret = nil
		err = op_err
		return
	}
	ret = append(rints[:len(rints)-op.numArgs()],
		op_val...)
	return
}

func tokens(s string, custom_words map[string]([]Op)) (ops []Op, custom_word string, err error) {
	s = strings.Trim(s, " ")
	n := len(s)
	isCustomWord := false
	if s[0] == ':' && s[n-1] == ';' {
		isCustomWord = true
		s = s[1 : n-1]
	}
	idx := 0
	for _, x := range strings.Split(s, " ") {
		if len(x) == 0 {
			continue
		}
		idx += 1
		v, cerr := strconv.Atoi(x)
		if isCustomWord && idx == 1 {
			if cerr == nil {
				err = errors.New("attempting to redefine an integer")
				return
			}
			custom_word = strings.ToLower(x)
		} else {

			if cerr == nil {
				ops = append(ops, IInt{v})
			} else if instructions, ok := custom_words[strings.ToLower(x)]; ok {
				ops = append(ops, instructions...)
			} else {
				err = errors.New("undefined word used")
				return
			}
		}
	}
	ops, err = simplify(ops)
	return
}

func Forth(inp []string) ([]int, error) {
	words := map[string]([]Op){
		"+":    []Op{IArithmetic{Add}},
		"-":    []Op{IArithmetic{Sub}},
		"*":    []Op{IArithmetic{Mul}},
		"/":    []Op{IArithmetic{Div}},
		"dup":  []Op{IDup{}},
		"drop": []Op{IDrop{}},
		"swap": []Op{ISwap{}},
		"over": []Op{IOver{}},
	}
	for _, x := range inp {
		instructions, custom_word, err := tokens(x, words)
		if err != nil {
			return nil, err
		}
		if len(custom_word) > 0 {
			var cerr error
			words[custom_word], cerr = simplify(instructions)
			if cerr != nil {
				return nil, cerr
			}
		} else {
			return eval(instructions)
		}
	}
	return nil, nil
}
