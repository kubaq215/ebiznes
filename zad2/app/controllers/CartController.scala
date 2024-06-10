package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.CartItem

@Singleton
class CartController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  private var cart = List(
    CartItem(1, 1, 2),
    CartItem(2, 2, 1)
  )

  def getAllCartItems: Action[AnyContent] = Action {
    Ok(Json.toJson(cart))
  }

  def getCartItemById(id: Int): Action[AnyContent] = Action {
    cart.find(_.id == id) match {
      case Some(cartItem) => Ok(Json.toJson(cartItem))
      case None => NotFound(Json.obj("error" -> "Cart item not found"))
    }
  }

  def addCartItem: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid JSON")),
      cartItem => {
        cart = cart :+ cartItem
        Created(Json.toJson(cartItem))
      }
    )
  }

  def updateCartItem(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[CartItem].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid JSON")),
      updatedCartItem => {
        cart = cart.map { item =>
          if (item.id == id) updatedCartItem else item
        }
        Ok(Json.toJson(updatedCartItem))
      }
    )
  }

  def deleteCartItem(id: Int): Action[AnyContent] = Action {
    cart = cart.filterNot(_.id == id)
    NoContent
  }
}
