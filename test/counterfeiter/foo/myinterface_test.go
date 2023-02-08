package foo_test

import (
	"counterfeiter/foo/foofakes"
	"errors"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var fake = &foofakes.FakeSomething{}

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Test functions",
	func() {

		// It("can have its behavior configured using stub functions", func() {

		// 	fake.DoThingsStub = func(arg1 string, arg2 uint64) (int, error) {
		// 		Expect(arg1).To(Equal("stuff"))
		// 		Expect(arg2).To(Equal(uint64(5)))
		// 		return 3, errors.New("the-error")
		// 	}

		// 	num, err := fake.DoThings("stuff", 5)

		// 	Expect(num).To(Equal(5))
		// 	Expect(err).To(Equal(errors.New("the-error")))
		// })

		It("can have its return values configured", func() {
			fake.DoThingsReturns(3, errors.New("the-error"))

			num, err := fake.DoThings("stuff", 5)
			Expect(num).To(Equal(3))
			Expect(err).To(Equal(errors.New("the-error")))
		})

		// It("returns zero values when no return value or stub is provided", func() {
		// 	num, err := fake.DoThings("stuff", 5)

		// 	Expect(num).To(Equal(0))
		// 	Expect(err).To(BeNil())
		// })

		// It("records the arguments it was called with", func() {
		// 	Expect(fake.DoThingsCallCount()).To(Equal(0))

		// 	fake.DoThings("stuff", 5)

		// 	Expect(fake.DoThingsCallCount()).To(Equal(1))
		// 	arg1, arg2 := fake.DoThingsArgsForCall(0)
		// 	Expect(arg1).To(Equal("stuff"))
		// 	Expect(arg2).To(Equal(uint64(5)))
		// })

		// It("records an array argument as a copy", func() {
		// 	buffer := [4]byte{1, 2, 3, 4}

		// 	fake.DoAnArray(buffer)

		// 	buffer[0] = 2
		// 	arg1 := fake.DoAnArrayArgsForCall(0)
		// 	Expect(arg1).To(ConsistOf(byte(1), byte(2), byte(3), byte(4)))
		// })

		// It("records a slice argument as a copy", func() {
		// 	buffer := []byte{1}

		// 	buffer[0] = 2

		// 	fake.DoASlice(buffer)

		// 	arg1 := fake.DoASliceArgsForCall(0)
		// 	Expect(arg1).To(ConsistOf(byte(2)))
		// 	Expect(buffer).To(ConsistOf(byte(2)))
		// })

		// It("passes the original slice to a stub function", func() {
		// 	buffer := []byte{3}

		// 	fake.DoASliceStub = func(b []byte) {
		// 		b[0] = 2
		// 	}

		// 	fake.DoASlice(buffer)

		// 	arg1 := fake.DoASliceArgsForCall(0)
		// 	Expect(arg1).To(ConsistOf(byte(3)))
		// 	Expect(buffer).To(ConsistOf(byte(2)))
		// })

		// It("records its calls without race conditions", func() {
		// 	go fake.DoNothing()

		// 	Eventually(fake.DoNothingCallCount, 1.0).Should(Equal(1))
		// })

		// BeforeEach(func() {
		// 	var ifake interface{} = fake
		// 	_, ok := ifake.(InvocationRecorder)

		// 	Expect(ok).To(BeTrue())
		// })

		// It("records each invocation", func() {
		// 	Expect(len(fake.Invocations()["DoThings"])).To(Equal(0))
		// 	Expect(len(fake.Invocations()["DoNothing"])).To(Equal(0))
		// 	Expect(len(fake.Invocations()["DoASlice"])).To(Equal(0))
		// 	Expect(len(fake.Invocations()["DoAnArray"])).To(Equal(0))

		// 	fake.DoThings("hello", 0)
		// 	Expect(len(fake.Invocations()["DoThings"])).To(Equal(1))
		// 	Expect(fake.Invocations()["DoThings"][0][0]).To(Equal("hello"))
		// 	Expect(fake.Invocations()["DoThings"][0][1]).To(Equal(uint64(0)))

		// 	fake.DoNothing()
		// 	Expect(len(fake.Invocations()["DoNothing"])).To(Equal(1))

		// 	fake.DoASlice([]byte("HAI"))
		// 	Expect(len(fake.Invocations()["DoASlice"])).To(Equal(1))
		// 	Expect(fake.Invocations()["DoASlice"][0][0]).To(Equal([]byte("HAI")))

		// 	fake.DoAnArray([4]byte{})
		// 	Expect(len(fake.Invocations()["DoAnArray"])).To(Equal(1))
		// 	Expect(fake.Invocations()["DoAnArray"][0][0]).ToNot(BeNil())
		// })

		// var start1 chan struct{}
		// var start2 chan struct{}
		// var end1 chan struct{}
		// var end2 chan struct{}

		// BeforeEach(func() {
		// 	start1 = make(chan struct{})
		// 	end1 = make(chan struct{})
		// 	start2 = make(chan struct{})
		// 	end2 = make(chan struct{})

		// 	fake.DoNothingStub = func() {
		// 		close(start1)
		// 		end1 <- struct{}{}
		// 	}

		// 	fake.DoThingsStub = func(string, uint64) (int, error) {
		// 		close(start2)
		// 		end2 <- struct{}{}
		// 		return 0, nil
		// 	}

		// 	go fake.DoNothing()
		// 	<-start1
		// 	go fake.DoThings("abc", 1)
		// 	<-start2
		// })

		// AfterEach(func() {
		// 	close(end1)
		// 	close(end2)
		// })

		// It("does not deadlock", func() {
		// 	Eventually(start1).Should(BeClosed())
		// 	Eventually(end1).Should(Receive())

		// 	Eventually(start2).Should(BeClosed())
		// 	Eventually(end2).Should(Receive())
		// })

		// Context("sample", func() {

		// 	fake.DoThings("stuff", 5)
		// 	fake.DoThings("stuff", 4)

		// 	It("name should return name", func() {

		// 		// "stuff", 5

		// 		arg1, arg2 := fake.DoThingsArgsForCall(1)

		// 		Expect(arg1).To(Equal("stuff"))
		// 		Expect(arg2).To(Equal(uint64(4)))

		// 	})
		// 	// Expect(arg2).To(Equal(5))
		// 	// It("name should return name", func() {

		// 	// 	Expect(fake.DoThingsCallCount()).To(Equal(2))
		// 	// })

		// 	// str, num := fake.DoThingsArgsForCall(0)

		// 	// fmt.Println(str)
		// 	// It("name should return name", func() {

		// 	// 	Expect(str).To(Equal("satuff"))
		// 	// })
		// 	// Expect(num).To(Equal(uint64(5)))
		// })

		// // Context("sample", func() {

		// // 	fake.DoThingsReturns(3, errors.New("the-error"))
		// // 	num, err := fake.DoThings("stuff", 5)

		// // 	It("one should return 1", func() {
		// // 		Expect(num).To(Equal(3))
		// // 	})

		// // 	It("one should return 1", func() {
		// // 		Expect(err).To(Equal(errors.New("the-error")))
		// // 	})

		// // })
	},
)

type InvocationRecorder interface {
	Invocations() map[string][][]interface{}
}
