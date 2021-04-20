module cotton-example

go 1.16

require (
	github.com/onsi/ginkgo v1.16.1 // indirect
	github.com/onsi/gomega v1.11.0 // indirect
	github.com/tonny-zhang/cotton v0.3.0
	github.com/tonny-zhang/cotton-session v0.0.0-00010101000000-000000000000
)

replace (
	github.com/tonny-zhang/cotton => ../cotton
	github.com/tonny-zhang/cotton-session => ../cotton-session
)
