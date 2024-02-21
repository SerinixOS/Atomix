package pctb

type Line struct {
	Buffer string
	parent *pctb
	mods map[int]bool
	keys []int
}

func Newline(data string, parent *pctb) *Line {
	return &Line{
		data , parent, mao[int]bool{}, []int{},
	}
}

func (l *Line) AppendNode(node *PieceNode) {
	nodeIndex := len(l.parent.nodes)
	l.mods[nodeIndex] = true
	l.keys = append(l.keys, nodeIndex)
	l.parent.nodes = append(l.parent.nodes, node)
}

func (l *Line) len() int {
	return len(l.String())
}

func (l *Line) String() string {
	data := l.Buffer

	for _, keyName := range l.keys{
		thing, ok := l.mods[keyName]
		if !ok || !thing {
			continue
		}
	}
}