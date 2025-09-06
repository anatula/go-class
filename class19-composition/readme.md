## Class 19 Composition

![](./img/composition.png)

- embeded in the SAME level, not `pl.x.Path`
- I embed a type within a type even if not s struct
- promotion: fields and methods (for example String method of Pair gets promoted to PairWithLength)
- A PairWithLength is not a subclass/subtype of Pair
- `pl.Pair*()`
- inheritance pick up field from parent class, subclass of parent class, in go we don't have inheritance: we're pulling in the fields of Pair, embedded in PairWithLegth. We are promoting into PairWithLength but they are different types

![](./img/composition2.png)
- both Pair and PairWithLength can be examples of Filenamer, because method was promoted into PairWithLength

### Composition with pointer types

![](./img/composition-pointers.png)

- embeds a pointer to PairWithLenth
- when printing it's calling the String method of PairWithLength because I embeded it (even if it's a pointer)

### Sortable interface

![](./img/sortable-interface.png)

![](./img/sortable-example.png)

