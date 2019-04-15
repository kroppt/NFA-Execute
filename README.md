# NFA-Execute

### Install

Requires [Go](https://golang.org/)

```
$ go get github.com/kroppt/NFA-Execute
```

### Usage

```
$ NFA-Execute -in=str.txt nfa.txt
```
or
```
$ NFA-Execute nfa.txt < str.txt
```

### Format

NFA:
```
<node count>
<list of exit nodes>
<node>
...
<node>
```
Nodes follow format:
```
<from> <to> <character>
```
where character can also be special symbol ε

Input string can be anything.

State 0 is assumed to be the starting state.

### Example

`nfa.txt`:
```
4
3
0 1 0
0 2 1
1 2 1
1 3 ε
2 1 0
2 3 ε
```
`str.txt`:
```
010
```
Output:
```
input accepted
```
