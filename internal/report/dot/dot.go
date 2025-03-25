package dot

import (
	"dep-comparer/internal/parser/types"
	"os"
	"strconv"
	"time"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func NewReport(
	dependencies ...*types.Dependency,
) (string, error) {
	g := graph.New(graph.StringHash, graph.Directed())
	for _, dependency := range dependencies {
		_ = g.AddVertex(string(dependency.DependencyPath), graph.VertexAttributes(map[string]string{
			"URL":   "https://" + string(dependency.DependencyPath),
			"color": "lightblue",
			"style": "filled",
			"shape": "box",
		}))
		for path, _ := range dependency.Dependencies {
			_ = g.AddVertex(string(path), graph.VertexAttributes(map[string]string{
				"URL":   "https://" + string(path),
				"color": "lightgreen",
				"style": "filled",
				"shape": "hexagon",
			}))
			_ = g.AddEdge(string(dependency.DependencyPath), string(path)) //, graph.EdgeAttribute("taillabel", string(version)))
		}
	}

	fileName := "graph_" + strconv.FormatInt(time.Now().Unix(), 10) + ".dot"
	file, _ := os.Create(fileName)
	return fileName, draw.DOT(g, file)
}
