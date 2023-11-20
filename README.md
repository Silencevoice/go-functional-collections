This is only a coding exercise.
The purpose of this exercise is to mimic the behaviour if the *contains*, *reverse*, *sort*, *map*, *filter* and *reduce* functions that in other languages are provided to manipulate arrays (slices).

Also, I have toyed with the possibility of adding those methods to slices, but there were some limits that go imposes.

- slices cannot be receivers of new methods, unless an alias is used.
- type aliases cannot be receivers of generic methods. They have to specify the type explicitly.

This implementation is by no means performant or efficient.