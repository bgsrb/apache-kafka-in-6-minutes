# Apache Kafka in 6 minutes
A quick introduction to how Apache Kafka works and differs from other messaging systems using an example application. In this video I explain partitioning, consumer offsets, replication and many other concepts found in Kafka.

[![Apache Kafka in 6 minutes](http://img.youtube.com/vi/Ch5VhJzaoaI/0.jpg)](https://www.youtube.com/watch?v=Ch5VhJzaoaI "Apache Kafka in 6 minutes")

Output
```
producer_1 -> America Canada : Querter Start
producer_1 -> Malta Portugal : Foul
producer_1 -> America Canada : Score 39-46
producer_1 -> Brazil Australia : Free Throw
producer_1 -> America Canada : Score 41-46
producer_1 -> Brazil Australia : Querter End
--------------------------------------------
mobile_consumer_2 -> America Canada : Querter Start
mobile_consumer_3 -> Malta Portugal : Foul
mobile_consumer_1 -> America Canada : Score 39-46
mobile_consumer_2 -> Brazil Australia : Free Throw
mobile_consumer_1 -> Brazil Australia : Querter End
mobile_consumer_3 -> America Canada : Score 41-46
computer_consumer_2 -> America Canada : Querter Start
computer_consumer_2 -> Brazil Australia : Free Throw
computer_consumer_3 -> America Canada : Score 39-46
computer_consumer_1 -> Malta Portugal : Foul
computer_consumer_3 -> Brazil Australia : Querter End
computer_consumer_1 -> America Canada : Score 41-46
```
