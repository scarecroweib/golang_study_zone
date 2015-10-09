package functions

func Add(args1 int32, args2 int32) int32 {
	return args1 + args2
}

func SumAndProduct(A, B int32) (int32, int32) {
	return A + B, A * B
}

func SumAndProduct2(A, B int32) (sum int32, pro int32) {
	sum = A + B
	pro = A * B
	return
}
