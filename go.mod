module cotton-example

go 1.16

require (
	github.com/tonny-zhang/cotton v0.4.1
	github.com/tonny-zhang/cotton-session v0.0.0-20210429020325-cfeac247336a
)

// config for local
replace (
	github.com/tonny-zhang/cotton => ../cotton
	github.com/tonny-zhang/cotton-session => ../cotton-session
)
