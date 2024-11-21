# 数组和切片
## 数组
- 数组定义--数组定长
  - `var` 定义固定大小的数组，在定义时需要指定大小，分配一片连续的内存
  - `[...]` 自动推断数组大小，核心在于数组定义的时候已经确定了数组大小，分配固定的连续内存空间
```go
package main

import (
  "fmt"
  "unsafe"
)

func main() {
  // 指定数组大小
  var a1 [5]int // int == int64
  // 自动推断数组大小
  a2 := [...]int{1, 2, 3}
  fmt.Printf("a1 = %v , a2 = %v,a1Len = %d, a2Len = %d, a1Size = %d, a2Size = %d \n", a1, a2, len(a1), len(a2), unsafe.Sizeof(a1), unsafe.Sizeof(a2))
  // a1 = [0 0 0 0 0] , a2 = [1 2 3],a1Len = 5, a2Len = 3, a1Size = 40, a2Size = 24
  // 按索引赋值
  a3 := [...]int{2: 2, 4: 4} //... 表示自动核算大小，分配固定内存
  // 按照索引复制[0,0,2,0,4] // len = 5, Sizeof(a3) = 40
  fmt.Println(a3)
  // 按索引赋值
  a4 := [5]int{2: 2, 4: 4}
  fmt.Println(a4) //[0 0 2 0 4] // len = 5, Sizeof(a4) = 40
}
```
## 切片
- 切片定义-- 数组不定长 
  - 通过 `make` 分配内存，返回的是引用类型本身
    - make是生成一个可变大小的内存块，并返回一个它的引用
  - 通过 `new` 分配内存，返回的是指向类型的指针，并且内存置为0
    - `new` 可以申请任何类型变量内存块并返回一个指针指向它
  - `var` 直接定义不定长的数组
```go
func main() {
	// -------------------- 切片 -----------------
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3:4]), cap(sli[0:3:4]), sli[0:3:4]) //len=3 cap=4 slice=[1 2 3]
	// 定义切片
	var b1 []int //切片不定长
	b1 = append(b1, 1)
	fmt.Println(b1) //[1]

	list := new([]int)
	*list = append(*list, 2)
	fmt.Println(*list) //[2]

	var b3 = []int{1, 2, 3} //切片不定长
	b3 = append(b3, 4)
	fmt.Println(b3) //[1,2,3,4]

	// make初始化
	b2 := make([]int, 5, 5+3)                                    // make([],len,cap)
	fmt.Printf("b2 = %v ,len=%d,cap=%d\n", b2, len(b2), cap(b2)) //b2 = [0 0 0 0 0] ,len=5,cap=8
}
```
### 切片扩容
- 通过 `make` 定义的切片结构由三部分组成：`make(type,len,cap)`
  - `type` 表明数据类型
  - `len` 用来初始化内存数据
  - `cap` 表示当前切片的容量，用来分配初始化内存大小
- 当数据超出 `cap` 容量的时候，就会重新分配内存扩容：
  - 当原切片长度小于 `1024` 时，新的切片长度直接加上 `append` 元素的个数，容量则会直接 `*2`
  - 当原切片长度大于等于 `1024` 时，新的切片长度直接加上 `append` 元素的个数，容量则会增加 `1/4`
```go
func growslice(et *_type, old slice, cap int) slice {
//cap 输入的新cap值
    newcap := old.cap 
    doublecap := newcap + newcap
    // 如果新容量大于旧容量的两倍，则直接按照新容量大小申请
    if cap > doublecap {
			newcap = cap
    } else {
        // 如果原有长度小于1024，则新容量是旧容量的2倍
        if old.len < 1024 {
            newcap = doublecap
        } else {
            // 按照原有容量的 1/4 增加，直到满足新容量的需要
            for 0 < newcap && newcap < cap {
                newcap += newcap / 4
            }
            if newcap <= 0 {
                newcap = cap
            }
        }
    }
}
```
### 切片数据赋值
- 切片通过 `append` 在末尾增加数据
  - 切片之间通过 `append` 值传递，两部分数据互不影响
  - `append` 时没必要初始化新切片的 `len`， `append` 时会自动将数据依次加到新切片末尾
