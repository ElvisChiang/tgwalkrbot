# filename format

default-<PLANET>-placeholder@2x.png

example: default-witch-placeholder@2x.png

convert image to have background, script

install GraphicsMagick(or ImageMagick, different command)

$ sips -Z 225 bgPlanet.png
$ find . -name '*.png' | xargs -n1 -I XXX gm composite 'XXX' ../bgPlanet.png 'output/XXX'

