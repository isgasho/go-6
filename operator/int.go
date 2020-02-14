package operator

// Code generated by (tawesoft.co.uk/go/operator) template-numbers.py: DO NOT EDIT.


type intBinary struct {
    Add             func(int, int) int
    Sub             func(int, int) int
    Mul             func(int, int) int
    Div             func(int, int) int
    Mod             func(int, int) int
    
    And             func(int, int) int
    Or              func(int, int) int
    Xor             func(int, int) int
    AndNot          func(int, int) int
    
    Shl             func(int, uint) int
    Shr             func(int, uint) int
}

type intBinaryChecked struct {
    Add             func(int, int) (int, error)
    Sub             func(int, int) (int, error)
    Mul             func(int, int) (int, error)
    
    Shl             func(int, uint) (int, error)
    Shr             func(int, uint) (int, error)
}

var Int = struct {
    Binary          intBinary
}{
    Binary:          intBinary{
        Add:        func(a int, b int) int { return a + b },
        Sub:        func(a int, b int) int { return a - b },
        Mul:        func(a int, b int) int { return a * b },
        Div:        func(a int, b int) int { return a / b },
        
        And:        func(a int, b int) int { return a & b },
        Or:         func(a int, b int) int { return a | b },
        Xor:        func(a int, b int) int { return a ^ b },
        AndNot:     func(a int, b int) int { return a &^ b },
        Mod:        func(a int, b int) int { return a % b },
        
        Shl:        func(a int, b uint) int { return a << b },
        Shr:        func(a int, b uint) int { return a >> b },
    },
}

var IntChecked = struct {
    Binary          intBinaryChecked
}{
    Binary:         intBinaryChecked{
        Add:        intBinaryCheckedAdd,
        Sub:        intBinaryCheckedSub,
        Mul:        intBinaryCheckedMul,
        Shl:        intBinaryCheckedShl,
    },
}

func intBinaryCheckedAdd(a int, b int) (v int, err error) {
    if (b > 0) && (a > (maxInt - b)) { return v, ErrorOverflow }
    if (b < 0) && (a < (minInt - b)) { return v, ErrorOverflow }
    return a + b, nil
}

func intBinaryCheckedSub(a int, b int) (v int, err error) {
    if (b < 0) && (a > (maxInt + b)) { return v, ErrorOverflow }
    if (b > 0) && (a < (minInt + b)) { return v, ErrorOverflow }
    return a - b, nil
}

func intBinaryCheckedMul(a int, b int) (v int, err error) {
    if (a == -1) && (b == minInt) { return v, ErrorOverflow }
    if (b == -1) && (a == minInt) { return v, ErrorOverflow }
    if (a > (maxInt / b)) { return v, ErrorOverflow }
    if (a < (minInt / b)) { return v, ErrorOverflow }
    
    return a * b, nil
}

func intBinaryCheckedShl(a int, b uint) (v int, err error) {
    if a < 0 { return v, ErrorUndefined }
    if b > uint(intMostSignificantBit(maxInt)) { return v, ErrorOverflow }
    return v, err
}

func intMostSignificantBit(a int) (result int) {
  for a > 0 {
      a >>= 1
      result++
  }
  return result;
}