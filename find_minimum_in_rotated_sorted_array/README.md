```
[3, 4, 5, 1, 2]
 L     m     R

          L
	  m

	     L
```

要是 array[m] >= array[L] 所有的 left subarray number 必定大於 right subarray。那我們就不用找 left subarray 了

array[3] > 5, ye currMin = 1

1 >= 1 yes L = m+1
