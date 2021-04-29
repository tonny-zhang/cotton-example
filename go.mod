module cotton-example

go 1.16

require (
	github.com/tonny-zhang/cotton v0.4.0
	github.com/tonny-zhang/cotton-session v0.0.0-00010101000000-000000000000
)

replace (
	github.com/tonny-zhang/cotton => ../cotton
	github.com/tonny-zhang/cotton-session => ../cotton-session
)
