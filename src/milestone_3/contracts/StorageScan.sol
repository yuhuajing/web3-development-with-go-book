pragma solidity 0.8.26;

contract StorageScan {
    enum Tx {
        Pending,
        Awaiting,
        Success,
        Failed
    }
    // int type
    int8 private int1 = -8; // 0x0
    Tx txenum; // 0x0 8bit
    int128 private int2 = 128; // 0x0
    int256 private int3 = 256; // 0x1

    // uint type
    uint16 private uint1 = 8; // 0x2
    uint64 private uint2 = 128; // 0x2
    uint128 private uint21 = 128; // 0x2
    uint256 private uint3 = 0x123456789abcef1; // 0x3

    // bool type
    bool private bool1 = true; // 0x4 8bit
    bool private bool2 = false; // 0x4

    // string type
    string private string1 = "abc"; // 0x5
    string private string2 =
        "solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts.solidity is an object-oriented, high-level language for implementing smart contracts."; // 0x6

    // bytes typeva
    bytes1 private b1 = "a"; // 0x7
    bytes8 private b2 = "byte2"; //0x7
    bytes32 private b3 = "string bytes cannot exceed 32"; //0x8

    // address type
    address private addr1 = 0x2729E5DFDeeCB92C884470EF6CaD9e844e34502D; // 0x9

    // struct type
    struct Entity {
        uint64 age; // //0xa
        uint128 id; // //0xa
        string value; // //0xb
    }
    struct Ert {
        string nme;
        mapping(uint256 => mapping(address => uint256)) ty;
    }

    Entity i; // //0xa

    // slice value 动态数组
    uint8[] private slice1 = [1, 2, 3, 4, 5]; // 0xc
    uint256[] private slice2 = [256, 257, 258, 259, 260]; // // 0xd
    bool[] private slice3 = [true, false, false, true, false]; // 0xe
    string[] private slice4 = [
        "abc",
        "solidity is an object-oriented, high-level language for implementing smart contracts."
    ]; //0xf
    Entity[] private slice5; // 0x10

    // array value
    uint8[2][3] private tt = [[1, 2], [4, 5], [6, 7]]; //0x11
    uint8[5] private array1 = [1, 2, 3, 4, 5]; // 0x14
    uint256[5] private array2 = [256, 257, 258, 259, 260]; // // 0x15-0x19
    bool[5] private array3 = [true, false, false, true, false]; // 0x1a
    string[2] private array4 = [
        "abc",
        "solidity is an object-oriented, high-level language for implementing smart contracts."
    ]; //0x1b-0x1c
    Entity[2] private array5; // 0x1d-0x20

    // mapping value
    mapping(uint256 => string) private mapping1; // 0x21
    mapping(string => uint256) private mapping2; // 0x22
    mapping(address => uint256) private mapping3; // 0x23
    mapping(int256 => uint256) private mapping4; // 0x24
    mapping(bytes1 => uint256) private mapping5; // 0x25
    mapping(uint256 => Entity) private mapping6; // 0x26
    mapping(uint256 => mapping(address => uint256)) private mapping7; // 0x27
    Ert ert; //0x28
    Entity ii; // //0xa

    constructor() {
        txenum = Tx.Awaiting;
        i.age = 16;
        i.id = 1;
        i.value = "entity";
        ii = i;
        ii.value = string2;

        slice5.push(Entity(12, 1, "slice50"));
        slice5.push(Entity(23, 2, "slice51"));

        array5[0] = Entity(34, 1, "arry50");
        array5[1] = Entity(18, 2, "array51");
        ert.nme = "Hello";
        ert.ty[1][address(5)] = 9;

        mapping1[1] = "mapping1";
        mapping2["mapping2"] = 1;
        mapping3[0x2729E5DFDeeCB92C884470EF6CaD9e844e34502D] = 1;
        mapping4[-256] = 1;
        mapping5["a"] = 1;
        mapping6[123] = Entity(27, 1, "mapping6");
        mapping7[123][0x2729E5DFDeeCB92C884470EF6CaD9e844e34502D] = 2;
    }
}

