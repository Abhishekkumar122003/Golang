What is primitive Data Types?
    -First what are data types? Since Go is statically and storngly typed language
      every variable and constant shold have have a type at commpile time.
    -Compiler wont compile the code until the data type is assigned to an variable or a     constant.
    -Primitive data type are the data types from which all the other non primitive data types are constructed.
    -Composite data type are made from basic data types ex:-[homogeneous - Array,list. Hetrogeneous- uct, record or object. Tabular or semy-structured - Table(row, col), list, dictionary/map, tree etc. Object-Orianted Classes ]

Primitive Data Type
1. Boolean - True or False
2. String are used for text in code. There are single line or multi-line strings.
3. Go support unicode strings.
4. Numbers - int, float, complex.
5. The concept of zero values in Go.


Integer
1. In Golang, there are different types of integer such as, "int8 int16, int32, int 64, etc"
2. The different int data types have different size in memory and as such have a different range of numbers that it can store.
3. int8=2^7 digit range:-128 to 127(0 included), int16=2^17 digit range: -32768 to 32767 , int32=2^31 range: -2147483648 to 2147483647, int63 range: -9223372036854775808 to 9223372036854775807

Unsigned Integer- same as integer but it can only store positive integer -> "uint"
1. All the bits are used to store the value as such unit has greater range.
2. uint8= 2^8 range-0 to 255, uint16=2^16 range-0 to 65535, uint 32=2^32 range-0 to 429467295, uint64=2^64 range-0 to 18446744073709551615

Other Ralated Data Types

Type                    Description
byte                    alias for uint8
rune                    alias for int32
uint                    size would be either 32 or 64 bits, it depent on the system architectur
int                     same goes for int also