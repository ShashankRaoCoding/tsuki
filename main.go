package main

import (
	"fmt"
	"time" 
	c "tsuki/components"
)

func main() {
	clock := &c.Clock{
		Unit: time.Second,
		Progress: 0,
		Total: 86400, 
	}

	app := c.App{
		clock , 
		// &c.FlexDiv{
		// 	&c.FileTree,
		// 	&c.FileEditor, 
		// },
		// &c.Console, 
	}

	c.FOCUS = app[0] 
	app.Start()
	fmt.Println(c.SCREEN.Size()) 
}









































