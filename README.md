# Draw

`draw` is an simple drawing tool in the terminal.
Hold your mouse down and move it across the screen to draw anything you want!

<p align="center">
  <img src="./assets/draw.png?raw=true" alt="Draw" />
</p>

### Installation

```
go install github.com/maaslalani/draw@latest
```

### Usage
```
draw
```

### Boxes

Draw boxes by clicking and dragging the right mouse button, releasing the
right mouse button will draw the box.

Voila! You have a box drawn where you want it.

<p align="center">
  <img src="./assets/boxes.png?raw=true" alt="Draw Boxes" />
</p>

### Text Insertion

Tap the right mouse button to set an anchor for inserting text and type away!
Press <kbd>ESC</kbd> to exit insert mode.

<p align="center">
  <img src="./assets/text.png?raw=true" alt="Insert Text" />
</p>


### Saving Drawings

`draw` automatically saves your latest masterpiece to `/tmp/draw.txt` by
default. You can change this location by setting the environment variable
`$DRAW_FILE`.

```
export DRAW_FILE=/path/to/draw.txt
```

To view this drawing use `cat` or `mv` the drawing to somewhere more permanent.
If you already know you want to save the drawing, run `draw` with the file path
where you want to save the drawing: `draw masterpiece.txt`.

### Controls
* <kbd>ctrl+c</kbd> to exit.
* Press any key to draw with that character.
* Press <kbd>1</kbd> for red.
* Press <kbd>2</kbd> for green.
* Press <kbd>3</kbd> for yellow.
* Press <kbd>4</kbd> for blue.
* Press <kbd>5</kbd> for magenta.
* Press <kbd>6</kbd> for cyan.
* Press <kbd>7</kbd> for gray.
