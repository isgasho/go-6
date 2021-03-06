package operator

// Code generated by (tawesoft.co.uk/go/operator) template-numbers.py: DO NOT EDIT.


// Some overflow checks with reference to stackoverflow.com/a/1514309/5654201

type int16Unary struct {
    Identity        func(int16) int16
    Abs             func(int16) int16
    Negation        func(int16) int16
    Zero            func(int16) bool
    NonZero         func(int16) bool
    Positive        func(int16) bool
    Negative        func(int16) bool
}

type int16UnaryChecked struct {
    Abs             func(int16) (int16, error)
    Negation        func(int16) (int16, error)
}

type int16Binary struct {
    Add             func(int16, int16) int16
    Sub             func(int16, int16) int16
    Mul             func(int16, int16) int16
    Div             func(int16, int16) int16
    Mod             func(int16, int16) int16
    
    Eq              func(int16, int16) bool
    Neq             func(int16, int16) bool
    Lt              func(int16, int16) bool
    Lte             func(int16, int16) bool
    Gt              func(int16, int16) bool
    Gte             func(int16, int16) bool
    
    And             func(int16, int16) int16
    Or              func(int16, int16) int16
    Xor             func(int16, int16) int16
    AndNot          func(int16, int16) int16
    
    Shl             func(int16, uint) int16
    Shr             func(int16, uint) int16
}

type int16BinaryChecked struct {
    Add             func(int16, int16) (int16, error)
    Sub             func(int16, int16) (int16, error)
    Mul             func(int16, int16) (int16, error)
    Div             func(int16, int16) (int16, error)
    
    Shl             func(int16, uint) (int16, error)
    Shr             func(int16, uint) (int16, error)
}

type int16Nary struct {
    Add             func(... int16) int16
    Sub             func(... int16) int16
    Mul             func(... int16) int16
}

type int16NaryChecked struct {
    Add             func(... int16) (int16, error)
    Sub             func(... int16) (int16, error)
    Mul             func(... int16) (int16, error)
}

// Int16 implements operations on one (unary), two (binary), or many (nary) arguments of type int16.
var Int16 = struct {
    Unary           int16Unary
    Binary          int16Binary
    Nary            int16Nary
    Reduce          func(operatorIdentity int16, operator func(int16, int16) int16, elements ... int16) int16
}{
    Unary:          int16Unary{
        Identity:   func(a int16) int16 { return a },
        Abs:        int16UnaryAbs,
        Negation:   func(a int16) int16 { return -a },
        Zero:       func(a int16) bool { return a == 0 },
        NonZero:    func(a int16) bool { return a != 0 },
        Positive:   int16UnaryPositive,
        Negative:   int16UnaryNegative,
    },
    
    Binary:          int16Binary{
        Add:        func(a int16, b int16) int16 { return a + b },
        Sub:        func(a int16, b int16) int16 { return a - b },
        Mul:        func(a int16, b int16) int16 { return a * b },
        Div:        func(a int16, b int16) int16 { return a / b },
        
        Eq:         func(a int16, b int16) bool { return a == b },
        Neq:        func(a int16, b int16) bool { return a != b },
        Lt:         func(a int16, b int16) bool { return a <  b },
        Lte:        func(a int16, b int16) bool { return a <= b },
        Gt:         func(a int16, b int16) bool { return a >  b },
        Gte:        func(a int16, b int16) bool { return a >= b },
        
        And:        func(a int16, b int16) int16 { return a & b },
        Or:         func(a int16, b int16) int16 { return a | b },
        Xor:        func(a int16, b int16) int16 { return a ^ b },
        AndNot:     func(a int16, b int16) int16 { return a &^ b },
        Mod:        func(a int16, b int16) int16 { return a % b },
        
        Shl:        func(a int16, b uint) int16 { return a << b },
        Shr:        func(a int16, b uint) int16 { return a >> b },
    },
    
    Nary:           int16Nary{
        Add:        int16NaryAdd,
        Mul:        int16NaryMul,
    },
    
    Reduce:         int16Reduce,
}

