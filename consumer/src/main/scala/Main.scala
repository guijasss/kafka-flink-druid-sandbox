package org.xqdl

import org.apache.flink.api.common.eventtime.WatermarkStrategy
import org.apache.flink.api.common.serialization.SimpleStringSchema
import org.apache.flink.connector.kafka.sink.{KafkaRecordSerializationSchema, KafkaSink}
import org.apache.flink.connector.kafka.source.KafkaSource
import org.apache.flink.connector.kafka.source.enumerator.initializer.OffsetsInitializer
import org.apache.flink.streaming.api.scala._

//https://medium.com/@erkansirin/apache-flink-and-kafka-simple-example-with-scala-97f0b338ee36

object Main {
  def main(args: Array[String]): Unit = {
    val env = StreamExecutionEnvironment.getExecutionEnvironment

    val kafkaSource = KafkaSource.builder()
      .setBootstrapServers("localhost:9092")
      .setTopics("ingestion-sales")
      .setStartingOffsets(OffsetsInitializer.latest())
      .setValueOnlyDeserializer(new SimpleStringSchema())
      .build()

    val serializer = KafkaRecordSerializationSchema.builder()
      .setValueSerializationSchema(new SimpleStringSchema())
      .setTopic("load-sales")
      .build()

    val kafkaSink = KafkaSink.builder()
      .setBootstrapServers("localhost:9092")
      .setRecordSerializer(serializer)
      .build()

    val lines = env.fromSource(kafkaSource, WatermarkStrategy.noWatermarks(), "Kafka Source")
    lines.print()
    lines.sinkTo(kafkaSink)
    env.execute("Read from Kafka")
  }
}
