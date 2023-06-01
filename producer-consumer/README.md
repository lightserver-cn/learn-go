# Redis Message Queue with Consumer Groups - Implementation README

This code provides a basic implementation of a Redis Message Queue with Consumer Groups using the `go-redis` library in Golang. The implementation demonstrates how to create a consumer group, produce messages to the queue, and consume messages from the queue using a consumer.

## Prerequisites

Before running the code, ensure that you have the following prerequisites:

- Golang installed on your system.
- Redis server running on `localhost` with the default port `6379`.
- The `go-redis` package installed. You can install it using the following command:

  ```
  go get github.com/redis/go-redis/v9
  ```

## Code Overview

The code can be divided into the following sections:

1. Setting up the Redis client: In the `main` function, a Redis client is created using the `redis.NewClient` function. The client is configured with the Redis server's address, password (if required), and database number.

2. Checking if the consumer group exists: The code checks if the specified consumer group exists using the `XInfoGroups` command. If the group exists, the `groupExists` variable is set to `true`.

3. Creating the consumer group: If the consumer group does not exist, it is created using the `XGroupCreateMkStream` command. This command creates a new consumer group and a stream if they don't already exist.

4. Consuming messages: The code starts a goroutine that acts as a consumer. The consumer continuously reads messages from the queue using the `XReadGroup` command. It processes each message and acknowledges its consumption using the `XAck` command.

5. Producing messages: The code uses the `XAdd` command to add messages to the queue. It demonstrates how to send multiple messages to the queue at different times.

6. Handling message processing: The consumer goroutine receives messages and processes them within a loop. It prints the received message and then acknowledges its consumption. You can customize the processing logic as per your requirements.

7. Closing the Redis connection: After the consumer goroutine completes, the Redis client is closed using the `Close` method.

## Running the Code

To run the code, follow these steps:

1. Save the code to a file with a `.go` extension (e.g., `main.go`).

2. Open a terminal and navigate to the directory containing the code file.

3. Build and run the code using the `go run` command:

   ```
   go run main.go
   ```

4. The code will produce and consume messages from the Redis Message Queue. The output will be displayed in the terminal as the messages are processed.

5. Wait for the code to finish processing the messages. Afterward, the Redis connection will be closed, and the program will exit.

Note: Make sure the Redis server is running before executing the code.

## Customization

You can customize the code according to your specific requirements:

- Adjust the Redis server connection settings in the `redis.NewClient` function call.
- Modify the queue name, consumer group, and other relevant variables as needed.
- Modify the message processing logic inside the consumer goroutine.
- Change the sleep durations to control the timing of message production and consumption.

Feel free to experiment with different configurations and logic to suit your application's needs.

## Conclusion

This README provides an overview of the Redis Message Queue with Consumer Groups implementation using Golang and the `go-redis` library. By following the instructions and customizing the code, you can create a simple message queue system that enables message production and consumption with multiple consumers.
