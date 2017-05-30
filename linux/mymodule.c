/*
 *  hello-1.c - The simplest kernel module.
 */
#include <linux/slab.h>
#include <linux/module.h>	/* Needed by all modules */
#include <linux/kernel.h>	/* Needed for KERN_INFO */
#include <linux/fs.h>
#include <linux/types.h> /* size_t */
#include <linux/uaccess.h>
#include <linux/string.h>
#include <linux/ioctl.h>

#define MAJOR_NUM 60
#define DRIVER_AUTHOR "SU JIN CHO"
#define DRIVER_DESC   "A simple driver"

/*iocil command*/
#define IOC_MAGIC MAJOR_NUM
#define IOCTL_READ_SIZE _IO(IOC_MAGIC,0)
#define IOCTL_INIT_BUFFER _IO(IOC_MAGIC,1)

/*function declaration */
int my_open(struct inode *inode, struct file *filp);
int my_close(struct inode *inode, struct file *filp);
ssize_t my_read(struct file *filp, char *buf, size_t count, loff_t *f_pos);
ssize_t my_write(struct file *filp, const char *buf, size_t count, loff_t *f_pos);
long my_ioctl(struct file *f, unsigned int cmd, unsigned long arg);
void my_exit(void);
int my_init(void);

/*access functions */
static struct file_operations fops = {
    .read = my_read,
    .write = my_write,
    .open = my_open,
    .release = my_close,
    .unlocked_ioctl = my_ioctl
};

/* Declaration of the init and exit functions */
module_init(my_init);
module_exit(my_exit);

/*module information */
MODULE_AUTHOR(DRIVER_AUTHOR);
MODULE_DESCRIPTION(DRIVER_DESC);
MODULE_LICENSE("GPL");

int major = 60; /*Major number assign to my device driver */
char* buffer;

int my_init(void)
{
    int result;

     //register module
    printk(KERN_INFO "mymodule: register start\n");
    result = register_chrdev(MAJOR_NUM, "mymodule", &fops);
    if(result < 0){
        printk(KERN_ERR "ERROR: Cannot register the module!\n");
        return result;
    }

	//allocate buffer
	buffer = (char*)kmalloc(sizeof(char)*1024, GFP_KERNEL);
	if(!buffer){
	  printk(KERN_ERR "mymodule: (ERROR) Cannot allocate the buffer!\n");
	  return -1;
	}

    printk(KERN_INFO "mymodule: (SUCCESS) register end\n");
	return 0;
}

int my_open(struct inode *inode, struct file *filp){

    printk(KERN_INFO "mymodule: open start\n");
    //initialize the buffer to a known text
    buffer = "Welcome to UCSC Second Assignment";
    printk(KERN_INFO "mymodule: (SUCCESS) open end\n");
    return 0;
}

int my_close(struct inode *inode, struct file *filp){

    printk(KERN_INFO "mymodule: close start\n");
    buffer = "";
    printk(KERN_INFO "mymodule: (SUCCESS) close end\n");
    return 0;
}

ssize_t my_read(struct file *filp, char *userBuf, size_t count, loff_t *f_pos){

    long result;
    printk(KERN_INFO "mymodule: read start\n");

    //copy_to_user( to, from, n );
    result = copy_to_user(userBuf, buffer, 1024);
    if(result != 0){
        printk(KERN_ERR "mymodule: (ERROR) cannot from the kernel space buffer\n");
    }
    printk(KERN_INFO "mymodule: read end(SUCCESS)\n");
    return 0;
}

ssize_t my_write( struct file *filp, const char *userBuf, size_t count, loff_t *f_pos) {

    int err;
    printk(KERN_INFO "mymodule: write start\n");

    //copy_from_user( to, from, n );
    buffer = kmalloc(count, GFP_KERNEL);
    err = copy_from_user(buffer, userBuf, count);
    if(err < 0){
        printk(KERN_ERR "mymodule: (ERROR) cannot write to the kernel space buffer\n");
        kfree(buffer);
        return err;
    }

    printk(KERN_INFO "mymodule: write end (SUCCESS)\n");
    return 0;
}

long my_ioctl(struct file *f, unsigned int ioctl_cmd, unsigned long arg){

    int result;
    printk(KERN_INFO "mymodule: ioctl start\n");
    switch(ioctl_cmd){
        case IOCTL_READ_SIZE:
            printk(KERN_INFO "mymodule: ioctl, read the size of string\n");
            result = strlen(buffer);
        break;
        case IOCTL_INIT_BUFFER:
            printk(KERN_INFO "mymodule: ioctl, init buffer\n");
            buffer = kmalloc(1024, GFP_KERNEL);
            buffer = "Welcome to UCSC Second Assignment";
            result =  strlen(buffer);
        break;
        default:
            printk(KERN_ERR "mymodule: (ERROR) Wrong ioctl command\n");
        return -1;
    }

    printk(KERN_INFO "mymodule: ioctl end (SUCCESS)\n");
    return result;
}

void my_exit(void)
{
	/* Freeing the major number */
	printk(KERN_INFO "mymodule: unregister start\n");
    unregister_chrdev(MAJOR_NUM, "mymodule");

     /* Freeing buffer memory */
    if (buffer) {
        kfree(buffer);
    }

    printk(KERN_INFO "mymodule: unregister (SUCCESS) end\n");
}
