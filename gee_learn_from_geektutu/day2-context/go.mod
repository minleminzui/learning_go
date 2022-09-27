module example

go 1.18

require gee v0.0.0

// 在go 1.11版本开始后，引用相对路径的package需要使用上述方式
replace gee => ./gee