// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Commit is a set of changes to nodes
type Commit struct{ capnp.Struct }
type Commit_merge Commit

// Commit_TypeID is the unique identifier for the type Commit.
const Commit_TypeID = 0x8da013c66e545daf

func NewCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return Commit{st}, err
}

func NewRootCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6})
	return Commit{st}, err
}

func ReadRootCommit(msg *capnp.Message) (Commit, error) {
	root, err := msg.RootPtr()
	return Commit{root.Struct()}, err
}

func (s Commit) String() string {
	str, _ := text.Marshal(0x8da013c66e545daf, s.Struct)
	return str
}

func (s Commit) Message() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Commit) HasMessage() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Commit) MessageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Commit) SetMessage(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Commit) Author() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Commit) HasAuthor() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Commit) AuthorBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Commit) SetAuthor(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Commit) Parent() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Commit) HasParent() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Commit) SetParent(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s Commit) Root() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Commit) HasRoot() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Commit) SetRoot(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s Commit) Index() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Commit) SetIndex(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Commit) Merge() Commit_merge { return Commit_merge(s) }

func (s Commit_merge) With() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Commit_merge) HasWith() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Commit_merge) WithBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Commit_merge) SetWith(v string) error {
	return s.Struct.SetText(4, v)
}

func (s Commit_merge) Head() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return []byte(p.Data()), err
}

func (s Commit_merge) HasHead() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Commit_merge) SetHead(v []byte) error {
	return s.Struct.SetData(5, v)
}

// Commit_List is a list of Commit.
type Commit_List struct{ capnp.List }

// NewCommit creates a new list of Commit.
func NewCommit_List(s *capnp.Segment, sz int32) (Commit_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 6}, sz)
	return Commit_List{l}, err
}

func (s Commit_List) At(i int) Commit { return Commit{s.List.Struct(i)} }

func (s Commit_List) Set(i int, v Commit) error { return s.List.SetStruct(i, v.Struct) }

func (s Commit_List) String() string {
	str, _ := text.MarshalList(0x8da013c66e545daf, s.List)
	return str
}

// Commit_Promise is a wrapper for a Commit promised by a client call.
type Commit_Promise struct{ *capnp.Pipeline }

func (p Commit_Promise) Struct() (Commit, error) {
	s, err := p.Pipeline.Struct()
	return Commit{s}, err
}

func (p Commit_Promise) Merge() Commit_merge_Promise { return Commit_merge_Promise{p.Pipeline} }

// Commit_merge_Promise is a wrapper for a Commit_merge promised by a client call.
type Commit_merge_Promise struct{ *capnp.Pipeline }

func (p Commit_merge_Promise) Struct() (Commit_merge, error) {
	s, err := p.Pipeline.Struct()
	return Commit_merge{s}, err
}

// A single directory entry
type DirEntry struct{ capnp.Struct }

// DirEntry_TypeID is the unique identifier for the type DirEntry.
const DirEntry_TypeID = 0x8b15ee76774b1f9d

func NewDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func NewRootDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func ReadRootDirEntry(msg *capnp.Message) (DirEntry, error) {
	root, err := msg.RootPtr()
	return DirEntry{root.Struct()}, err
}

func (s DirEntry) String() string {
	str, _ := text.Marshal(0x8b15ee76774b1f9d, s.Struct)
	return str
}

func (s DirEntry) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s DirEntry) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DirEntry) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s DirEntry) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s DirEntry) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s DirEntry) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DirEntry) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

// DirEntry_List is a list of DirEntry.
type DirEntry_List struct{ capnp.List }

// NewDirEntry creates a new list of DirEntry.
func NewDirEntry_List(s *capnp.Segment, sz int32) (DirEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return DirEntry_List{l}, err
}

func (s DirEntry_List) At(i int) DirEntry { return DirEntry{s.List.Struct(i)} }

func (s DirEntry_List) Set(i int, v DirEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s DirEntry_List) String() string {
	str, _ := text.MarshalList(0x8b15ee76774b1f9d, s.List)
	return str
}

// DirEntry_Promise is a wrapper for a DirEntry promised by a client call.
type DirEntry_Promise struct{ *capnp.Pipeline }

