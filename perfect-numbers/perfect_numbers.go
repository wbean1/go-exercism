/*
	Package perfect provides a means to classify positive integers
	into the distict classification groups of 'perfect', 'abundant',
	or 'deficient' based on its sum of factors.
*/
package perfect

import "errors"

type Classification struct {
	name string
}

var ClassificationPerfect Classification = Classification{"ClassificationPerfect"}
var ClassificationAbundant Classification = Classification{"ClassificationAbundant"}
var ClassificationDeficient Classification = Classification{"ClassificationDeficient"}
var ErrOnlyPositive error = errors.New("perfect.Classify: input must be positive integer")

// Classify takes a positive non-zero integer, and returns an appropriate
// Classification.
func Classify(i int64) (Classification, error) {
	if i <= 0 {
		return Classification{""}, ErrOnlyPositive
	}
	var sum int64
	for j := int64(1); j <= i/2; j++ {
		if i%j == 0 {
			sum += j
		}
	}
	if sum == i {
		return ClassificationPerfect, nil
	} else if sum > i {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}
}
