package qid

import (
	"hash/fnv"
	"os"
	"syscall"
)

// FileInfo creates a neinp.Qid from os.FileInfo, using Sys of FileInfo if possible.
func FileInfo(fi os.FileInfo) Qid {
	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return Generic(fi)
	}

	t := TypeFile
	if fi.IsDir() {
		t = TypeDir
	}

	// fvn hashes match the size of version and path fields. they are no cryptographical hashes, but should work good enough.
	v := fnv.New32a()
	v.Write([]byte{uint8(s.Mtim.Sec), uint8(s.Mtim.Sec >> 8), uint8(s.Mtim.Sec >> 16), uint8(s.Mtim.Sec >> 24), uint8(s.Mtim.Sec >> 32), uint8(s.Mtim.Sec >> 40), uint8(s.Mtim.Sec >> 48), uint8(s.Mtim.Sec >> 56)})
	v.Write([]byte{uint8(s.Mtim.Nsec), uint8(s.Mtim.Nsec >> 8), uint8(s.Mtim.Nsec >> 16), uint8(s.Mtim.Nsec >> 24), uint8(s.Mtim.Nsec >> 32), uint8(s.Mtim.Nsec >> 40), uint8(s.Mtim.Nsec >> 48), uint8(s.Mtim.Nsec >> 56)})

	p := fnv.New64a()
	p.Write([]byte{uint8(s.Dev), uint8(s.Dev >> 8), uint8(s.Dev >> 16), uint8(s.Dev >> 24), uint8(s.Dev >> 32), uint8(s.Dev >> 40), uint8(s.Dev >> 48), uint8(s.Dev >> 56)})
	p.Write([]byte{uint8(s.Ino), uint8(s.Ino >> 8), uint8(s.Ino >> 16), uint8(s.Ino >> 24), uint8(s.Ino >> 32), uint8(s.Ino >> 40), uint8(s.Ino >> 48), uint8(s.Ino >> 56)})

	return Qid{Type: t, Version: v.Sum32(), Path: p.Sum64()}
}

// Generic creates a neinp.Qid not using FileInfo.Sys().
func Generic(fi os.FileInfo) Qid {
	t := TypeFile
	if fi.IsDir() {
		t = TypeDir
	}

	sec := fi.ModTime().Unix()
	nsec := fi.ModTime().UnixNano()

	v := fnv.New32a()
	v.Write([]byte{uint8(sec), uint8(sec >> 8), uint8(sec >> 16), uint8(sec >> 24), uint8(sec >> 32), uint8(sec >> 40), uint8(sec >> 48), uint8(sec >> 56)})
	v.Write([]byte{uint8(nsec), uint8(nsec >> 8), uint8(nsec >> 16), uint8(nsec >> 24), uint8(nsec >> 32), uint8(nsec >> 40), uint8(nsec >> 48), uint8(nsec >> 56)})

	p := fnv.New64a()
	p.Write([]byte(fi.Name()))

	return Qid{Type: t, Version: v.Sum32(), Path: p.Sum64()}
}
