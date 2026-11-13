# do not edit — managed by intent-os (intent: edd6c7ad-7b59-46ee-875f-b42dccc61b36, version: v1)
# Context Summary

**Intent:** Add a generic Set type to the smoke library

## Description

Add a new file set.go to package smoke that declares the generic type alias `type Set[T comparable] = map[T]struct{}` (a generic type alias) plus the generic functions NewSet, SetAdd, SetHas and SetLen over it, and set_test.go with table-driven tests. Do not change calc.go or calc_test.go.
