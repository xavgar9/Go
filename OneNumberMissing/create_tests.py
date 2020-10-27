from sys import stdin
import numpy as np

def main():
	for _ in range(10):
		case = stdin.readline().split()
		print("{[]int{",end="")
		for n in case:
			print(str(n)+",",end="")
		print("},",-1,"},")
main()