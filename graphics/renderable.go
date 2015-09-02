package graphics

type Renderable interface {
    Render()
}

var Renderables []Renderable

func RenderAllRenderables() {

}