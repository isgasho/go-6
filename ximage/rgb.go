// Based on https://golang.org/src/image/image.go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ximage

import (
    "image"
    "image/color"
    "tawesoft.co.uk/go/ximage/xcolor"
)

// RGB is an in-memory image whose At method returns color.RGB values.
type RGB struct {
	// Pix holds the image's pixels, as Red values. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *RGB) ColorModel() color.Model { return xcolor.RGBModel }

func (p *RGB) Bounds() image.Rectangle { return p.Rect }

func (p *RGB) At(x, y int) color.Color {
	return p.RGBAt(x, y)
}

func (p *RGB) RGBAt(x, y int) xcolor.RGB {
	if !(image.Point{x, y}.In(p.Rect)) {
		return xcolor.RGB{}
	}
	i := p.PixOffset(x, y)
	return xcolor.RGB{R: p.Pix[i], G: p.Pix[i+1], B: p.Pix[i+2]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*1
}

func (p *RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
    i := p.PixOffset(x, y)
    rgba := xcolor.RGBModel.Convert(c).(color.RGBA)
    p.Pix[i]   = rgba.R
    p.Pix[i+1] = rgba.G
    p.Pix[i+2] = rgba.B
}

func (p *RGB) SetRGB(x, y int, c xcolor.RGB) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
    p.Pix[i]   = c.R
    p.Pix[i+1] = c.G
    p.Pix[i+2] = c.B
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB) Opaque() bool {
	return true
}

// NewRGB returns a new RGB image with the given bounds.
func NewRGB(r image.Rectangle) *RGB {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 3 * w * h)
	return &RGB{pix, 3 * w, r}
}
