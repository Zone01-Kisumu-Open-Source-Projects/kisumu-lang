# Kisumu Programming Language Specification

## 8. Concurrency

Concurrency in Kisumu is inspired by Go's lightweight goroutines and channels, enabling developers to write highly concurrent and scalable applications. Kisumu's concurrency model simplifies the creation, management, and synchronization of concurrent tasks.

### 8.1 Goroutines

Goroutines are lightweight threads managed by the Kisumu runtime. They enable functions to run concurrently without the overhead of traditional threads.

#### Syntax

```ksm
go functionName(arguments);
```

#### Example

```ksm
func sayHello() {
    print("Hello from a goroutine!");
}

go sayHello();
print("Hello from the main function!");
```

**Output:**  
The order of the outputs may vary due to concurrency.

### 8.2 Channels

Channels are used for communication between goroutines. They provide a safe way to send and receive data across concurrent tasks.

#### 8.2.1 Creating Channels

Channels are created using the `make` keyword and have a specific type.

```ksm
channel<int> ch = make(channel<int>);
```

#### 8.2.2 Sending and Receiving

- **Send**: Use the `<-` operator to send data into the channel.
- **Receive**: Use the `<-` operator to receive data from the channel.

```ksm
channel<int> ch = make(channel<int>);

go func() {
    ch <- 42; // Sending data
}();

int value = <-ch; // Receiving data
print(value); // Output: 42
```

### 8.3 Buffered Channels

Buffered channels allow for a fixed number of elements to be stored without requiring an immediate receiver.

#### Syntax

```ksm
channel<int> ch = make(channel<int>, 3); // Buffered channel with a size of 3
```

#### Example

```ksm
channel<int> ch = make(channel<int>, 2);

ch <- 1;
ch <- 2;
// ch <- 3; // Blocks because the buffer is full

print(<-ch); // Output: 1
print(<-ch); // Output: 2
```

### 8.4 Channel Operations

#### Closing Channels

Channels can be closed using the `close` function to indicate no further data will be sent.

```ksm
channel<int> ch = make(channel<int>);

go func() {
    ch <- 10;
    close(ch);
}();

int value = <-ch;
print(value); // Output: 10
```

#### Select Statement

The `select` statement allows waiting on multiple channel operations simultaneously.

```ksm
channel<string> ch1 = make(channel<string>);
channel<string> ch2 = make(channel<string>);

go func() {
    ch1 <- "Hello from ch1!";
}();

go func() {
    ch2 <- "Hello from ch2!";
}();

select {
    case msg1 = <-ch1:
        print(msg1);
    case msg2 = <-ch2:
        print(msg2);
    default:
        print("No messages received.");
}
```

### 8.5 Synchronization Primitives

Kisumu provides additional tools for synchronization:

#### 8.5.1 Wait Groups

Wait groups allow waiting for a set of goroutines to complete.

```ksm
var wg = make(waitgroup);

wg.add(1);
go func() {
    print("Task completed.");
    wg.done();
}();

wg.wait();
print("All tasks are done.");
```

#### 8.5.2 Mutexes

Mutexes provide mutual exclusion for critical sections.

```ksm
var mtx = make(mutex);
int counter = 0;

go func() {
    mtx.lock();
    counter++;
    mtx.unlock();
}();
```

### 8.6 Error Handling in Concurrent Code

Kisumu channels can also be used to propagate errors from goroutines back to the main function. Structured error handling for concurrency will be introduced in future updates.
