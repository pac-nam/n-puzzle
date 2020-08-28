package messages

const Help = `Usage of ./n-puzzle [options...] <file> :
-C	Activate visual checker
-H	Heuristic you can use:
	'manhattan' or 'm' or '1'
	'hamming'   or 'h' or '2'
	'euclidian' or 'e' or '3'
	'complete'  or 'c' or '4'
	(default "manhattan")
-c	Activate simple checker
-i	Activate ID A*
-h	Print this help
`
const AtoiError = "Atoi error"
const OpenError = "Open error"
const ReadError = "Read error"
const InvalidFirstLine = "Invalid first line"
const TooSmall = "Size should be at least 2"
const InvalidHeuristic = `Invalid heuristic
	you can use:
	'manhattan' or 'm' or '1'
	'hamming'   or 'h' or '2'
	'euclidian' or 'e' or '3'
	'complete'  or 'c' or '4'
`
const Heuristic = `Heuristic you can use:
'manhattan' or 'm' or '1'
'hamming'   or 'h' or '2'
'euclidian' or 'e' or '3'
'complete'  or 'c' or '4'
`