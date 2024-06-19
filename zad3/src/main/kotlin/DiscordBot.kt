import dev.kord.core.Kord
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import kotlinx.coroutines.runBlocking

class DiscordBot(private val token: String) {
    private val kord = runBlocking { Kord(token) }

    suspend fun start() {
        kord.on<MessageCreateEvent> {
            if (message.content.startsWith("!categories")) {
                message.channel.createMessage("Available categories: Category1, Category2, Category3")
            }
            if (message.content.startsWith("!products")) {
                val category = message.content.removePrefix("!products").trim()
                message.channel.createMessage("Products for $category: Product1, Product2, Product3")
            }
        }

        kord.login()
    }
}
