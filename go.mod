module github.com/anytypeio/go-anytype-middleware

go 1.12

require (
	github.com/PuerkitoBio/goquery v1.5.0 // indirect
	github.com/anytypeio/go-anytype-library v0.0.0-20191019100520-f545fa654778
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/ipfs/go-log v0.0.1
	github.com/mauidude/go-readability v0.0.0-20141216012317-2f30b1a346f1
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/otiai10/opengraph v1.1.0
	github.com/stretchr/testify v1.3.0
	github.com/textileio/go-textile v0.7.2-0.20190907000013-95a885123536
)

replace github.com/textileio/go-textile => github.com/anytypeio/go-textile v0.0.0-20190924115707-a0dcb5a893ec

replace github.com/libp2p/go-eventbus => github.com/libp2p/go-eventbus v0.1.0
