package lense

import (
	"math"
	"math/rand"

	"github.com/tanema/amore/gfx"
)

const (
	maxShake        = float32(5)
	atenuationSpeed = float32(4)
)

// Camera is the main object to handle all camera operations
type Camera struct {
	x, y           float32
	width, height  float32
	halfWidth      float32
	halfHeight     float32
	scale          float32
	rot            float32
	shakeIntensity float32
}

// New will return a new Camera object for the current screen extents
func New() *Camera {
	w, h := gfx.GetWidth(), gfx.GetHeight()

	return &Camera{
		width:      w,
		height:     h,
		halfWidth:  w * 0.5,
		halfHeight: h * 0.5,
		scale:      1,
	}
}

// LookAt will set the camera to looking at the provided x, y coords.
func (camera *Camera) LookAt(x, y float32) {
	camera.x, camera.y = x, y
}

// Move will move the camera by the difference of x or y
func (camera *Camera) Move(dx, dy float32) {
	camera.x, camera.y = camera.x+dx, camera.y+dy
}

// Rotate will rotate the view of the camera around the center coords that the
// camera is looking at
func (camera *Camera) Rotate(phi float32) {
	camera.rot = camera.rot + phi
}

// RotateTo will set the camera rotation to the radian provided
func (camera *Camera) RotateTo(phi float32) {
	camera.rot = phi
}

// Zoom will scale the scene by the multiplyer provided. for instance if 2 is provided
// then everything is scaled to twice the size.
func (camera *Camera) Zoom(mul float32) {
	camera.scale = camera.scale * mul
}

// ZoomTo will set the scale factor. If 1 is passed then the zoom will be default
func (camera *Camera) ZoomTo(zoom float32) {
	camera.scale = zoom
}

// GetVisible returns the extents of the visible area that the camera can see. The
// top left x, y and the width and height of the frame.
func (camera *Camera) GetVisible() (l, t, w, h float32) {
	return camera.x - camera.halfWidth, camera.y - camera.halfHeight, camera.width, camera.height
}

// Draw is the function to scope all your draw calls to be transformed to the camera veiw.
// Pass in a function that takes the extents of the camera to be called for drawing the
// scene within the camera view.
func (camera *Camera) Draw(draw func(l, t, w, h float32)) {
	gfx.Push()
	{
		gfx.Translate(-camera.x, -camera.y)
		gfx.Scale(camera.scale)
		gfx.Rotate(camera.rot)
		gfx.Translate(camera.halfWidth, camera.halfHeight)

		l, t, w, h := camera.GetVisible()
		draw(l, t, w, h)
	}
	gfx.Pop()
}

// Shake will cause the camera to jump around making it seem like the whole view
// is shaking. If you want to use this function then you will need to call Update
// in your global update function.
func (camera *Camera) Shake(intensity float32) {
	camera.shakeIntensity = min(maxShake, camera.shakeIntensity+intensity)
}

// Update will update the shake functionality. This method is only needed if Shake
// is used.
func (camera *Camera) Update(dt float32) {
	camera.shakeIntensity = max(0, camera.shakeIntensity-atenuationSpeed*dt)
	if camera.shakeIntensity > 0 {
		camera.x += (100 - 200*randMax(camera.shakeIntensity)) * dt
		camera.y += (100 - 200*randMax(camera.shakeIntensity)) * dt
	}
}

func randMax(max float32) float32 {
	return randRange(0, max)
}

func randRange(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}

func min(x, y float32) float32 {
	return float32(math.Min(float64(x), float64(y)))
}

func max(x, y float32) float32 {
	return float32(math.Max(float64(x), float64(y)))
}
