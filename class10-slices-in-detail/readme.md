## class10

## 1. Introduction

Slices in Go are a convenient and flexible way to work with sequences of elements. They are built on top of arrays but offer dynamic sizing.

---

## 2. Slice Basics

- A slice is a descriptor of an array segment.
- It consists of:
  - A pointer to the underlying array
  - A length
  - A capacity
- Slices are more flexible than arrays since they can grow and shrink dynamically.

---

## 3. Slice Descriptor

Internally, a slice is represented by 3 fields:

1. **Pointer** to the underlying array
2. **Length** – number of elements in the slice
3. **Capacity** – maximum number of elements the slice can grow to before reallocating

> Multiple slices can share the same underlying array.  
> Mutating one slice may affect the others if they overlap.

---

## 4. Creating Slices

### From an array

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // slice contains [2, 3, 4]
```
### Using make
```
slice := make([]int, 3, 5) // length 3, capacity 5
```
## 5. Slice Operations
Accessing elements
```
value := slice[0] // first element
slice[1] = 10     // modify second element
```
#### Appending elements
```
slice = append(slice, 6, 7) // add elements to the slice
```
#### Copying slices
```
newSlice := make([]int, len(slice))
copy(newSlice, slice)
```

## 6. Length and Capacity
```
len(slice) // returns length
cap(slice) // returns capacity
```

- Length: Number of elements in the slice
- Capacity: Maximum elements the slice can hold before growing

## 7. Sharing Underlying Array
```
arr := [5]int{1, 2, 3, 4, 5}
a := arr[1:4] // [2, 3, 4]
b := arr[2:5] // [3, 4, 5]

a[1] = 99
fmt.Println(b) // [3, 99, 5]
```

- Changing a affects b because they share part of the same array.

## 8. Summary
- Slices are lightweight descriptors over arrays.
- They provide flexible, dynamic-length sequences.
- Multiple slices can share the same underlying array.
- Use len() and cap() to inspect slices.
- Use append() to grow slices dynamically.