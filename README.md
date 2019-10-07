# prisma-golang-tutorial
prisma tutorial - downgraded golang to 1.11.x had to use dep for this, plus the generated code interfaces did not match the tutorial - i.e the return values were slices of pointers, where the tutorial implementation did not have this; I had to use slice pointers on at least one api - due to a conversion error.  Also I had to downgrade a package from v2 otherwise there were start up errors

It was not the easiest tutorial to get through, enough typos to cause headaches

```
[[override]]
  name = "github.com/russross/blackfriday"
  version = "1.5.2"
```

I started on 1.13 and using modules, but had issues, that I thought were related to the ecosystem, but it was just the tutorial.  Will try this with 1.13 and modules next and check it in as a different branch
