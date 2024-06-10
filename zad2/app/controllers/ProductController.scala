package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.Product

@Singleton
class ProductController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  private var products = List(
    Product(1, "Product A", 10.0),
    Product(2, "Product B", 20.0)
  )

  def getAllProducts: Action[AnyContent] = Action {
    Ok(Json.toJson(products))
  }

  def getProductById(id: Int): Action[AnyContent] = Action {
    products.find(_.id == id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(Json.obj("error" -> "Product not found"))
    }
  }

  def createProduct: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid JSON")),
      product => {
        products = products :+ product
        Created(Json.toJson(product))
      }
    )
  }

  def updateProduct(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid JSON")),
      updatedProduct => {
        products = products.map { product =>
          if (product.id == id) updatedProduct else product
        }
        Ok(Json.toJson(updatedProduct))
      }
    )
  }

  def deleteProduct(id: Int): Action[AnyContent] = Action {
    products = products.filterNot(_.id == id)
    NoContent
  }
}
