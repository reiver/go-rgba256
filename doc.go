/*
Package rgba32 provides a type that represents an RGBA color with 8-bits per color channel,
for a total of 32-bits altogether,
that seamlessly works with Go's built-in "image", "image/color", and "image/draw" packages.

An rgba32.Slice is primarily used as a Go built-in `color.Color`.

But it can also be used as an `image.Image`. When it is used as an `image.Image` it is an
“infinite” image with the same color everywhere.
*/
package rgba32
