package operator

// Code generated by (tawesoft.co.uk/go/operator) template-numbers.py: DO NOT EDIT.


type int8Binary struct {
    Add             func(int8, int8) int8
    Sub             func(int8, int8) int8
    Mul             func(int8, int8) int8
    Div             func(int8, int8) int8
    Mod             func(int8, int8) int8
    
    And             func(int8, int8) int8
    Or              func(int8, int8) int8
    Xor             func(int8, int8) int8
    AndNot          func(int8, int8) int8
    
    Shl             func(int8, uint) int8
    Shr             func(int8, uint) int8
}

type int8BinaryChecked struct {
    Add             func(int8, int8) (int8, error)
    Sub             func(int8, int8) (int8, error)
    Mul             func(int8, int8) (int8, error)
    
    Shl             func(int8, uint) (int8, error)
    Shr             func(int8, uint) (int8, error)
}

var Int8 = struct {
    Binary          int8Binary
}{
    Binary:          int8Binary{
        Add:        func(a int8, b int8) int8 { return a + b },
        Sub:        func(a int8, b int8) int8 { return a - b },
        Mul:        func(a int8, b int8) int8 { return a * b },
        Div:        func(a int8, b int8) int8 { return a / b },
        
        And:        func(a int8, b int8) int8 { return a & b },
        Or:         func(a int8, b int8) int8 { return a | b },
        Xor:        func(a int8, b int8) int8 { return a ^ b },
        AndNot:     func(a int8, b int8) int8 { return a &^ b },
        Mod:        func(a int8, b int8) int8 { return a % b },
        
        Shl:        func(a int8, b uint) int8 { return a << b },
        Shr:        func(a int8, b uint) int8 { return a >> b },
    },
}

var Int8Checked = struct {
    Binary          int8BinaryChecked
}{
    Binary:         int8BinaryChecked{
        Add:        int8BinaryCheckedAdd,
        Sub:        int8BinaryCheckedSub,
        Mul:        int8BinaryCheckedMul,
        Shl:        int8BinaryCheckedShl,
    },
}

func int8BinaryCheckedAdd(a int8, b int8) (v int8, err error) {
    if (b > 0) && (a > (maxInt8 - b)) { return v, ErrorOverflow }
    if (b < 0) && (a < (minInt8 - b)) { return v, ErrorOverflow }
    return a + b, nil
}

func int8BinaryCheckedSub(a int8, b int8) (v int8, err error) {
    if (b < 0) && (a > (maxInt8 + b)) { return v, ErrorOverflow }
    if (b > 0) && (a < (minInt8 + b)) { return v, ErrorOverflow }
    return a - b, nil
}

func int8BinaryCheckedMul(a int8, b int8) (v int8, err error) {
    if (a == -1) && (b == minInt8) { return v, ErrorOverflow }
    if (b == -1) && (a == minInt8) { return v, ErrorOverflow }
    if (a > (maxInt8 / b)) { return v, ErrorOverflow }
    if (a < (minInt8 / b)) { return v, ErrorOverflow }
    
    return a * b, nil
}

func int8BinaryCheckedShl(a int8, b uint) (v int8, err error) {
    if a < 0 { return v, ErrorUndefined }
    if b > uint(int8MostSignificantBit(maxInt8)) { return v, ErrorOverflow }
    return v, err
}

func int8MostSignificantBit(a int8) (result int) {
  for a > 0 {
      a >>= 1
      result++
  }
  return result;
}