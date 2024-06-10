package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.Category

@Singleton
class CategoryController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  private var categories = List(
    Category(1, "Category A"),
    Category(2, "Category B")
  )

  def getAllCategories: Action[AnyContent] = Action {
    Ok(Json.toJson(categories))
  }

  def getCategoryById(id: Int): Action[AnyContent] = Action {
    categories.find(_.id == id) match {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound(Json.obj("error" -> "Category not found"))
    }
  }

  def createCategory: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid JSON")),
      category => {
        categories = categories :+ category
        Created(Json.toJson(category))
      }
    )
  }

  def updateCategory(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest(Json.obj("error" -> "Invalid JSON")),
      updatedCategory => {
        categories = categories.map { category =>
          if (category.id == id) updatedCategory else category
        }
        Ok(Json.toJson(updatedCategory))
      }
    )
  }

  def deleteCategory(id: Int): Action[AnyContent] = Action {
    categories = categories.filterNot(_.id == id)
    NoContent
  }
}
