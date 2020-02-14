package operator

// Code generated by (tawesoft.co.uk/go/operator) template-numbers.py: DO NOT EDIT.



type int32Unary struct {
    Identity        func(int32) int32
    Abs             func(int32) int32
    Negation        func(int32) int32
    Zero            func(int32) bool
    NonZero         func(int32) bool
    Positive        func(int32) bool
    Negative        func(int32) bool
}

type int32UnaryChecked struct {
    Abs             func(int32) (int32, error)
    Negation        func(int32) (int32, error)
}

type int32Binary struct {
    Add             func(int32, int32) int32
    Sub             func(int32, int32) int32
    Mul             func(int32, int32) int32
    Div             func(int32, int32) int32
    Mod             func(int32, int32) int32
    
    Eq              func(int32, int32) bool
    Neq             func(int32, int32) bool
    Lt              func(int32, int32) bool
    Lte             func(int32, int32) bool
    Gt              func(int32, int32) bool
    Gte             func(int32, int32) bool
    
    And             func(int32, int32) int32
    Or              func(int32, int32) int32
    Xor             func(int32, int32) int32
    AndNot          func(int32, int32) int32
    
    Shl             func(int32, uint) int32
    Shr             func(int32, uint) int32
}

type int32BinaryChecked struct {
    Add             func(int32, int32) (int32, error)
    Sub             func(int32, int32) (int32, error)
    Mul             func(int32, int32) (int32, error)
    
    Shl             func(int32, uint) (int32, error)
    Shr             func(int32, uint) (int32, error)
}

type int32Nary struct {
    Add             func(... int32) int32
    Sub             func(... int32) int32
    Mul             func(... int32) int32
}

type int32NaryChecked struct {
    Add             func(... int32) (int32, error)
    Sub             func(... int32) (int32, error)
    Mul             func(... int32) (int32, error)
}

var Int32 = struct {
    Unary           int32Unary
    Binary          int32Binary
    Nary            int32Nary
    Reduce          func(operatorIdentity int32, operator func(int32, int32) int32, elements ... int32) int32
}{
    Unary:          int32Unary{
        Identity:   func(a int32) int32 { return a },
        Abs:        int32UnaryAbs,
        Negation:   func(a int32) int32 { return -a },
        Zero:       func(a int32) bool { return a == 0 },
        NonZero:    func(a int32) bool { return a != 0 },
        Positive:   int32UnaryPositive,
        Negative:   int32UnaryNegative,
    },
    
    Binary:          int32Binary{
        Add:        func(a int32, b int32) int32 { return a + b },
        Sub:        func(a int32, b int32) int32 { return a - b },
        Mul:        func(a int32, b int32) int32 { return a * b },
        Div:        func(a int32, b int32) int32 { return a / b },
        
        Eq:         func(a int32, b int32) bool { return a == b },
        Neq:        func(a int32, b int32) bool { return a != b },
        Lt:         func(a int32, b int32) bool { return a <  b },
        Lte:        func(a int32, b int32) bool { return a <= b },
        Gt:         func(a int32, b int32) bool { return a >  b },
        Gte:        func(a int32, b int32) bool { return a >= b },
        
        And:        func(a int32, b int32) int32 { return a & b },
        Or:         func(a int32, b int32) int32 { return a | b },
        Xor:        func(a int32, b int32) int32 { return a ^ b },
        AndNot:     func(a int32, b int32) int32 { return a &^ b },
        Mod:        func(a int32, b int32) int32 { return a % b },
        
        Shl:        func(a int32, b uint) int32 { return a << b },
        Shr:        func(a int32, b uint) int32 { return a >> b },
    },
    
    Nary:           int32Nary{
        Add:        int32NaryAdd,
        Mul:        int32NaryMul,
    },
    
    Reduce:         int32Reduce,
}

