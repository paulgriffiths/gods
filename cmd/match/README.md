# match

**match** demonstrates the simple regular expression package.

## Usage examples

	paul@horus:match$ ./match '0*|1*' numbers.txt
	0
	1
	00
	11
	000
	111
	0000
	1111
	00000
	11111
	000000
	111111
	paul@horus:match$ ./match '1*|b*' numbers.txt letters.txt
	1
	11
	111
	1111
	11111
	111111
	b
	bb
	bbb
	bbbb
	bbbbb
	bbbbbb
	paul@horus:match$ cat numbers.txt | ./match '1101(0|1)*'
	1101
	11010
	11011
	110100
	110101
	110110
	110111
	paul@horus:match$ ./match '(aa|bb)*' letters.txt
	aa
	bb
	aaaa
	aabb
	bbaa
	bbbb
	aaaaaa
	aaaabb
	aabbaa
	aabbbb
	bbaaaa
	bbaabb
	bbbbaa
	bbbbbb
	paul@horus:match$ 
