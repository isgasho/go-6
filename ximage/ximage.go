/*
Package ximage implements Red, RedGreen, and RGB images matching the core
image interface.

Note that there are good reasons these image types aren't in the core image
package. The native image types may have optimized fast-paths for many use
cases.

This package is a tradeoff of these optimizations against lower memory
usage. This package is intended to be used in computer graphics (e.g.
OpenGL) where images are uploaded to the GPU in a specific format (such as
GL_R, GL_RG, or GL_RGB) and we don't care too much about the performance of
native Go image manipulation.

Home page https://tawesoft.co.uk/go

For source code see https://github.com/tawesoft/go/tree/master/ximage

For documentation see https://godoc.org/tawesoft.co.uk/go/ximage

See also https://godoc.org/tawesoft.co.uk/go/ximage/xcolor
*/
package ximage // import "tawesoft.co.uk/go/ximage"
