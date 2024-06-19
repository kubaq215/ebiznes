import io.ktor.application.*
import io.ktor.http.*
import io.ktor.response.*
import io.ktor.routing.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import kotlinx.coroutines.runBlocking

fun main() {
    val token = "MTI1MzA5ODg1OTIyODI5OTMwNQ.Gwm7_J.cV17IbF52e-vIlyGbUutOa5XHVsJGtVvdSvv-M"
    val bot = DiscordBot(token)

    runBlocking {
        bot.start()
    }

    embeddedServer(Netty, port = 8080) {
        routing {
            get("/") {
                call.respondText("Hello, world!", ContentType.Text.Plain)
            }
            get("/categories") {
                call.respondText("Available categories: Category1, Category2, Category3", ContentType.Text.Plain)
            }
            get("/products/{category}") {
                val category = call.parameters["category"]
                call.respondText("Products for $category: Product1, Product2, Product3", ContentType.Text.Plain)
            }
        }
    }.start(wait = true)
}
