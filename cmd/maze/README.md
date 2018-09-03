# maze

**maze** generates random mazes and finds paths through them.

## Usage

	  -o string
			output file (default "stdout")
	  -p	show path through maze
	  -s int
			dimension of each maze cell in pixels (default 30)
	  -x int
			x dimension of maze (default 25)
	  -y int
			y dimension of maze (default 25)

## Examples

    maze -x 60 -y 40 -s 10 -p -o maze1.png

![maze1](https://user-images.githubusercontent.com/5059971/44964372-1b6f3800-aefe-11e8-925a-6ff9d560a575.png)

    maze -x 30 -y 50 -s 15 -p -o maze2.png

![maze2](https://user-images.githubusercontent.com/5059971/44964442-93d5f900-aefe-11e8-8293-30f14f225956.png)

    maze -x 10 -y 10 -s 15 -p -o maze2.png

![maze3](https://user-images.githubusercontent.com/5059971/44964453-a819f600-aefe-11e8-88c7-6f1058aa22bc.png)