var Int32Checked = struct {
    Unary           int32UnaryChecked
    Binary          int32BinaryChecked
    Nary            int32NaryChecked
    Reduce          func(operatorIdentity int32, operator func(int32, int32) (int32, error), elements ... int32) (int32, error)
}{
    Unary:          int32UnaryChecked{
        Abs:        int32UnaryCheckedAbs,
        Negation:   int32UnaryCheckedNegation,
    },
    
    Binary:         int32BinaryChecked{
        Add:        int32BinaryCheckedAdd,
        Sub:        int32BinaryCheckedSub,
        Mul:        int32BinaryCheckedMul,
        Shl:        int32BinaryCheckedShl,
    },
    
    Nary:           int32NaryChecked{
        Add:        int32NaryCheckedAdd,
        Mul:        int32NaryCheckedMul,
    },
    
    Reduce:         int32CheckedReduce,
}

func int32UnaryPositive(a int32) bool {
    return a > 0
}

func int32UnaryNegative(a int32) bool {
    return a < 0
}

func int32UnaryAbs(a int32) int32 {
    if a < 0 { return -a }
    return a
}

func int32UnaryCheckedAbs(a int32) (v int32, err error) {
    if a == minInt32 { return v, ErrorOverflow }
    if a < 0 { return -a, nil }
    return a, nil
}

func int32UnaryCheckedNegation(a int32) (v int32, err error) {
    if (a == minInt32) { return v, ErrorOverflow }
    return -a, nil
}

func int32BinaryCheckedAdd(a int32, b int32) (v int32, err error) {
    if (b > 0) && (a > (maxInt32 - b)) { return v, ErrorOverflow }
    if (b < 0) && (a < (minInt32 - b)) { return v, ErrorOverflow }
    return a + b, nil
}

func int32BinaryCheckedSub(a int32, b int32) (v int32, err error) {
    if (b < 0) && (a > (maxInt32 + b)) { return v, ErrorOverflow }
    if (b > 0) && (a < (minInt32 + b)) { return v, ErrorOverflow }
    return a - b, nil
}

func int32BinaryCheckedMul(a int32, b int32) (v int32, err error) {
    if (a == -1) && (b == minInt32) { return v, ErrorOverflow }
    if (b == -1) && (a == minInt32) { return v, ErrorOverflow }
    if (a > (maxInt32 / b)) { return v, ErrorOverflow }
    if (a < (minInt32 / b)) { return v, ErrorOverflow }
    
    return a * b, nil
}

func int32BinaryCheckedShl(a int32, b uint) (v int32, err error) {
    if a < 0 { return v, ErrorUndefined }
    if b > uint(int32MostSignificantBit(maxInt32)) { return v, ErrorOverflow }
    return v, err
}

func int32MostSignificantBit(a int32) (result int) {
  for a > 0 {
      a >>= 1
      result++
  }
  return result;
}

func int32NaryAdd(xs ... int32) (result int32) {
    for i := 0; i < len(xs); i++ {
        result += xs[i]
    }
    return result
}

func int32NaryCheckedAdd(xs ... int32) (result int32, err error) {
    for i := 0; i < len(xs); i++ {
        result, err = int32BinaryCheckedAdd(result, xs[i])
        if err != nil { return result, err }
    }
    return result, nil
}

func int32NaryMul(xs ... int32) (result int32) {
    result = 1
    for i := 0; i < len(xs); i++ {
        result *= xs[i]
    }
    return result
}

func int32NaryCheckedMul(xs ... int32) (result int32, err error) {
    result = 1
    for i := 0; i < len(xs); i++ {
        result, err = int32BinaryCheckedMul(result, xs[i])
        if err != nil { return result, err }
    }
    return result, nil
}

func int32Reduce(operatorIdentity int32, operator func(int32, int32) int32, elements ... int32) (result int32) {
    result = operatorIdentity
    for i := 0; i < len(elements); i++ {
        result = operator(result, elements[i])
    }
    return result
}

func int32CheckedReduce(operatorIdentity int32, operator func(int32, int32) (int32, error), elements ... int32) (result int32, err error) {
    result = operatorIdentity
    for i := 0; i < len(elements); i++ {
        result, err = operator(result, elements[i])
        if err != nil { return result, err }
    }
    return result, err
}