// Int16Checked implements operations on one (unary), two (binary), or many (nary) arguments of type int16, returning an
// error in cases such as overflow or an undefined operation.
var Int16Checked = struct {
    Unary           int16UnaryChecked
    Binary          int16BinaryChecked
    Nary            int16NaryChecked
    Reduce          func(operatorIdentity int16, operator func(int16, int16) (int16, error), elements ... int16) (int16, error)
}{
    Unary:          int16UnaryChecked{
        Abs:        int16UnaryCheckedAbs,
        Negation:   int16UnaryCheckedNegation,
    },
    
    Binary:         int16BinaryChecked{
        Add:        int16BinaryCheckedAdd,
        Sub:        int16BinaryCheckedSub,
        Mul:        int16BinaryCheckedMul,
        Div:        int16BinaryCheckedDiv,
        Shl:        int16BinaryCheckedShl,
    },
    
    Nary:           int16NaryChecked{
        Add:        int16NaryCheckedAdd,
        Mul:        int16NaryCheckedMul,
    },
    
    Reduce:         int16CheckedReduce,
}

func int16UnaryPositive(a int16) bool {
    return a > 0
}

func int16UnaryNegative(a int16) bool {
    return a < 0
}

func int16UnaryAbs(a int16) int16 {
    if a < 0 { return -a }
    return a
}

func int16UnaryCheckedAbs(a int16) (v int16, err error) {
    if a == minInt16 { return v, ErrorOverflow }
    if a < 0 { return -a, nil }
    return a, nil
}

func int16UnaryCheckedNegation(a int16) (v int16, err error) {
    if (a == minInt16) { return v, ErrorOverflow }
    return -a, nil
}

func int16BinaryCheckedAdd(a int16, b int16) (v int16, err error) {
    if (b > 0) && (a > (maxInt16 - b)) { return v, ErrorOverflow }
    if (b < 0) && (a < (minInt16 - b)) { return v, ErrorOverflow }
    return a + b, nil
}

func int16BinaryCheckedSub(a int16, b int16) (v int16, err error) {
    if (b < 0) && (a > (maxInt16 + b)) { return v, ErrorOverflow }
    if (b > 0) && (a < (minInt16 + b)) { return v, ErrorOverflow }
    return a - b, nil
}

func int16BinaryCheckedMul(a int16, b int16) (v int16, err error) {
    if (a == -1) && (b == minInt16) { return v, ErrorOverflow }
    if (b == -1) && (a == minInt16) { return v, ErrorOverflow }
    if (a > (maxInt16 / b)) { return v, ErrorOverflow }
    if (a < (minInt16 / b)) { return v, ErrorOverflow }
    
    return a * b, nil
}

func int16BinaryCheckedDiv(a int16, b int16) (v int16, err error) {
    if (b == -1) && (a == minInt16) { return v, ErrorOverflow }
    if (b == 0) { return v, ErrorUndefined }
    
    return a / b, nil
}

func int16BinaryCheckedShl(a int16, b uint) (v int16, err error) {
    if a < 0 { return v, ErrorUndefined }
    if b > uint(int16MostSignificantBit(maxInt16)) { return v, ErrorOverflow }
    return v, err
}

func int16MostSignificantBit(a int16) (result int) {
  for a > 0 {
      a >>= 1
      result++
  }
  return result;
}

func int16NaryAdd(xs ... int16) (result int16) {
    for i := 0; i < len(xs); i++ {
        result += xs[i]
    }
    return result
}

func int16NaryCheckedAdd(xs ... int16) (result int16, err error) {
    for i := 0; i < len(xs); i++ {
        result, err = int16BinaryCheckedAdd(result, xs[i])
        if err != nil { return result, err }
    }
    return result, nil
}

func int16NaryMul(xs ... int16) (result int16) {
    result = 1
    for i := 0; i < len(xs); i++ {
        result *= xs[i]
    }
    return result
}

func int16NaryCheckedMul(xs ... int16) (result int16, err error) {
    result = 1
    for i := 0; i < len(xs); i++ {
        result, err = int16BinaryCheckedMul(result, xs[i])
        if err != nil { return result, err }
    }
    return result, nil
}

func int16Reduce(operatorIdentity int16, operator func(int16, int16) int16, elements ... int16) (result int16) {
    result = operatorIdentity
    for i := 0; i < len(elements); i++ {
        result = operator(result, elements[i])
    }
    return result
}

func int16CheckedReduce(operatorIdentity int16, operator func(int16, int16) (int16, error), elements ... int16) (result int16, err error) {
    result = operatorIdentity
    for i := 0; i < len(elements); i++ {
        result, err = operator(result, elements[i])
        if err != nil { return result, err }
    }
    return result, err
}

