// Package tally counts the TAG WORDs of the live documents of pack archives by
// Issuer and code: the Length bits (LLLL) are dropped and the 9-bit code w>>4
// (section KKKK and C, TAG ID 0x000L~0x1FFL) is the key, nothing else folded.
package tally
