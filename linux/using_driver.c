#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>		/* open */
#include <unistd.h>		/* exit */
#include <errno.h>
#include <string.h>
#include <sys/ioctl.h>

#define MAJOR_NUM 60
#define IOC_MAGIC MAJOR_NUM
#define IOCTL_READ_SIZE _IO(IOC_MAGIC,0)
#define IOCTL_INIT_BUFFER _IO(IOC_MAGIC,1) // defines our ioctl call


void printError(int num){
    printf ("Error no is : %d\n", num);
    char * msg = strerror(num);
    printf("Error description is : %s\n", msg);
};

int main(){

    int file_desc;
    int result;
    char *ptr;
    ptr = (char*)malloc(sizeof(char) * 1024);

    //open device driver
    printf("/dev/mymodule: open!!\n");
    file_desc = open("/dev/mymodule", O_RDWR);
    if (file_desc < 0)
    {
        printError(errno);
    }
    printf("****************************\n");

    printf("/dev/mymodule: ioctl!!\n");
    result = ioctl(file_desc,IOCTL_READ_SIZE);
    if(result < 0)
    {
        printError(errno);
    }
    printf("Size of default string: %d \n", result);
    printf("****************************\n");

    printf("/dev/mymodule: read!!\n");
    result = read(file_desc, ptr, 1024);
    if(result < 0)
    {
        printError(errno);
    }
    printf("Content of string: %s \n", ptr);
    printf("****************************\n");


    printf("/dev/mymodule: write!!\n");
    char* newString = "String updated";
    result = write(file_desc, newString, strlen(newString));
    if(result < 0)
    {
       printError(errno);
    }
    printf("****************************\n");


    printf("/dev/mymodule: read!!\n");
    result = read(file_desc, ptr, strlen(newString));
    if(result < 0)
    {
            printError(errno);
    }
    printf("Content of string: %s \n", ptr);
    printf("****************************\n");


    printf("/dev/mymodule: ioctl!!, read the size of string\n");
    result = ioctl(file_desc,IOCTL_READ_SIZE);
    if(result < 0)
    {
        printError(errno);
    }
    printf("Size of string: %d \n", result);
    printf("****************************\n");


    printf("/dev/mymodule: ioctl!!, init buffer\n");
    result = ioctl(file_desc, IOCTL_INIT_BUFFER);
    if(result < 0)
    {
        printError(errno);
    }
    printf("****************************\n");

    printf("/dev/mymodule: read!!\n");
    result = read(file_desc, ptr, 1024);
    if(result < 0)
    {
            printError(errno);
    }
    printf("Content of string: %s \n", ptr);
    printf("****************************\n");

    printf("/dev/mymodule: close!!\n");
    close(file_desc);
    printf("****************************\n");

    return 0;
 }

