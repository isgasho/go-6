/*
Package xcolor implements Red, RedGreen, and RGB color models matching the core
image/color interface.

Note that there are good reasons these color types aren't in the core
image.color package. The native color types may have optimized fast-paths
for many use cases.

This package is a tradeoff of these optimizations against lower memory
usage. This package is intended to be used in computer graphics (e.g.
OpenGL) where images are uploaded to the GPU in a specific format (such as
GL_R, GL_RG, or GL_RGB) and we don't care about the performance of native
Go image manipulation.

OpenGL® and the oval logo are trademarks or registered trademarks of Hewlett Packard Enterprise in
the United States and/or other countries worldwide.

See also: ximage (https://tawesoft.co.uk/go/ximage)

For license information, documentation, source code, support, links, etc. please see
https://tawesoft.co.uk/go/ximage/xcolor

This module is part of https://tawesoft.co.uk/go
*/
package xcolor // import "tawesoft.co.uk/go/ximage/xcolor"

// SPDX-License-Identifier: BSD-3-Clause

// Code generated by (tawesoft.co.uk/go/) fluff.py: DO NOT EDIT.