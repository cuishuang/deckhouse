module github.com/deckhouse/deckhouse/testing

go 1.13

require (
	github.com/flant/shell-operator v1.0.0-beta.7.0.20200206175208-7a40246730f7 // branch: master
	github.com/gammazero/deque v0.0.0-20190521012701-46e4ffb7a622
	github.com/imdario/mergo v0.3.8
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.9.0
	github.com/otiai10/copy v1.0.2
	github.com/segmentio/go-camelcase v0.0.0-20160726192923-7085f1e3c734
	github.com/tidwall/gjson v1.3.4
	github.com/tidwall/sjson v1.0.4
	golang.org/x/sys v0.0.0-20200113162924-86b910548bc1
	gopkg.in/evanphx/json-patch.v4 v4.5.0
	gopkg.in/yaml.v2 v2.2.7
	gopkg.in/yaml.v3 v3.0.0-20191120175047-4206685974f2
	k8s.io/apimachinery v0.17.0
	sigs.k8s.io/yaml v1.1.1-0.20191128155103-745ef44e09d6 // branch master, with fixes in yaml.v2.2.7
)

//replace github.com/flant/shell-operator => ../../shell-operator
