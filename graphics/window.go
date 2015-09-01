package graphics

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/veandco/go-sdl2/sdl"
)

type Window struct {
    Title string
    Width int
    Height int
    Window *sdl.Window
    Context sdl.GLContext
}

func NewWindow() *Window {
    var err error

    w := &Window{
        Title: "Bottom Dollar",
        Width: 800,
        Height: 600,
    }

    if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        panic(err)
    }

    if err = gl.Init(); err != nil {
        panic(err)
    }

    w.Window, err = sdl.CreateWindow(w.Title,
                                     sdl.WINDOWPOS_UNDEFINED,
                                     sdl.WINDOWPOS_UNDEFINED,
                                     w.Width, w.Height, sdl.WINDOW_OPENGL)
    if err != nil {
        panic(err)
    }

    w.Context, err = sdl.GL_CreateContext(w.Window)
    if err != nil {
        panic(err)
    }

    gl.Enable(gl.DEPTH_TEST)
    gl.ClearColor(0.2, 0.2, 0.3, 1.0)
    gl.ClearDepth(1)
    gl.DepthFunc(gl.LEQUAL)
    gl.Viewport(0, 0, int32(w.Width), int32(w.Height))

    return w
}

func (w *Window) ProcessOneFrame() {
    drawgl()
    sdl.GL_SwapWindow(w.Window)
}


func drawgl() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    gl.Begin(gl.TRIANGLES)
    gl.Color3f(1.0, 0.0, 0.0)
    gl.Vertex2f(0.5, 0.0)
    gl.Color3f(0.0, 1.0, 0.0)
    gl.Vertex2f(-0.5, -0.5)
    gl.Color3f(0.0, 0.0, 1.0)
    gl.Vertex2f(-0.5, 0.5)
    gl.End()
}

func (w *Window) Destroy() {
    sdl.GL_DeleteContext(w.Context)
    w.Window.Destroy()
    sdl.Quit()
}

