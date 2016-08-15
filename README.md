# go-promise

Go promise implementation heavily influenced by JavaScript
[Promise A+ spec](https://promisesaplus.com/).


## General specification

### 1. Terminology

- **Promise** is a structure with the `Wait()` method whose behavior
conforms to this specification.
- **Waitable** is a structure that defines a `Wait()` method.
- **Value** is any legal Go value (generally, an `interface{}`).
- **Error** is a Go error.
- **Reason** is an error that indicates why a promise was rejected.

### 2. Requirements

#### 2.1. Promise States

A promise must be in one of three states: pending, fulfilled,
or rejected.

Fulfilled and rejected promises are generally called settled.

- When pending, a promise:
    - may transition to either the fulfilled or rejected state.
- When fulfilled, a promise:
    - must not transition to any other state.
    - must have a value, which must not change.
- When rejected, a promise:
    - must not transition to any other state.
    - must have a reason, which must not change.

Here, “must not change” means immutable identity, but does not imply
deep immutability.

#### 2.2. The `Wait()` method

`Wait()` is a blocking call defined as follows:

    func (p *Promise) Wait() (ret interface{}, err error)

When it's called on a pending promise, the goroutine blocks
until it's settled. When called on a settled promise, the method
returns immediately.
