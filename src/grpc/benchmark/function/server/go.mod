module bm

go 1.16

require (
	byfuls.com/generate/proto v0.0.0
	google.golang.org/grpc v1.39.0 // indirect
)

replace byfuls.com/generate/proto v0.0.0 => ../generate/proto