func (p DirEntry_Promise) Struct() (DirEntry, error) {
	s, err := p.Pipeline.Struct()
	return DirEntry{s}, err
}

// Directory contains one or more directories or files
type Directory struct{ capnp.Struct }

// Directory_TypeID is the unique identifier for the type Directory.
const Directory_TypeID = 0xe24c59306c829c01

func NewDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Directory{st}, err
}

func NewRootDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Directory{st}, err
}

func ReadRootDirectory(msg *capnp.Message) (Directory, error) {
	root, err := msg.RootPtr()
	return Directory{root.Struct()}, err
}

func (s Directory) String() string {
	str, _ := text.Marshal(0xe24c59306c829c01, s.Struct)
	return str
}

func (s Directory) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Directory) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Directory) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Directory) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Directory) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Directory) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Directory) Children() (DirEntry_List, error) {
	p, err := s.Struct.Ptr(1)
	return DirEntry_List{List: p.List()}, err
}

func (s Directory) HasChildren() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Directory) SetChildren(v DirEntry_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewChildren sets the children field to a newly
// allocated DirEntry_List, preferring placement in s's segment.
func (s Directory) NewChildren(n int32) (DirEntry_List, error) {
	l, err := NewDirEntry_List(s.Struct.Segment(), n)
	if err != nil {
		return DirEntry_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Directory_List is a list of Directory.
type Directory_List struct{ capnp.List }

// NewDirectory creates a new list of Directory.
func NewDirectory_List(s *capnp.Segment, sz int32) (Directory_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Directory_List{l}, err
}

func (s Directory_List) At(i int) Directory { return Directory{s.List.Struct(i)} }

func (s Directory_List) Set(i int, v Directory) error { return s.List.SetStruct(i, v.Struct) }

func (s Directory_List) String() string {
	str, _ := text.MarshalList(0xe24c59306c829c01, s.List)
	return str
}

// Directory_Promise is a wrapper for a Directory promised by a client call.
type Directory_Promise struct{ *capnp.Pipeline }

func (p Directory_Promise) Struct() (Directory, error) {
	s, err := p.Pipeline.Struct()
	return Directory{s}, err
}

// A leaf node in the MDAG
type File struct{ capnp.Struct }

// File_TypeID is the unique identifier for the type File.
const File_TypeID = 0x8ea7393d37893155

func NewFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return File{st}, err
}

func NewRootFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return File{st}, err
}

func ReadRootFile(msg *capnp.Message) (File, error) {
	root, err := msg.RootPtr()
	return File{root.Struct()}, err
}

func (s File) String() string {
	str, _ := text.Marshal(0x8ea7393d37893155, s.Struct)
	return str
}

func (s File) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s File) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s File) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s File) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s File) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s File) Key() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s File) HasKey() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s File) SetKey(v []byte) error {
	return s.Struct.SetData(1, v)
}

// File_List is a list of File.
type File_List struct{ capnp.List }

// NewFile creates a new list of File.
func NewFile_List(s *capnp.Segment, sz int32) (File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return File_List{l}, err
}

func (s File_List) At(i int) File { return File{s.List.Struct(i)} }

func (s File_List) Set(i int, v File) error { return s.List.SetStruct(i, v.Struct) }

func (s File_List) String() string {
	str, _ := text.MarshalList(0x8ea7393d37893155, s.List)
	return str
}

// File_Promise is a wrapper for a File promised by a client call.
type File_Promise struct{ *capnp.Pipeline }

func (p File_Promise) Struct() (File, error) {
	s, err := p.Pipeline.Struct()
	return File{s}, err
}

// Ghost indicates that a certain node was at this path once
type Ghost struct{ capnp.Struct }
type Ghost_Which uint16

const (
	Ghost_Which_commit    Ghost_Which = 0
	Ghost_Which_directory Ghost_Which = 1
	Ghost_Which_file      Ghost_Which = 2
)

