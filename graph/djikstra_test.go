package graph

import "testing"

func TestShortestPath_SimpleChain(t *testing.T) {
    g := NewGraph()
    g.AddEdge(1, 2, 3)
    g.AddEdge(2, 3, 4)

    dist := ShortestPath(1, g)

    if dist[1] != 0 {
        t.Fatalf("dist to 1 = %d, want 0", dist[1])
    }
    if dist[2] != 3 {
        t.Fatalf("dist to 2 = %d, want 3", dist[2])
    }
    if dist[3] != 7 {
        t.Fatalf("dist to 3 = %d, want 7", dist[3])
    }
}

func TestShortestPath_ChoosesCheapestIndirect(t *testing.T) {
    g := NewGraph()
    g.AddEdge(1, 2, 10)
    g.AddEdge(1, 3, 1)
    g.AddEdge(3, 2, 1)

    dist := ShortestPath(1, g)

    if dist[2] != 2 {
        t.Fatalf("dist to 2 = %d, want 2 (via 1->3->2)", dist[2])
    }
    if dist[3] != 1 {
        t.Fatalf("dist to 3 = %d, want 1", dist[3])
    }
}

func TestShortestPath_UnreachableVerticesRemainMaxInt(t *testing.T) {
    g := NewGraph()
    // Connected component
    g.AddEdge(1, 2, 5)
    // Disconnected vertex 4
    g.AddVertex(4)

    dist := ShortestPath(1, g)

    if dist[1] != 0 {
        t.Fatalf("dist to 1 = %d, want 0", dist[1])
    }
    if dist[2] != 5 {
        t.Fatalf("dist to 2 = %d, want 5", dist[2])
    }
    if d, ok := dist[4]; !ok {
        t.Fatalf("no distance recorded for vertex 4")
    } else if d != maxInt {
        t.Fatalf("dist to 4 = %d, want maxInt (%d)", d, maxInt)
    }
}

func TestShortestPath_ZeroWeightEdges(t *testing.T) {
    g := NewGraph()
    g.AddEdge(1, 2, 0)
    g.AddEdge(2, 3, 0)
    g.AddEdge(1, 3, 5)

    dist := ShortestPath(1, g)

    if dist[2] != 0 {
        t.Fatalf("dist to 2 = %d, want 0", dist[2])
    }
    if dist[3] != 0 { // 1->2->3 beats direct 1->3
        t.Fatalf("dist to 3 = %d, want 0", dist[3])
    }
}

