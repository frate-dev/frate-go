package utils


func RemoveIndex[K any](s []K, index int) []K {
    return append(s[:index], s[index+1:]...)
}
