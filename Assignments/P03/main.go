package main

import (
	"github.com/jstevens8185/img_colors"
	"github.com/jstevens8185/img_get"
	"github.com/jstevens8185/img_gray_scale"
	"github.com/jstevens8185/img_text"
)

func main() {
	img_get.GetImage("https://scontent.cdninstagram.com/v/t51.2885-15/47474109_128024318196493_7634142724408351138_n.jpg?stp=dst-jpg_e35&efg=eyJ2ZW5jb2RlX3RhZyI6ImltYWdlX3VybGdlbi43NTB4NzUwLnNkciJ9&_nc_ht=scontent.cdninstagram.com&_nc_cat=101&_nc_ohc=g83IpmM2LfsAX90z638&edm=APs17CUBAAAA&ccb=7-5&ig_cache_key=MTkyNzQyNTY4MTU3MjI0NTQ0OA%3D%3D.2-ccb7-5&oh=00_AfBVyBi46JfzqDEVJHmSEf37SA4sLrFKqWKid9hlpD8LSQ&oe=6528777A&_nc_sid=10d13b", "", "colors.jpeg")
	img_gray_scale.ConvertImageToGrayscale("colors.jpeg", "colors_gray_scale.jpeg")
	img_colors.PrintImagePixelsToFile("colors.jpeg", "colors_pixel_counts.csv")
	img_text.AddTextOverlay("colors.jpeg", "colors_labeled.jpg", "COLORS")
}
