package org.xqdl
package utils

import protocols.{ProcessedSale, RawSale}

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.scala.DefaultScalaModule

object Json {
  def deserialize(jsonString: String): RawSale = {
    // ObjectMapper for JSON parsing
    val objectMapper = new ObjectMapper()
    objectMapper.registerModule(DefaultScalaModule)

    // Deserialize the JSON string into a case class instance
    objectMapper.readValue(jsonString, classOf[RawSale])
  }

  def serialize(sale: ProcessedSale): String = {
    val objectMapper = new ObjectMapper()
    objectMapper.registerModule(DefaultScalaModule)

    // Serialize the case class instance to a JSON string
    val jsonString: String = objectMapper.writeValueAsString(sale)

    // Now you have the JSON string
    jsonString
    // Output: {"name":"John Doe","age":30,"email":"johndoe@example.com"}
  }
}