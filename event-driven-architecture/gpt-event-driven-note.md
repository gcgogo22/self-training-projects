# How to run a Java Maven project in vs code terminal
When using **VS Code** and running the `mvn exec` command in its terminal, you need to run the command from the **root directory** of your Maven project. This is the directory where your `pom.xml` file is located, which defines your project's dependencies and build configurations.

### Steps to Run `mvn exec` in VS Code Terminal:

1. **Navigate to the Root Directory**:
   Ensure you are in the root directory of your project where the `pom.xml` file is located. You can do this by running:

   ```bash
   cd /path/to/your/project/root
   ```

   For example, if your project directory structure looks like this:
   ```
   ├── src
   │   └── main
   │       └── java
   │           └── com
   │               └── globomantics
   │                   └── Main.java
   ├── target
   ├── pom.xml
   ```

   You should run the `mvn exec` command from the directory where the `pom.xml` file is located.

2. **Run the `mvn exec` Command**:
   Once you are in the root directory (where the `pom.xml` file is), you can run the `mvn exec` command to execute your Java program.

   If you have added the `exec-maven-plugin` to your `pom.xml`, the command to run your `Main.java` class would be:

   ```bash
   mvn exec:java
   ```

   This command will look for the `mainClass` (like `com.globomantics.Main`) defined in your `pom.xml` file’s `exec-maven-plugin` configuration.

3. **Alternative Command (Specifying the Main Class Directly)**:
   If you haven’t configured the `exec-maven-plugin` in your `pom.xml` file, you can specify the main class directly when running the `mvn exec` command:

   ```bash
   mvn exec:java -Dexec.mainClass="com.globomantics.Main"
   ```

   This command explicitly tells Maven which class contains the `main` method to execute.

### Example of `exec-maven-plugin` Configuration:
If you haven’t configured the `exec-maven-plugin` yet, you can add it to your `pom.xml`:

```xml
<build>
    <plugins>
        <plugin>
            <groupId>org.codehaus.mojo</groupId>
            <artifactId>exec-maven-plugin</artifactId>
            <version>3.0.0</version>
            <configuration>
                <mainClass>com.globomantics.Main</mainClass>
            </configuration>
        </plugin>
    </plugins>
</build>
```

### Summary:
- **Navigate to the root directory of your project**, where the `pom.xml` file is located, using the terminal in VS Code.
- Run the `mvn exec:java` command to execute your Java program.
- If not configured, you can specify the main class directly using `-Dexec.mainClass`.

