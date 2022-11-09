package FacadePattern

type BucketItem struct {
	name string
}

type Bucket struct{}

func (Bucket) fill(item BucketItem) {
	println("Filling bucket with", item.name)
}

func (Bucket) pack() {
	println("Packed! Ready to ship.")
}
