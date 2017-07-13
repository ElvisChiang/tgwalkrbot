
## Planet

- path in apk
```
assets/Images/thumbs/preview/
```
- combine new planet pictures
```bash
find . -name '*.png' | xargs -n1 -I XXX gm composite 'XXX' ../bgPlanet.png 'output/XXX'
```

## Satellite

- path in apk
```
assets/Images/thumbs/satellites/
```
- combine new satellite pictures

```bash
find . -name '*.png' | xargs -n1 -I XXX gm composite 'XXX' ../bgSatellite.png 'output/XXX'
```
