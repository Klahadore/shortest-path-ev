package graph

import (
	"reflect"
	"testing"
)

func TestNewGraph_InitializesAdjList(t *testing.T) {
	g := NewGraph()
	if g == nil {
		t.Fatalf("NewGraph returned nil")
	}
	if g.adjList == nil {
		t.Fatalf("adjList should be initialized")
	}
	if len(g.adjList) != 0 {
		t.Fatalf("expected empty adjList, got %d entries", len(g.adjList))
	}
}

func TestAddVertex_CreatesEmptySlice(t *testing.T) {
	g := NewGraph()
	g.AddVertex(1)

	got, ok := g.adjList[1]
	if !ok {
		t.Fatalf("vertex 1 should exist in adjList")
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("expected empty neighbor list for vertex 1, got %v", got)
	}

	// Adding the same vertex again should not duplicate or change neighbors
	g.AddVertex(1)
	got2, ok2 := g.adjList[1]
	if !ok2 {
		t.Fatalf("vertex 1 should still exist in adjList")
	}
	if !reflect.DeepEqual(got2, got) {
		t.Fatalf("neighbor list should remain unchanged; got %v, want %v", got2, got)
	}
}

func TestAddEdge_AddsVerticesAndEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2, 3)

	// Both vertices should exist
	if _, ok := g.adjList[1]; !ok {
		t.Fatalf("from-vertex 1 should exist")
	}
	if _, ok := g.adjList[2]; !ok {
		t.Fatalf("to-vertex 2 should exist")
	}

	// Edge 1 -> 2 should be present
	neighbors := g.adjList[1]
	if len(neighbors) != 1 || neighbors[0].To != 2 {
		t.Fatalf("expected neighbors[1] = [2], got %v", neighbors)
	}

	// Adding another edge from 1 to 3 should append
	g.AddEdge(1, 3, 2)
	neighbors = g.adjList[1]
	if len(neighbors) != 2 || neighbors[0].To != 2 || neighbors[1].To != 3 {
		t.Fatalf("expected neighbors[1] = [2 3], got %v", neighbors)
	}
}

func TestReturnVertices(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 3, 4)
	g.AddEdge(2, 4, 3)
	g.AddVertex(5)

	vertices := g.GetVertices()

	// Check that we have the correct number of vertices
	expectedCount := 5
	if len(vertices) != expectedCount {
		t.Errorf("Expected %d vertices, but got %d", expectedCount, len(vertices))
	}

	// Check that all expected vertices are present
	expectedVertices := []int{1, 2, 3, 4, 5}
	for _, expected := range expectedVertices {
		found := false
		for _, vertex := range vertices {
			if vertex == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected vertex %d not found in vertices", expected)
		}
	}

	// Check that there are no duplicate vertices
	vertexSet := make(map[int]bool)
	for _, vertex := range vertices {
		if vertexSet[vertex] {
			t.Errorf("Duplicate vertex %d found", vertex)
		}
		vertexSet[vertex] = true
	}
}
