package org.xqdl
package protocols

case class ProcessedSale(id: String,
                         salesman: String,
                         customer: String,
                         brand: String,
                         product: String,
                         originalPrice: Float,
                         discountRate: Int,
                         timestamp: String,
                         price: Float
                         )
