# gosplit
gosplit is tool for split image.

# install
````
go install github.com/desktopgame/gosplit/cmd/gosplit
````

# format
````
gosplit -r $rowCount -c $columnCount -i $input_file
````

# example
````
gosplit -r 3 -c 13 -i /Users/koya/Documents/Gimp/AlphaMapFill_14x28.png
````

# screenshot
input image  
![image 1](image/AlphaMapFill_14x28.png)

command
````
gosplit -r 3 -c 13 -i /Users/koya/go/src/github.com/desktopgame/gosplit/image/AlphaMapFill_14x28.png
````
output image  
![image 2](image/gosplit/AlphaMapFill_14x28_0_0.png)  
![image 3](image/gosplit/AlphaMapFill_14x28_0_1.png)  
![image 4](image/gosplit/AlphaMapFill_14x28_0_2.png)  
![image 5](image/gosplit/AlphaMapFill_14x28_0_3.png)  