func (w Ghost_Which) String() string {
	const s = "commitdirectoryfile"
	switch w {
	case Ghost_Which_commit:
		return s[0:6]
	case Ghost_Which_directory:
		return s[6:15]
	case Ghost_Which_file:
		return s[15:19]

	}
	return "Ghost_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Ghost_TypeID is the unique identifier for the type Ghost.
const Ghost_TypeID = 0x80c828d7e89c12ea

func NewGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Ghost{st}, err
}

func NewRootGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Ghost{st}, err
}

func ReadRootGhost(msg *capnp.Message) (Ghost, error) {
	root, err := msg.RootPtr()
	return Ghost{root.Struct()}, err
}

func (s Ghost) String() string {
	str, _ := text.Marshal(0x80c828d7e89c12ea, s.Struct)
	return str
}

func (s Ghost) Which() Ghost_Which {
	return Ghost_Which(s.Struct.Uint16(8))
}
func (s Ghost) GhostInode() uint64 {
	return s.Struct.Uint64(0)
}

func (s Ghost) SetGhostInode(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Ghost) GhostPath() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Ghost) HasGhostPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Ghost) GhostPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Ghost) SetGhostPath(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Ghost) Commit() (Commit, error) {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != commit")
	}
	p, err := s.Struct.Ptr(1)
	return Commit{Struct: p.Struct()}, err
}

func (s Ghost) HasCommit() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetCommit(v Commit) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Ghost) NewCommit() (Commit, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) Directory() (Directory, error) {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != directory")
	}
	p, err := s.Struct.Ptr(1)
	return Directory{Struct: p.Struct()}, err
}

func (s Ghost) HasDirectory() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetDirectory(v Directory) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Ghost) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) File() (File, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != file")
	}
	p, err := s.Struct.Ptr(1)
	return File{Struct: p.Struct()}, err
}

func (s Ghost) HasFile() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetFile(v File) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Ghost) NewFile() (File, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// Ghost_List is a list of Ghost.
type Ghost_List struct{ capnp.List }

// NewGhost creates a new list of Ghost.
func NewGhost_List(s *capnp.Segment, sz int32) (Ghost_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return Ghost_List{l}, err
}

func (s Ghost_List) At(i int) Ghost { return Ghost{s.List.Struct(i)} }

func (s Ghost_List) Set(i int, v Ghost) error { return s.List.SetStruct(i, v.Struct) }

func (s Ghost_List) String() string {
	str, _ := text.MarshalList(0x80c828d7e89c12ea, s.List)
	return str
}

// Ghost_Promise is a wrapper for a Ghost promised by a client call.
type Ghost_Promise struct{ *capnp.Pipeline }

func (p Ghost_Promise) Struct() (Ghost, error) {
	s, err := p.Pipeline.Struct()
	return Ghost{s}, err
}

func (p Ghost_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Ghost_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Ghost_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

// Node is a node in the merkle dag of brig
type Node struct{ capnp.Struct }
type Node_Which uint16

const (
	Node_Which_commit    Node_Which = 0
	Node_Which_directory Node_Which = 1
	Node_Which_file      Node_Which = 2
	Node_Which_ghost     Node_Which = 3
)

func (w Node_Which) String() string {
	const s = "commitdirectoryfileghost"
	switch w {
	case Node_Which_commit:
		return s[0:6]
	case Node_Which_directory:
		return s[6:15]
	case Node_Which_file:
		return s[15:19]
	case Node_Which_ghost:
		return s[19:24]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xa629eb7f7066fae3

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 7})
	return Node{st}, err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 7})
	return Node{st}, err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.RootPtr()
	return Node{root.Struct()}, err
}

func (s Node) String() string {
	str, _ := text.Marshal(0xa629eb7f7066fae3, s.Struct)
	return str
}

func (s Node) Which() Node_Which {
	return Node_Which(s.Struct.Uint16(8))
}
func (s Node) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Node) TreeHash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Node) HasTreeHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Node) SetTreeHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Node) ModTime() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Node) HasModTime() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Node) ModTimeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Node) SetModTime(v string) error {
	return s.Struct.SetText(2, v)
}

func (s Node) Inode() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node) SetInode(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Node) ContentHash() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Node) HasContentHash() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetContentHash(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s Node) User() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Node) HasUser() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Node) UserBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Node) SetUser(v string) error {
	return s.Struct.SetText(4, v)
}

