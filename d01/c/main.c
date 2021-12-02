#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <limits.h>

int * readInput(size_t *length) {
    size_t len = 4;
    int *buf = malloc(len * sizeof *buf);

    size_t i = 0;
    int *temp; // to save buf in case realloc fails

    while(scanf("%d", buf+i) == 1) { 
        i++;
        if(i == len) {
            temp = buf;
            len *= 2;
            buf = realloc(buf, len * sizeof *buf);
            if(buf == NULL) {
                printf("OOM\n");
                buf = temp;
                break;
            }
        }
    }

    *length = (size_t)i;
    return buf;
}

int solveP1(int *buf, size_t size) {
	int lastDepth = INT_MAX;
	int increases = 0;

    for(size_t j = 0; j < size; j++) {
		if(buf[j] > lastDepth){
			increases++;
		}
		lastDepth = buf[j];
    }
	return increases;
}

int solveP2(int *buf, size_t size) {
	int lastSum = INT_MAX;
	int increases = 0;

    for(size_t j = 0; j < size-2; j++) {
		int sum = buf[j] + buf[j+1] + buf[j+2];
		if(sum > lastSum){
			increases++;
		}
		lastSum = sum;
    }
	return increases;
}

int main(void) {
    size_t bufLen = 0;
    int *buf = readInput(&bufLen);

    printf("p1: %d\n", solveP1(buf, bufLen));
    printf("p2: %d\n", solveP2(buf, bufLen));

    free(buf);
    buf = NULL;
    return 0;
}
