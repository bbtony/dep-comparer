package dot

import (
	"dep-comparer/internal/parser"
	"os"
	"strconv"
	"time"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func NewReport(
	modules ...*parser.Module,
) (string, error) {
	g := graph.New(graph.StringHash, graph.Directed())
	for _, module := range modules {
		_ = g.AddVertex(string(module.ModulePath), graph.VertexAttributes(map[string]string{
			"URL":   "https://" + string(module.ModulePath),
			"color": "lightblue",
			"style": "filled",
			"shape": "box",
		}))
		for path, _ := range module.Dependencies {
			_ = g.AddVertex(string(path), graph.VertexAttributes(map[string]string{
				"URL":   "https://" + string(path),
				"color": "lightgreen",
				"style": "filled",
				"shape": "hexagon",
			}))
			_ = g.AddEdge(string(module.ModulePath), string(path)) //, graph.EdgeAttribute("taillabel", string(version)))
		}
	}

	fileName := "graph_" + strconv.FormatInt(time.Now().Unix(), 10) + ".dot"
	file, _ := os.Create(fileName)
	return fileName, draw.DOT(g, file)
}
