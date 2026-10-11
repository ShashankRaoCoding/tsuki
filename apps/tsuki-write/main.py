from textual.app import App
from textual.widgets import Static  ,Label

class App(App):
    def compose(self):
        self.Label = Label("Textual2") 
        self.Static = Static(
            "[bold green]Textual[/bold green]", 
        )

        yield self.Label
        yield self.Static

    def on_mount(self):
        self.Label.styles.background = "darkgreen"
        self.Label.styles.border = ("solid","green")
        self.Static.styles.background = "purple" 
        self.Static.styles.text_align = "center" 
        self.Static.styles.border = ("solid", "black") 

if __name__ == "__main__": App().run() 












































