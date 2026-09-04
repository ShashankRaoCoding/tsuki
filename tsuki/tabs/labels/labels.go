package labels

type Label struct {
	Content string 
}

func (l Label) Render(w int) string {
	label := l.Content
	elipses := "..." 

	if len(label) == 0 {
		label = "Untitled" 
	}

	if len(label) <= len(elipses) {
		return label 
	}

	if len(label) <= w {
		return label 
	}

	label = label[0:w - len(elipses)] + elipses
	return label 
}

func New(label string) Label {
	l := Label{Content: label} 
	return l 
}








































