package twobucket

// Solve implements solving of buckets problem. It means pouring water to two buckets of different sizes and
// measuring certain amount of water in one of them. Output is which bucket contains the desired amount, how many
// actions did it take, how much water is in the other bucket and eventual error of operation.
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	var actions int
	var otherBucket string
	buckets := map[string]int{
		"one": 0,
		"two": 0,
	}

	fillBucket := func(bucket string, buckets map[string]int) {
		bu
	}
	// First action always same - fill start bucket.
	switch startBucket {
	case "one":
		buckets[startBucket] = sizeBucketOne
		otherBucket = "two"
	case "two":
		buckets[startBucket] = sizeBucketTwo
		otherBucket = "one"
	}

	// Second action always same - pour start budket to the other until start is empty or other full.
	if buckets[startBucket] > buckets[otherBucket] {
		buckets[startBucket] -= sizeBucketTwo
		buckets[otherBucket] = sizeBucketOne
	} else {
		buckets[startBucket] = 0
		buckets[otherBucket] =
	}
	for buckets["one"] != goalAmount && buckets["two"] != goalAmount {
		actions++
	}
	// fmt.Println(actions, bucketOne)
	return "", actions, 0, nil
}
