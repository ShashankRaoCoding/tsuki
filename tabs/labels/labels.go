package labels

type Label string 

func (l Label) Render(w width) string {
	label , _ := l.(string) 
	elipses := "..." 

	if len(label) == 0 {
		label = "Untitled" 
	}

	if len(label) <= len(elipses) || len(label) <= w {
		return label 
	}

}










































