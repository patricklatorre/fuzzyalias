package main

import "testing"

func Benchmark_FindNearestLink(b *testing.B) {
	config.Aliases = map[string]string{
		"/foo": "https://foo.com",
		"/bar": "https://bar.com",
		"/baz": "https://baz.com",
	}

	var (
		exactQuery   = "/baz"
		prefixQuery  = "/ba"
		suffixQuery  = "/az"
		noMatchQuery = "/az"
		bigQuery     = "/az/cont/inued?n=99999"
	)

	b.Run("Exact match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			searchNearestLink(exactQuery)
		}
	})

	b.Run("Prefix match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			searchNearestLink(prefixQuery)
		}
	})

	b.Run("Suffix match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			searchNearestLink(suffixQuery)
		}
	})

	b.Run("No match", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			searchNearestLink(noMatchQuery)
		}
	})

	b.Run("Big query", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			searchNearestLink(bigQuery)
		}
	})
}
