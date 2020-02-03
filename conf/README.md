# Configuration

Place all configuration files, videos and images in this folder.

A configuration file consists of a JSON array that contains a number of groups, which contain any number of *Element*s.

Each group and each *Element* has a `Color` attribute. Supported colors (by TentCSS) are:
* `primary`: blue
* `secondary`: green
* `tertiary`: red

Video and image files that are referenced in a configuration file must be present in this directory or its subdirectories. Use forward slashes to reference files in subdirectories: `subdir/example.mp4`.

```json
[
	{
		"Color": "primary",
		"Name": "Group 1",
		"Elements": [
			{
				"Name": "Some scrolling text",
				"Color": "primary",
				"Endpoint": "marquee",
				"Parameter": "Hello World"
			},
			{
				"Name": "A video",
				"Color": "primary",
				"Endpoint": "video",
				"Parameter": "example.mp4"
			}
		]
	},
	{
		"Color": "secondary",
		"Name": "Group 2",
		"Elements": [
			{
				"Name": "Another scrolling text",
				"Color": "secondary",
				"Endpoint": "marquee",
				"Parameter": "Test Test"
			},
			{
				"Name": "An image",
				"Color": "secondary",
				"Endpoint": "image",
				"Parameter": "example.png"
			},
			{
				"Name": "Blank screen",
				"Color": "tertiary",
				"Endpoint": "black",
				"Parameter": ""
			}
		]
	}
]
```

Result:

[![Screenshot: Controller View](https://i.imgur.com/g2YZya7l.png)](https://i.imgur.com/g2YZya7.png)