```go
func appendvalue() {
	// -------------------- 切片 直接引用 复制 -----------------
	a := make([]int, 0, 6) //默认0值为空 容量3
	a = append(append(append(a, 1), 2), 3)
	fmt.Printf("alen=%d acap=%d aslice=%v \n", len(a), cap(a), a) // alen=3 acap=6 aslice=[1 2 3]

	b := make([]int, 0) //声明一个长度为a 切片指针变量b
	b = append(b, a...)
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[1 2 3] blen=3 bcap=3 bslice=[1 2 3]
	// 值传递，两部分数据互不影响
	a[0] = 99
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 2 3] blen=3 bcap=3 bslice=[1 2 3]
	b[1] = 100
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 2 3] blen=3 bcap=3 bslice=[1 100 3]
}
```
- `copy` 复制切片数组
  - `copy` 值传递，实现单独的内存分配，复制出的切片和原切片处于独立的内存区，两部分数据互不影响
  - `copy(dst, src)` 从源数据拷贝 `min(len(dst), len(src))`个元素
    - 因此，拷贝的 `dst` 数组需要初始化长度
```go
func copyvalue() {
	// -------------------- 切片 直接引用 复制 -----------------
	a := make([]int, 0, 6) //默认0值为空 容量3
	a = append(append(append(a, 1), 2), 3)
	fmt.Printf("alen=%d acap=%d aslice=%v \n", len(a), cap(a), a) // alen=3 acap=6 aslice=[1 2 3]

	b := make([]int, len(a)) //声明一个长度为a 切片指针变量b
	copy(b[:], a[:])
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[1 2 3] blen=3 bcap=3 bslice=[1 2 3]
	// 值传递，两部分数据互不影响
	a[0] = 99
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 2 3] blen=3 bcap=3 bslice=[1 2 3]
	b[1] = 100
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 2 3] blen=3 bcap=3 bslice=[1 100 3]
}
```
- 直接复制出的新切片是引用传递
  - 只要新切片的 `cap` 不超过原切片，则新切片和旧切片指向同一个引用类型，修改任意切片数据都会影响相同的引用类型
  - 但是，新切片的 `cap` 超过原切片后，就会重新申请内存容量，此时两者指向不同的内存引用，对于数组数据的操作就互不影响
```go
func reference() {
	// -------------------- 切片 直接引用 复制 -----------------
	a := make([]int, 0, 6) //默认0值为空 容量3
	a = append(append(append(a, 1), 2), 3)
	fmt.Printf("alen=%d acap=%d aslice=%v \n", len(a), cap(a), a) // alen=3 acap=6 aslice=[1 2 3]

	b := make([]int, 0) //声明一个长度为0 切片指针变量b
	b = a
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[1 2 3] blen=3 bcap=3 bslice=[1 2 3]
	// 引用复制，在未超过旧切片的 cap 时，两部分指向相同的地址
	a[0] = 99
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 2 3] blen=3 bcap=6 bslice=[99 2 3]
	b[1] = 100
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 100 3] blen=3 bcap=6 bslice=[99 100 3]

	b = append(append(append(b, 4), 5), 6)
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 100 3] blen=6 bcap=6 bslice=[99 100 3 4 5 6]
	// 再次新增数据，超出旧切片的容量
	b = append(b, 7)
	b[1] = 98
	fmt.Printf("alen=%d acap=%d aslice=%v blen=%d bcap=%d bslice=%v \n", len(a), cap(a), a, len(b), cap(b), b) //alen=3 acap=6 aslice=[99 100 3] blen=7 bcap=12 bslice=[99 98 3 4 5 6 7]
}
```