In VS Code, you can also open the terminal with **Ctrl + `** and ensure you're in the correct directory before running the Maven commands.

# Continue
If you are using **VS Code** and want to run a **Java program** in a Maven project without adding the `exec-maven-plugin`, the **normal approach** would be to compile the program using Maven and then run the compiled `.class` files directly using the `java` command.

### Normal Approach Without the `exec-maven-plugin`:

1. **Ensure Your Program is Compiled**:
   First, use Maven to compile your Java program. This ensures that all dependencies are downloaded and the code is compiled into the `target/classes` directory.

   In the VS Code terminal, run:

   ```bash
   mvn compile
   ```

   This will compile your Java source files (`src/main/java`) and place the compiled `.class` files in the `target/classes` directory.

2. **Run the Program Using the `java` Command**:
   Once the program is compiled, you can run it using the `java` command from the terminal. You need to specify the classpath (`target/classes`) and the fully qualified name of the main class (e.g., `com.globomantics.Main`).

   Run the following command from the **root of your project** (where the `pom.xml` is located):

   ```bash
   java -cp target/classes com.globomantics.Main
   ```

   This tells Java to:
   - Use the classpath `-cp target/classes`, which is where Maven has compiled your classes.
   - Execute the main class `com.globomantics.Main`.

3. **Handling Dependencies**:
   If your program depends on external libraries (specified in the `pom.xml`), those dependencies won’t automatically be included when you run the program using the `java` command. To include those dependencies, you can either:
   
   - **Use `mvn exec:java`** (which handles dependencies automatically, as mentioned earlier), or
   - **Manually include dependencies** in the classpath.

   If you prefer to manually handle dependencies:

   1. Run this Maven command to copy the dependencies into the `target/dependency` folder:
      ```bash
      mvn dependency:copy-dependencies
      ```

   2. Then, run the program and include both the `target/classes` directory and the `target/dependency` folder (which now contains all the necessary JAR files) in the classpath:
      ```bash
      java -cp target/classes:target/dependency/* com.globomantics.Main
      ```

   This will include all the compiled classes and dependencies.

### Summary of the Normal Approach:

1. **Compile** your Java program using:
   ```bash
   mvn compile
   ```

2. **Run** the compiled class using the `java` command:
   ```bash
   java -cp target/classes com.globomantics.Main
   ```

3. If you have external dependencies:
   - Copy dependencies to `target/dependency`:
     ```bash
     mvn dependency:copy-dependencies
     ```
   - Run with dependencies:
     ```bash
     java -cp target/classes:target/dependency/* com.globomantics.Main
     ```

### Alternative: Use Maven's Built-in Lifecycle Phases

Another common approach is to use Maven's **built-in lifecycle phases** to compile and package the program. You can create a JAR file and run it with dependencies:

1. **Package the Program**:
   ```bash
   mvn package
   ```

   This will create a JAR file in the `target` directory.

2. **Run the Program**:
   You can run the JAR using:

   ```bash
   java -jar target/your-jar-file.jar
   ```

This approach is simple, and by running `mvn compile` followed by `java -cp`, you avoid the need to add additional plugins or change the configuration in `pom.xml`.

# How to resolve the * pattern issue
The error `zsh: no matches found: target/classes:target/dependency/*` happens because **Zsh** is interpreting the `*` (wildcard) in the classpath as a pattern that needs to match files, and if there are no matches, it throws an error. This behavior is specific to **Zsh** (unlike Bash, which handles wildcards differently).

### Two Ways to Fix It:

#### 1. **Escape the Wildcard**
You can escape the `*` wildcard so that Zsh doesn't interpret it, and it gets passed correctly to the Java command:

```bash
java -cp target/classes:target/dependency/\* com.globomantics.Main
```

In this case, the `\*` ensures that the wildcard is passed as part of the classpath for Java to interpret, not Zsh.

#### 2. **Use Quotes Around the Classpath**
Alternatively, you can enclose the classpath in quotes, which prevents Zsh from interpreting the wildcard:

```bash
java -cp "target/classes:target/dependency/*" com.globomantics.Main
```

This way, Zsh won't expand the wildcard, and Java will be able to handle the classpath correctly.

### Explanation:
- **Zsh Shell**: In Zsh, unquoted wildcards are treated as patterns for file matching. If no file matches, Zsh will throw an error (`no matches found`).
- **Java Classpath**: When you use `target/dependency/*`, you are telling Java to include all JAR files from the `target/dependency/` directory in the classpath.

Either of these approaches should solve the issue, and the program should now run with the dependencies included.

# Method Reference 

Let’s break down the two lines of code and explain the `::` syntax and its meaning.

### 1. **`Thread haltedHook = new Thread(consumer::close);`**

- **`new Thread(consumer::close)`**: This creates a new `Thread` object.
- **`consumer::close`**: This is using a **method reference** in Java, introduced in **Java 8**. The `::` operator is the **method reference operator** and is used to refer to a method without invoking it. In this case, `consumer::close` refers to the `close` method of the `consumer` object, but it is not being called immediately. Instead, it will be called later when the thread starts.

#### Meaning:
- This statement creates a new thread, and the action performed by the thread is to invoke the `close()` method of the `consumer` object. However, the `close()` method will only be executed when the thread is started (e.g., during a shutdown event, which we'll see in the next line).

### 2. **`Runtime.getRuntime().addShutdownHook(haltedHook);`**

- **`Runtime.getRuntime()`**: This gets the singleton `Runtime` object associated with the current Java application. The `Runtime` class allows the application to interact with the environment in which the application is running (e.g., managing memory, adding shutdown hooks, etc.).
  
- **`addShutdownHook(Thread hook)`**: This method registers a **shutdown hook**, which is a thread that runs when the JVM is shutting down (either normally or due to an external signal like `Ctrl+C` or a termination signal). The thread passed to `addShutdownHook` will be executed when the JVM starts the shutdown process.

#### Meaning:
- This statement adds the `haltedHook` thread (which we defined earlier to close the `consumer`) as a shutdown hook to be executed when the JVM is shutting down. When the JVM receives a shutdown signal, it will call the `close()` method on the `consumer` object to ensure that resources are properly released before the application exits.

### The `::` Syntax (Method Reference)
The **`::`** is the **method reference** operator. It is a shorthand that allows us to refer to methods without executing them immediately. It's often used in situations where we want to pass a method as an argument to another method (such as passing the `consumer::close` method to the `Thread` constructor).

#### Common Use of `::` in Method References:
- **Static Method Reference**: `ClassName::staticMethod`
- **Instance Method Reference**: `instance::method`
- **Constructor Reference**: `ClassName::new`

In this case, `consumer::close` refers to an **instance method** of the `consumer` object.

### Putting It All Together:
```java
Thread haltedHook = new Thread(consumer::close);
Runtime.getRuntime().addShutdownHook(haltedHook);
```

- **`consumer::close`** refers to the `close()` method of the `consumer` object.
- A new `Thread` is created that will call the `close()` method when it runs.
- This thread (`haltedHook`) is registered as a **shutdown hook**, meaning it will run when the JVM is shutting down, ensuring the `close()` method of the `consumer` is called to gracefully release resources.

### Why Use a Shutdown Hook?
Shutdown hooks are useful when you want to ensure certain cleanup tasks (like closing network connections, releasing file handles, or flushing logs) happen when the application is terminated. In this case, the shutdown hook ensures that the Kafka consumer is closed properly when the application is shutting down.

### Example Use Case:
Imagine you are consuming messages from a Kafka topic. When the application is interrupted or shuts down, you want to ensure that the consumer properly closes its connections and commits any offsets. By adding the shutdown hook, the `consumer.close()` method will automatically be called when the JVM is shutting down, ensuring a graceful shutdown.

```java
// Process record

private static void processRecord(ConsumerRecord record) {
    log.info("Received message with key: " + record.key() + " and value " + record.value());
    log.info("It comes from partition: " + record.partition()); 
    try {
        Thread.sleep(1000);
    } catch (InterruptedException e) {
        System.out.println(e.getMessage)
    }
}
```

# How does ksqlDB connect to Kafka

**KSQL** (or **ksqlDB**, its more recent evolution) is a SQL-based stream processing platform built on top of **Apache Kafka**. It allows you to query, transform, and manipulate data in Kafka topics in real-time using SQL-like syntax. **ksqlDB** simplifies working with Kafka by enabling you to create, read, and process streams of data without writing complex code.

### How ksqlDB Connects to Kafka Topics:
**ksqlDB** (formerly known as KSQL) interacts with **Kafka topics** by leveraging Kafka's producer and consumer APIs. It reads from and writes to Kafka topics just like any other Kafka client (producers and consumers). Here's how it works:

### 1. **Streams and Tables** in ksqlDB:
- **Streams**: A **stream** in ksqlDB is an unbounded sequence of events, which corresponds to a Kafka topic where each message is treated as an event.
- **Tables**: A **table** in ksqlDB represents a changelog or an evolving state. Internally, it reads from a compacted Kafka topic, where the latest value for each key is maintained.

When you create a **stream** or **table** in ksqlDB, it is mapped to an underlying **Kafka topic**. You can think of streams and tables as abstractions over Kafka topics, allowing you to query and manipulate topic data using SQL.

### 2. **Creating a Stream from a Kafka Topic**:
You can create a ksqlDB **stream** that reads data from an existing Kafka topic with a simple SQL command. Here’s an example:

```sql
CREATE STREAM user_logins (
    user_id STRING,
    login_time BIGINT
) WITH (
    KAFKA_TOPIC = 'user_logins_topic',
    VALUE_FORMAT = 'JSON'
);
```

- **`CREATE STREAM`**: This creates a stream that is directly connected to the Kafka topic `user_logins_topic`.
- **`KAFKA_TOPIC = 'user_logins_topic'`**: Specifies the Kafka topic to connect to.
- **`VALUE_FORMAT = 'JSON'`**: Indicates the format of the messages in the Kafka topic. Other formats could include `AVRO` or `DELIMITED`.

Once the stream is created, ksqlDB will automatically **consume messages from the Kafka topic** and make them available for querying.

### 3. **Inserting Data into a Kafka Topic via ksqlDB**:
You can also use ksqlDB to produce (insert) data into Kafka topics, which can either be existing topics or new ones that ksqlDB creates. Here’s an example of inserting data:

```sql
INSERT INTO user_logins (user_id, login_time) VALUES ('user_123', 1632848200);
```

This will insert a new record into the `user_logins` stream, which is backed by the Kafka topic `user_logins_topic`.

### 4. **Querying Data from Kafka Topics**:
Once the stream is created, you can query the stream in real-time. For example, to get all login events from the `user_logins_topic`, you can run a `SELECT` query like this:

```sql
SELECT * FROM user_logins EMIT CHANGES;
```

This query continuously listens to the Kafka topic and outputs each new record that gets produced to the topic.

### 5. **Materialized Views with Tables**:
You can create a **table** in ksqlDB that represents an aggregated or "materialized" view of a Kafka topic. This is useful for tracking state or counts.

For example, you can create a table that counts the number of logins per user:

```sql
CREATE TABLE login_counts AS
SELECT user_id, COUNT(*) AS login_count
FROM user_logins
GROUP BY user_id;
```

Here, the `login_counts` table will continuously aggregate the login events and store the count for each user in a Kafka topic that backs the table.

### 6. **Kafka Consumer and Producer Underneath**:
Under the hood, ksqlDB acts as both a **Kafka consumer** and a **Kafka producer**:
- **Kafka Consumer**: When you create a stream, ksqlDB uses a Kafka consumer to read messages from the specified Kafka topic.
- **Kafka Producer**: When you insert data into a stream or when a query writes results to a new stream or table, ksqlDB produces the output back to a Kafka topic using Kafka's producer API.

### Summary of the Connection Between ksqlDB and Kafka Topics:
1. **Streams and tables** in ksqlDB are directly connected to **Kafka topics**.
2. **ksqlDB consumes data** from Kafka topics using Kafka consumers and processes the data according to SQL-like queries.
3. **ksqlDB produces data** (e.g., results from queries, transformations, or insertions) back to Kafka topics using Kafka producers.
4. You can query Kafka topics in real-time using `SELECT` statements on streams and tables.
5. You can manipulate, transform, and aggregate the data from Kafka topics via SQL queries in ksqlDB, and ksqlDB ensures the communication with Kafka topics behind the scenes.

This setup makes **ksqlDB** a powerful tool for streaming data processing, allowing users to easily work with Kafka topics using a familiar SQL syntax without needing to manage the low-level details of Kafka producers and consumers.

# How to connect ksqlDB to kafka topics

To connect **ksqlDB** to a Kafka cluster and interact with Kafka topics, you need to configure **ksqlDB** to be aware of the **Kafka broker addresses**, **Zookeeper**, and possibly other components depending on your deployment environment.

Here’s a step-by-step guide to setting up **ksqlDB** to connect to Kafka, including the configuration details like the **broker address**, schema registry (if applicable), and other necessary settings.

### 1. **Install and Start Kafka**
First, you need to have a running **Kafka cluster** and optionally **Zookeeper** (if you are not using Kafka's newer KRaft mode). If you already have Kafka set up, proceed to the next steps. If not, follow these steps:

- Download the Kafka distribution from [Apache Kafka Downloads](https://kafka.apache.org/downloads).
- Start Zookeeper (if required) and Kafka brokers.

Example commands to start Zookeeper and Kafka brokers:

```bash
# Start Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties

# Start Kafka broker
bin/kafka-server-start.sh config/server.properties
```

### 2. **Install ksqlDB**
You can install **ksqlDB** using Docker, a package manager, or manually download and run it.

#### a. **Running ksqlDB with Docker**:
Using Docker is the easiest way to set up **ksqlDB**:

```bash
docker run -d --name ksqldb-server \
  -p 8088:8088 \
  -e KSQL_CONFIG_DIR="/etc/ksqldb" \
  -e KSQL_BOOTSTRAP_SERVERS="your-kafka-broker:9092" \
  -e KSQL_LISTENERS="http://0.0.0.0:8088" \
  confluentinc/ksqldb-server:latest
```

In this command:
- **`KSQL_BOOTSTRAP_SERVERS`**: Specifies the Kafka broker addresses that **ksqlDB** will use to connect to Kafka. Replace `your-kafka-broker:9092` with the actual Kafka broker address(es).
- **`KSQL_LISTENERS`**: Configures the interface and port on which ksqlDB listens for requests (HTTP REST API). In this case, ksqlDB will be accessible on `http://localhost:8088`.

#### b. **Running ksqlDB Locally**:
Download **ksqlDB** from the official [Confluent Downloads](https://ksqldb.io/), then follow these steps:

```bash
# Extract the ksqlDB tarball
tar -xzf confluent-7.x.x.tar.gz
cd confluent-7.x.x

# Start ksqlDB with required configurations
bin/ksql-server-start etc/ksqldb/ksql-server.properties
```

Make sure that `ksql-server.properties` contains the necessary configurations like Kafka brokers and Zookeeper.

### 3. **Configure ksqlDB to Connect to Kafka**
To connect **ksqlDB** to your Kafka cluster, you need to set several key configurations in **ksqlDB's configuration file** (`ksql-server.properties`) or environment variables.

Here are the important properties:

#### a. **Kafka Broker Configuration** (`bootstrap.servers`):
This property tells **ksqlDB** where the Kafka brokers are located so that it can interact with the Kafka cluster.

```properties
# The address of your Kafka brokers
bootstrap.servers=PLAINTEXT://your-kafka-broker:9092
```

You can specify multiple brokers by separating them with commas:
```properties
bootstrap.servers=PLAINTEXT://broker1:9092,PLAINTEXT://broker2:9092,PLAINTEXT://broker3:9092
```

- `PLAINTEXT://` specifies the protocol. You can also use `SSL://` or `SASL_PLAINTEXT://` if your Kafka cluster requires secure connections.
- **Port 9092** is the default port for Kafka brokers.

#### b. **Advertised Kafka Client Configurations**:
In addition to `bootstrap.servers`, **ksqlDB** needs information about how it should connect as a Kafka client. For example:

```properties
listeners=http://0.0.0.0:8088
```
This tells **ksqlDB** to expose its REST interface on port **8088**.

#### c. **Zookeeper Configuration** (if applicable):
If you’re using Zookeeper to manage your Kafka cluster, you might also need to specify the **Zookeeper connection string**. However, with the newer **KRaft mode** in Kafka (introduced in Kafka 2.8), Zookeeper is no longer required for broker management.

```properties
zookeeper.connect=localhost:2181
```

This tells **ksqlDB** how to connect to the Zookeeper node that manages the Kafka brokers.

### 4. **Schema Registry (Optional)**
If you’re using **Avro** or **Protobuf** data formats in Kafka, you’ll likely need to configure **Confluent Schema Registry**. The Schema Registry provides a way to manage and validate message schemas in Kafka topics.

Add this line to your `ksql-server.properties` file or Docker environment variables:

```properties
ksql.schema.registry.url=http://schema-registry:8081
```

This tells **ksqlDB** where the Schema Registry is located.

### 5. **Starting ksqlDB and Testing the Connection**

Once you have configured `ksql-server.properties` or set up the environment variables, you can start **ksqlDB**.

To test that **ksqlDB** can connect to Kafka, you can access the **ksqlDB CLI** or the **REST API**:

#### a. **Using the CLI**:
You can use the ksqlDB CLI to interact with your Kafka cluster. Run the following command to start the CLI:

```bash
docker exec -it ksqldb-server ksql http://localhost:8088
```

You can now create streams, query data, and work with Kafka topics using SQL-like syntax.

#### b. **Using the REST API**:
You can also use **ksqlDB's REST API** to interact with the Kafka cluster. For example, you can create a stream using a POST request to the REST API.

```bash
curl -X POST "http://localhost:8088/ksql" \
     -H "Content-Type: application/vnd.ksql.v1+json; charset=utf-8" \
     --data '{
       "ksql": "CREATE STREAM user_logins (user_id STRING, login_time BIGINT) WITH (KAFKA_TOPIC = '\''user_logins_topic'\'', VALUE_FORMAT = '\''JSON'\'');",
       "streamsProperties": {}
     }'
```

This creates a stream that reads from the `user_logins_topic` Kafka topic.

### 6. **Creating a Stream or Table in ksqlDB (Connection to Kafka Topic)**
Once **ksqlDB** is properly connected to Kafka, you can create a **stream** or **table** by connecting it to a Kafka topic. Here’s an example of creating a stream from a Kafka topic:

```sql
CREATE STREAM user_logins (
  user_id STRING,
  login_time BIGINT
) WITH (
  KAFKA_TOPIC = 'user_logins_topic',
  VALUE_FORMAT = 'JSON'
);
```

This command tells **ksqlDB** to consume data from the Kafka topic `user_logins_topic`.

### 7. **Querying Data in Real-Time**
Once the stream is created, you can run real-time queries against the Kafka topic. For example:

```sql
SELECT * FROM user_logins EMIT CHANGES;
```

This will display new records from the `user_logins_topic` Kafka topic as they are ingested by Kafka and processed by **ksqlDB**.

---

### Summary:

1. **ksqlDB** connects to **Kafka** by specifying the **Kafka broker addresses** via the `bootstrap.servers` property in `ksql-server.properties` or Docker environment variables.
2. **ksqlDB** communicates with Kafka like any other Kafka client, using the **Kafka producer and consumer APIs** to read from and write to Kafka topics.
3. **Streams and tables** in ksqlDB are abstractions over Kafka topics, allowing you to perform SQL-like queries on real-time Kafka data.
4. You can configure additional properties like the **Schema Registry** and **Zookeeper** (if applicable) to enable schema validation and distributed cluster management.
5. You interact with ksqlDB via its **CLI** or **REST API** to create streams, tables, and perform real-time queries against Kafka topics.

This setup ensures that **ksqlDB** is fully integrated with your Kafka cluster, enabling powerful, real-time stream processing and querying.