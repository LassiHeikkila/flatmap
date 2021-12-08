module github.com/LassiHeikkila/flatmap

go 1.14

require github.com/google/go-cmp v0.5.6 // indirect

retract (
	[v1.0.0, v1.0.10] // relic from original Terraform repo
	v1.0.11 // contains just the retractions
	v1.0.12 // contains just the retractions
)
