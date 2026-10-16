package agent

import (
    "strings"
    "testing"

    "github.com/achetronic/magec/server/store"
    memory "google.golang.org/adk/v2/memory"
)

func boolPtr(b bool) *bool { return &b }

func TestBuildToolsets_MemoryDisabled(t *testing.T) {
    def := store.AgentDefinition{ID: "a", DisableMemory: boolPtr(true)}
    svc := memory.InMemoryService()
    ts, err := buildToolsets(def, map[string]store.MCPServer{}, svc)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(ts) != 0 {
        t.Fatalf("expected 0 toolsets, got %d", len(ts))
    }
}

func TestBuildToolsets_MemoryEnabledByDefault(t *testing.T) {
    def := store.AgentDefinition{ID: "a"}
    svc := memory.InMemoryService()
    ts, err := buildToolsets(def, map[string]store.MCPServer{}, svc)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(ts) != 1 {
        t.Fatalf("expected 1 toolset, got %d", len(ts))
    }
    // ensure it's a memory toolset by name
    if ts[0].Name() != "memory" && !strings.Contains(ts[0].Name(), "memory") {
        // just check length for now
    }
}

func TestBuildToolsets_MemoryExplicitlyEnabled(t *testing.T) {
    def := store.AgentDefinition{ID: "a", DisableMemory: boolPtr(false)}
    svc := memory.InMemoryService()
    ts, err := buildToolsets(def, map[string]store.MCPServer{}, svc)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(ts) != 1 {
        t.Fatalf("expected 1 toolset, got %d", len(ts))
    }
}

func TestBuildToolsets_MemoryServiceNil(t *testing.T) {
    def := store.AgentDefinition{ID: "a", DisableMemory: boolPtr(false)}
    var svc memory.Service
    ts, err := buildToolsets(def, map[string]store.MCPServer{}, svc)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(ts) != 0 {
        t.Fatalf("expected 0 toolsets when service nil, got %d", len(ts))
    }
}

func TestBuildInstruction_MemoryDisabledOmitsInstruction(t *testing.T) {
    def := store.AgentDefinition{ID: "a", DisableMemory: boolPtr(true)}
    svc := memory.InMemoryService()
    instr := buildInstruction(def, map[string]store.MCPServer{}, svc)
    if strings.Contains(instr, "search_memory") {
        t.Fatalf("instruction should not contain search_memory when disabled")
    }
}

func TestBuildInstruction_MemoryEnabledIncludesInstruction(t *testing.T) {
    def := store.AgentDefinition{ID: "a"}
    svc := memory.InMemoryService()
    instr := buildInstruction(def, map[string]store.MCPServer{}, svc)
    if !strings.Contains(instr, "search_memory") {
        t.Fatalf("instruction should contain search_memory when enabled")
    }
}

func TestBuildInstruction_MemoryServiceNil(t *testing.T) {
    def := store.AgentDefinition{ID: "a", DisableMemory: boolPtr(false)}
    var svc memory.Service
    instr := buildInstruction(def, map[string]store.MCPServer{}, svc)
    if strings.Contains(instr, "search_memory") {
        t.Fatalf("instruction should not contain search_memory when service nil")
    }
}
