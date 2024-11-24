package jocker

import (
	// "context"
	"context"
	"log"
	"os"

	// "os"

	// "github.com/heph2/jocker/internal/builder"
	"github.com/heph2/jocker/internal/parser"
	"github.com/moby/buildkit/client/llb"
	// "github.com/moby/buildkit/client/llb"
)

func DebugDump() error {
	// Initialize Jsonnet VM and evaluate Jockerfile
	jsonStr, err := parser.EvaluateJsonnetFile("Jockerfile")
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON into Jockerfile struct
	j, err := parser.ParseJockerfile(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	// Generate LLB state from Jockerfile
	dt, err := j.ToLLB().Marshal(context.TODO(), llb.LinuxAmd64)
	if err != nil {
		log.Fatal(err)
	}

	// Write LLB definition to stdout
	return llb.WriteTo(dt, os.Stdout)
}
