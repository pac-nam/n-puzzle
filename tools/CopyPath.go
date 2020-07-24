package tools

import s "n-puzzle/structures"

func CopyPath(path []s.Tnumber) []s.Tnumber {
	tmp := make([]s.Tnumber, len(path))
	copy(tmp, path)
	return tmp
}