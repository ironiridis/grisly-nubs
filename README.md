# grisly-nubs

## Important

`grisly-nubs` requires a 24-bit framebuffer. Edit `config.txt` to include this line:

    framebuffer_depth=24

## General

This is a bare-to-the-metal application designed to render pre-loaded image files to the HDMI port. It does not do a lot of sanity checks. For now, `grisly-nubs` is a no-training-wheels, no-seat-belts app.

For best results:

1. Edit `config.txt` to configure your HDMI output framebuffer for a static resolution. Don't rely on EDID.
2. Make sure all of your content images are exactly the dimensions of the framebuffer.
3. Nearly all modern displays do not require overscan. Disable it in `config.txt` and set your display to be "dot-by-dot" or whatever.

## Coming soon

`grisly-nubs` will soon...

1. Provide an HTTP interface for managing assets on the device
2. Provide a very basic TCP interface for selecting assets for display

## Coming ... uh, later

`grisly-nubs` will eventually...

1. Provide configurable sections on-screen
2. Provide control via an interactive line-based TCP session
3. Provide control via HTTP endpoints and feedback via JSON
