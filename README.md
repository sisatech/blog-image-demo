# blog-image-demo

### Description

### Build Process
```sh
CGO_ENABLED=0 go build
```

### Modying Behaviour
The application looks for the following environment variables to influence its behaviour:

DEMO_SOURCE_DIR: should be the path to the directory where unprocessed pictures are stored.
DEMO_PROCESSED_DIR: should be the path to the directory where processed pictures should be stored.
DEMO_POLYGONS: the number of polygons used to paint the picture. This number dramatically effects processing time.
DEMO_RESOLUTION: the number of pixels to scale the image down to during processing. This number dramatically effects processing time.
DEMO_PIXELS: the number of pixels on each side of the finished processed picture.
