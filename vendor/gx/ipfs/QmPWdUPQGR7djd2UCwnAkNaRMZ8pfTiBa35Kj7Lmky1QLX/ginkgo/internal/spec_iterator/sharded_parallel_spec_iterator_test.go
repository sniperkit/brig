package spec_iterator_test

import (
	. "gx/ipfs/QmPWdUPQGR7djd2UCwnAkNaRMZ8pfTiBa35Kj7Lmky1QLX/ginkgo/internal/spec_iterator"

	"gx/ipfs/QmPWdUPQGR7djd2UCwnAkNaRMZ8pfTiBa35Kj7Lmky1QLX/ginkgo/internal/codelocation"
	"gx/ipfs/QmPWdUPQGR7djd2UCwnAkNaRMZ8pfTiBa35Kj7Lmky1QLX/ginkgo/internal/containernode"
	"gx/ipfs/QmPWdUPQGR7djd2UCwnAkNaRMZ8pfTiBa35Kj7Lmky1QLX/ginkgo/internal/leafnodes"
	"gx/ipfs/QmPWdUPQGR7djd2UCwnAkNaRMZ8pfTiBa35Kj7Lmky1QLX/ginkgo/internal/spec"
	"gx/ipfs/QmPWdUPQGR7djd2UCwnAkNaRMZ8pfTiBa35Kj7Lmky1QLX/ginkgo/types"

	. "github.com/onsi/gomega"
	. "gx/ipfs/QmPWdUPQGR7djd2UCwnAkNaRMZ8pfTiBa35Kj7Lmky1QLX/ginkgo"
)

var _ = Describe("ShardedParallelSpecIterator", func() {
	var specs []*spec.Spec
	var iterator *ShardedParallelIterator

	newSpec := func(text string, flag types.FlagType) *spec.Spec {
		subject := leafnodes.NewItNode(text, func() {}, flag, codelocation.New(0), 0, nil, 0)
		return spec.New(subject, []*containernode.ContainerNode{}, false)
	}

	BeforeEach(func() {
		specs = []*spec.Spec{
			newSpec("A", types.FlagTypePending),
			newSpec("B", types.FlagTypeNone),
			newSpec("C", types.FlagTypeNone),
			newSpec("D", types.FlagTypeNone),
		}
		specs[3].Skip()

		iterator = NewShardedParallelIterator(specs, 2, 1)
	})

	It("should report the total number of specs", func() {
		Ω(iterator.NumberOfSpecsPriorToIteration()).Should(Equal(4))
	})

	It("should report the number to be processed", func() {
		n, known := iterator.NumberOfSpecsToProcessIfKnown()
		Ω(n).Should(Equal(2))
		Ω(known).Should(BeTrue())
	})

	It("should report the number that will be run", func() {
		n, known := iterator.NumberOfSpecsThatWillBeRunIfKnown()
		Ω(n).Should(Equal(1))
		Ω(known).Should(BeTrue())
	})

	Describe("iterating", func() {
		It("should return the specs in order", func() {
			Ω(iterator.Next()).Should(Equal(specs[0]))
			Ω(iterator.Next()).Should(Equal(specs[1]))
			spec, err := iterator.Next()
			Ω(spec).Should(BeNil())
			Ω(err).Should(MatchError(ErrClosed))
		})
	})
})
