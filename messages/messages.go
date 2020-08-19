package messages

const Help = "Invalid number of arguments\n ./n-puzzle <heuristic> <file>"
const AtoiError = "Atoi error"
const OpenError = "Open error"
const ReadError = "Read error"
const InvalidFirstLine = "Invalid first line"
const TooSmall = "Size should be at least 2"
const InvalidHeuristic = `Invalid heuristic\nyou can use:
	'manhattan' or 'm' or '1'
	'hamming'   or 'h' or '2'
	'euclidian' or 'e' or '3'
	'complete'  or 'c' or '4'
`