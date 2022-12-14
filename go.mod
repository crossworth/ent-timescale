module ent-timescale

go 1.19

require (
	entgo.io/ent v0.11.4
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.7
	github.com/rs/xid v1.4.0
)

require (
	ariga.io/atlas v0.8.3-0.20221116151337-9e4e9cbf3baf // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/text v0.3.7 // indirect
)

// https://github.com/crossworth/ent/tree/timescale-support
replace entgo.io/ent => github.com/crossworth/ent v0.10.1-0.20221123205751-28352eaf5492
