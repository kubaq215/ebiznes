package models

import play.api.libs.json._

case class CartItem(id: Int, productId: Int, quantity: Int)
object CartItem {
  implicit val cartItemFormat: OFormat[CartItem] = Json.format[CartItem]
}
