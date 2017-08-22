using Go = import "/go.capnp";

@0x9195d073cb5c5953;

$Go.package("capnp");
$Go.import("github.com/disorganizer/brig/model/nodes/capnp");


struct Person  $Go.doc("Person might be any brig user") {
    ident @0 :Text;
    hash  @1 :Data;
}

struct Commit $Go.doc("Commit is a set of changes to nodes") {
    # Following attributes will be part of the hash:
    message @0 :Text;
    author  @1 :Person;
    parent  @2 :Data;     # Hash to parent.
    root    @3 :Data;     # Hash to root directory.

    # Attributes not being part of the hash:
    merge :group {
        isMerge @4 :Bool;
        with    @5 :Person;
        hash    @6 :Data;
    }
}

struct DirEntry $Go.doc("A single directory entry") {
    name @0 :Text;
    hash @1 :Data;
}

struct Directory $Go.doc("Directory contains one or more directories or files") {
    size     @0 :UInt64;
    parent   @1 :Text;
    children @2 :List(DirEntry);
}

struct File $Go.doc("A leaf node in the MDAG") {
    size     @0 :UInt64;
    parent   @1 :Text;
    key      @2 :Data;

    # Hash to the content in ipfs.
    content  @4 :Data;
}

struct Ghost $Go.doc("Ghost indicates that a certain node was at this path once") {
    # Refernce to where the successor of this ghost lives now
    # (In case of moved files, the ghost lives where the old file was)
    movedTo @0 :Data;

    union {
        commit    @1 :Commit;
        directory @2 :Directory;
        file      @3 :File;
    }
}

struct Node $Go.doc("Node is a node in the merkle dag of brig") {
    name    @0 :Text;
    hash    @1 :Data;
    modTime @2 :Text;     # Time as ISO8601
    inode   @3 :UInt64;

    union {
        commit    @4 :Commit;
        directory @5 :Directory;
        file      @6 :File;
        ghost     @7 :Ghost;
    }
}
