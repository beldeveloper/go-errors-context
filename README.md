# go-errors-context

## Purpose

This library is intended for adding context to the errors in Go.

## Examples

### New error

```
errors.NewWithContext("text", errors.Context{
    Path:   "component.method",
    Params: errors.Params{"entity_id": id},
})
```

### Wrap the error

```
errors.WrapContext(err, errors.Context{
    Path:   "component.method",
    Params: errors.Params{"entity_id": id},
})
```

### To string

```
error text; entity_id=1; component.foo -> component.bar
```
