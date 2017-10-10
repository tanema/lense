# lense

A camera system for [amore](https://github.com/tanema/amore)

Example
=======

```
package main

import (
  "github.com/tanema/amore"
  "github.com/tanema/amore/gfx"
  "github.com/tanema/amore/keyboard"
  "github.com/tanema/lense"
)

var (
  camera *lense.Camera
)

func main() {
  amore.Start(update, draw)
}

func update(dt float32) {
  if keyboard.IsDown(keyboard.KeyLeft) {
    camera.Move(-1, 0)
  } else if keyboard.IsDown(keyboard.KeyRight) {
    camera.Move(1, 0)
  }

  if keyboard.IsDown(keyboard.KeyUp) {
    camera.Move(0, -1)
  } else if keyboard.IsDown(keyboard.KeyDown) {
    camera.Move(0, 1)
  }

  if keyboard.IsDown(keyboard.KeyS) {
    camera.Shake(2)
  }

  camera.Update(dt)
}

func draw() {
  camera.Draw(func(l, t, w, h float32){
    gfx.Rect(gfx.FILL, 200, 200, 200, 200)
  })
}
```