func (s Node) Commit() (Commit, error) {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != commit")
	}
	p, err := s.Struct.Ptr(5)
	return Commit{Struct: p.Struct()}, err
}

func (s Node) HasCommit() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetCommit(v Commit) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Node) NewCommit() (Commit, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Directory() (Directory, error) {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != directory")
	}
	p, err := s.Struct.Ptr(5)
	return Directory{Struct: p.Struct()}, err
}

func (s Node) HasDirectory() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetDirectory(v Directory) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Node) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) File() (File, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != file")
	}
	p, err := s.Struct.Ptr(5)
	return File{Struct: p.Struct()}, err
}

func (s Node) HasFile() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetFile(v File) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Node) NewFile() (File, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Ghost() (Ghost, error) {
	if s.Struct.Uint16(8) != 3 {
		panic("Which() != ghost")
	}
	p, err := s.Struct.Ptr(5)
	return Ghost{Struct: p.Struct()}, err
}

func (s Node) HasGhost() bool {
	if s.Struct.Uint16(8) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetGhost(v Ghost) error {
	s.Struct.SetUint16(8, 3)
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewGhost sets the ghost field to a newly
// allocated Ghost struct, preferring placement in s's segment.
func (s Node) NewGhost() (Ghost, error) {
	s.Struct.SetUint16(8, 3)
	ss, err := NewGhost(s.Struct.Segment())
	if err != nil {
		return Ghost{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) BackendHash() ([]byte, error) {
	p, err := s.Struct.Ptr(6)
	return []byte(p.Data()), err
}

func (s Node) HasBackendHash() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s Node) SetBackendHash(v []byte) error {
	return s.Struct.SetData(6, v)
}

// Node_List is a list of Node.
type Node_List struct{ capnp.List }

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 7}, sz)
	return Node_List{l}, err
}

func (s Node_List) At(i int) Node { return Node{s.List.Struct(i)} }

func (s Node_List) Set(i int, v Node) error { return s.List.SetStruct(i, v.Struct) }

func (s Node_List) String() string {
	str, _ := text.MarshalList(0xa629eb7f7066fae3, s.List)
	return str
}

// Node_Promise is a wrapper for a Node promised by a client call.
type Node_Promise struct{ *capnp.Pipeline }

func (p Node_Promise) Struct() (Node, error) {
	s, err := p.Pipeline.Struct()
	return Node{s}, err
}

func (p Node_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Node_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Node_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Node_Promise) Ghost() Ghost_Promise {
	return Ghost_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

const schema_9195d073cb5c5953 = "x\xda\xb4V_h\x1c\xd5\x17>\xdf\xbd3;\x9d\x92" +
	"\xfe6\xfb\x9b\x14J\xf9\xa5{)-\xbf\xb4\xc46i" +
	"*6\xa1\xd2\xa6MlR\xd3\x92\xdbM\xb1-U\x98" +
	"\xee\xde\xec\x0c\xdd\x9dIg\xa6\xc6\x88\xa1\"\x0aU\xa9" +
	"\xb4X\xc1B\x8a\xad\xc4\x7f\xa0\xd4\x17\x1fE\x10\x14\xd1" +
	"\x17\xf1A\xc1G\xf5A\x14\xf4\xd5\x8a\xed\xc8\xdd\xbfi" +
	"\x88\xb6\x08\xbee\xbes\xee\xbd\xdf\xf9\xcewN\xb6\xef" +
	"W\xb6\x97\xf5\x9b\xff7\x88d\x9f\x99I\x7f\xfa\xef\xc2" +
	"\x8f\xdf\xf4|\xf64\xc9\x0d`i\xe1\xf8\xc9/\xe2/" +
	"_\xb9D\xa3\xcc\xe20\x06nb#\x1c\x9bY\x8e\xcd" +
	"\xf2\x03\xa3,\x0fBz5\xff\xf0\xec\xe3\xbf\xac}\x91" +
	"r\x1b\xd0>`2\x8bh\xc0\xe5Cp\xcep\xcb9" +
	"\xc3\xf3\xceU>KHo<:\x15|\xea\\\xbb\xa0" +
	"\x1fX\x9a\x9f\xd1\xf97\xf9V8\xb6a9\xb6\x91\x1f" +
	"\x184\x1e\xd1\xf7\x1f\xed\x7f\xfe\x81\x07\x07\xdfzi\xf9" +
	"\x81\xda\x03\xbe\xb9\x1e\xce\x9ci9sf\xde\xb9n\xde" +
	" \xa4\xdf\xff>=s\xee\xe7-o.\xaf\xc0\xb2\x0c" +
	"\x18\x03\x83\x99\xf5p\xc63\x963\x9e\xc9\x0f\xccgB" +
	"FH\x17\x7f\x98\xf86\xbb\xf8\xdbG$7c\x09\xc1" +
	"\xb5\x19\x0bD\x03\xa6}\x02\x04'gk\xf6Xx\xa6" +
	"\xd2w|\xe2\xbb\x15\xc9T\xed}p\xe6m\xcb\x99\xb7" +
	"\xf3\xce\x07\xf6\x0d\xda\x95\x16\xddd:\xde\x1e\x84\xbc\xa4" +
	"\xe2\xedEw&\x98\xd9\x1e\x84%\x15o\xab\xfd=t" +
	"\xc0\xb3\xc28\x99\x04\xa4\x01\x96>\xf6\xf2k\xf2\xc3\xaf" +
	"_\xf8\x84\xa4\xc10\xdc\x0bt\x10\xf5\xe3+\xa4\x07\xbc" +
	"0N\x84\x1fdJ~\xd1MT,\x12\xcfM\x84+" +
	"\x8a*J\\?\x10\xfaJ1\xeb\xc6\xc2MD\xe2\xf9" +
	"\xb1\x98q\x13O\x84A\x11\x8aHvq\x83\xc8\x00Q" +
	"n\xfe\x04\x91|\x8aC\x9eg\x00\xba\xa0\xb1\xe7\x8e\x10" +
	"\xc9g9\xe4E\x86n\x96\xa6\xe8\x02#\xca]\x18\"" +
	"\x92\xe79\xe4e\x86n~[\xc3\x9c(wIg_" +
	"\xe4\x90\x0b\x0c\xdd\xc6-\x0d\x1bD\xb9+[\x89\xe4e" +
	"\x0ey\x8d!-k\xb6\xe3AH\xbc\xa4`\x13\x83M" +
	"\x0dp\xd2M\x08\x1e:\x88\xa1\x83\xb0\xa7\x18V\xab~" +
	"\x82\xce\xb6\xe4\x04t\x12\xd2\x92\x1f\xa9b\x12F\x849" +
	"t\xb65\xafG\xb3\xd3~E\xa1\xb3\xed\x8b\xc6\xa1\xbb" +
	"H=\xe2\xef\x89F\x83$\x9a[Y\xed\xff\xd5\xd4\xce" +
	"\xe1\xf3tX\xc4~P\xae(&\x9a4\xe6\x84\xd2\x07" +
	"\x09rUK\xca-\xba\xe2M\x1c\xb2\x8f!\xd7\xd4\xf2" +
	">\x0d\xf6p\xc8\x9d\x0c\xd9\xc0\xad\xaaf\xa9Y\xcf\x8d" +
	"=\xac!\x865wg\xba?\xccj]V\xe6)\x1a" +
	"\xae\xd8\x88t\x7fM>\xe1\xf3X\xb8\"V\x89\x08\xa7" +
	"E\xd1s\x83\xb26H(\x82\xd0*\xa9\x98H\xaek" +
	"\x91\xbe\xb2\xaf\xdd\xa6\x16\xe9\xab\xba\xd3\xafr\xc8E\x86" +
	"\x1cc\xf5\xf6_\xd7\xe0\x02\x87|\x9b!\xc7y\xbd\xf9" +
	"o\xe8\xf2\xaeq\xc8w\x19`\xd4;\xff\xce\x0e\"\xb9" +
	"\xc8!\xdfg\x80\x89%\xc3\x94{o\x07\xb1sU\x15" +
	"\xc7n\xb9%\xc4\x1e\xf7l\xe2\x85Q\xebs\xc6\x8dT" +
	"\x904\x95\xc9Fa\xd8\xfa\xc8\xfbAI=\x01\x93\x18" +
	"LB\xbe\xaa\xa2\xb2\xba\x9bt\x0f\xf9\xbc\xa2V\x16n" +
	"]\xa3\xc1\x1f\xa7\xc3\xa2\xa2\xdci\x110=5~ " +
	"\x12O\x89C#\xc3\x07\x88Hv\xb4\xb4\x1a\xd5\xc5\xee" +
	"\xe5\x90\x13\xedY\x19\xd7\xaa\x8cp\xc8I-UcR" +
	"\x0em$\x92c\x1cr\x8a!\x1b\xfbO\xb6<\xdf," +
	"\xaeQ\xabuZ\xcd\xdd\xab\x05\x0e\xeb\xc0\xcaulj" +
	"\x18\xe0 \xd2\xc3\xb5\x02ba\xb8\xf5\x0d\xd0\xa8\xa5\xaa" +
	"\xa2\xd3\x15%JnY;\xe2T\xe4\x97\x09\xb2\xb7Y" +
	"\x98\xb3\x19[\x89\x0a\x02\x1c\x85^\xb4}\xe0l\xc1A" +
	"\xa2B\x8f\xc6w\xa2m\x05\xa7\x1f\xfb\x88\x0a\xbd\x1a\xdf" +
	"\x05\x06\xd4\xcd\xe0\xdc\x8f\x1dD\x85>\x0d\xef\xd6\xe9\x06" +
	"\xaf\x19\xc2\x19\xc4)\xa2\xc2.\x8d\x8fh\xdc4\xba`" +
	"\x129\xc3\xb5gwk|\x0c\x0c\xdd\x9945\xbb\x90" +
	"!rF1DT\xd8\xab#\x13:b\xdd\xd6\x11\x8b" +
	"\xc8\x19\xc7\x11\xa2\xc2\x98\x8eL\xe9\xc8\xaa[:\xb2\x8a" +
	"\xc8\x91\xb5\xdb&t\xe4\x98\x8e\xd8\x7f\xe8\x88M\xe4\x1c" +
	"\xad\xf1\x9a\xd4\x91\x93\xfa\xfd\xd5\x99.\xac&r\x8e\xd7" +
	"x\x1d\xd3x\x09\xcb\xc63M\"\xa5\xc6\xdc\xd8#\xa2" +
	"f\x8b\xceU\xc3\xd2\x94\xdf\xce\xc9\xfbZ\xe3\xd6>+" +
	"\x86A\xa2\x82d\x8c\xac%\x93\x9d=\x1b\xab\xe8\xdfY" +
	"o\xf9\xda\x02Eg\xfb\x1ft\xe3\xb2Sn\xf1\xb4\x0a" +
	"Jw\x12i\xf9\xcb\xf8\xab\x15\xa3\xa9m\xab\xaa\x88\x97" +
	"\x95\xdej\x9d\xf5.-[k\xf5\x06\xdd\xb9\xd6f\xfd" +
	"\xc4k\xaf5\xe5\x96\xee\xf5\xcd\x91\xe66\xa5\x95\x8d\xdd" +
	"\xd30\xf6\xebH\x9b\xa9\xe6\x9c\xd0:\xbb~\x10\x8b0" +
	"P\"\x8cD5\x8cTk1\xfb*\xd6\xd8\xb4oU" +
	"j\x9b\xee\x9fL\xefA\"9\xc1!\xbd\xbf\x9f\xde\xb4" +
	"\xe8\xf9\x95R\xa4\x02m\x91\xff\x10&9\xd0\xd9\xfe\xed" +
	"C\xd0\xe0\x9f\x01\x00\x00\xff\xff\x0e\xda3\x93"

func init() {
	schemas.Register(schema_9195d073cb5c5953,
		0x80c828d7e89c12ea,
		0x8b15ee76774b1f9d,
		0x8da013c66e545daf,
		0x8ea7393d37893155,
		0xa629eb7f7066fae3,
		0xbff8a40fda4ce4a4,
		0xe24c59306c829c01)
}
