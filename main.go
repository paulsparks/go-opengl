package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600

func init() {
	// make sure that main() runs on the main thread
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT)

	window.SetFramebufferSizeCallback(framebuffer_size_callback)

	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	renderLoop(window)
}

func framebuffer_size_callback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT)
}

func renderLoop(window *glfw.Window) {
	for !window.ShouldClose() {
		// input
		processInput(window)

		//rendering
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// call events, swap buffers
		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}
