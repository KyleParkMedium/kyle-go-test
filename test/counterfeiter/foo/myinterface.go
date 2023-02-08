package foo

// You only need **one** of these per package!

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate .

// -o mock/test.go --fake-name Test . CoinProcessor

// // You will add lots of directives like these in the same package...
// //
// //counterfeiter:generate . MySpecialInterface
// type MySpecialInterface interface {
// 	DoThings(string, uint64) (int, error)
// }

// // Like this...
// //
// //counterfeiter:generate . MyOtherInterface
// type MyOtherInterface interface {
// 	DoOtherThings(string, uint64) (int, error)
// }

//counterfeiter:generate . Something
type Something interface {
	DoThings(string, uint64) (int, error)
	DoNothing()
	DoASlice([]byte)
	DoAnArray([4]byte)
}
