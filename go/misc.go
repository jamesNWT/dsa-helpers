func isSameCharDifferentCase(a, b byte) bool {
    return int(a) - int(b) == 32 || int(b) - int(a) == 32
}
