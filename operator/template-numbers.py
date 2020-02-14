
template=""
template_file="template_numbers_test.go"

numbers = [
    # class,     type       signed
    ("integers", "int",     True),
    ("integers", "int8",    True),
    ("integers", "int16",   True),
    ("integers", "int32",   True),
    ("integers", "int64",   True),
    ("integers", "uint",    False),
    ("integers", "uint8",   False),
    ("integers", "uint16",  False),
    ("integers", "uint32",  False),
    ("integers", "uint64",  False),
    ("floats",   "float32", None),
    ("floats",   "float64", None),
]

with open(template_file, 'r') as fp:
    template = fp.read()

for _class, _type, signed in numbers:

    def replace(s):
        for r in (("_T", _type.title()), ("_t", _type)):
            s = s.replace(*r)
        return s

    with open("%s.go" % _type, 'w') as fp:
        for line in template.splitlines(keepends=True):
            if line.startswith("package"):
                fp.write(line)
                fp.write("\n// Code generated by (tawesoft.co.uk/go/operator) template-numbers.py: DO NOT EDIT.\n")
            elif "//" not in line:
                fp.write(replace(line))
            else:
                pos = line.rfind("//")
                left, right = line[0:pos].rstrip(), line[pos + 2:].strip()

                rest = None
                if ";" in right:
                    right, rest = right.split(";", maxsplit=1)
                    rest = set([x.strip() for x in rest.split(";")])

                if right.startswith("IGNORE"):
                    continue
                elif right.startswith("CLASS"):
                    classes = right[right.find(" "):].split(",")
                    classes = set([x.strip() for x in classes])
                    if _class in classes:

                        if (rest is None) or \
                            (("signed" in rest) and signed):

                            fp.write(replace(left)+"\n")


